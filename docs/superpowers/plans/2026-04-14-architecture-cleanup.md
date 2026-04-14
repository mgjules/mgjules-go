# Architecture Cleanup Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Clean up dead code, reduce fetcher repetition, and improve structural boundaries across the codebase.

**Architecture:** Three phases — dead code removal first (safe, no behavior change), then fetcher deduplication (reduces ~180 lines to ~60), then structural improvements (error returns, interface decoupling, lock fix). Each phase produces a compilable, working state.

**Tech Stack:** Go 1.24, templ, resty, ants, cron, endless

---

## File Structure

### Files to delete
- `internal/repository/edgedb/` (15 files — entire directory)
- `internal/db/edgedb/` (schema + 28 migrations — entire directory)
- `edgedb.toml`
- `internal/fetcher/copier.go`

### Files to modify
- `internal/config/config.go` — remove EdgeDBDSN field
- `internal/repository/repository.go` — remove EdgeDB branch, return error
- `internal/fetcher/fetcher.go` — generic fetchEntity, remove DeepCopy, fix notify lock
- `internal/projecter/projecter.go` — DataProvider interface, return error from New, remove comments
- `internal/http/server.go` — return error from NewServer, remove commented routes
- `internal/auth/auth.go` — return error from New
- `cmd/serve/main.go` — handle constructor errors
- `go.mod` / `go.sum` — cleaned by go mod tidy

---

### Task 1: Delete EdgeDB repository and related files

**Files:**
- Delete: `internal/repository/edgedb/` (entire directory)
- Delete: `internal/db/edgedb/` (entire directory)
- Delete: `edgedb.toml`

- [ ] **Step 1: Delete the EdgeDB repository directory**

```bash
rm -rf internal/repository/edgedb/
```

- [ ] **Step 2: Delete the EdgeDB schema and migrations directory**

```bash
rm -rf internal/db/edgedb/
```

- [ ] **Step 3: Delete the EdgeDB config file**

```bash
rm edgedb.toml
```

- [ ] **Step 4: Remove EdgeDBDSN from config**

In `internal/config/config.go`, remove the EdgeDBDSN field. The full file becomes:

```go
package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	// General
	Prod bool `envconfig:"PROD"`

	// HTTP server
	ServerHost      string `envconfig:"SERVER_HOST" default:"0.0.0.0"`
	ServerPort      int    `envconfig:"SERVER_PORT" default:"13337"`
	ServerTLSDomain string `envconfig:"SERVER_TLS_DOMAIN"`

	// Authentication
	AuthToken string `envconfig:"AUTH_TOKEN" required:"true"`

	// Static
	Static bool `envconfig:"STATIC"`

	// Directus
	DirectusURL   string `envconfig:"DIRECTUS_URL"`
	DirectusToken string `envconfig:"DIRECTUS_TOKEN"`
}

func Parse() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}
	return &cfg, nil
}
```

- [ ] **Step 5: Simplify repository factory — remove EdgeDB branch**

In `internal/repository/repository.go`, remove the EdgeDB import and branch. The full file becomes:

```go
package repository

import (
	"context"
	"errors"

	"github.com/mgjules/mgjules-go/internal/config"
	"github.com/mgjules/mgjules-go/internal/entity"
	"github.com/mgjules/mgjules-go/internal/repository/directus"
	"github.com/mgjules/mgjules-go/internal/repository/static"
)

type Repository interface {
	GetMeta(ctx context.Context, id string) (*entity.Meta, error)
	GetLinks(ctx context.Context) ([]entity.Link, error)
	GetIntroduction(ctx context.Context, id string) (*entity.Introduction, error)
	GetSections(ctx context.Context) ([]entity.Section, error)
	GetExperiences(ctx context.Context) ([]entity.Experience, error)
	GetProjects(ctx context.Context) ([]entity.Project, error)
	GetContributions(ctx context.Context) ([]entity.Contribution, error)
	GetAwards(ctx context.Context) ([]entity.Award, error)
	GetInterests(ctx context.Context) ([]entity.Interest, error)
	GetLanguages(ctx context.Context) ([]entity.Language, error)
	GetPosts(ctx context.Context) ([]entity.Post, error)
	Cleanup() error
}

func New(ctx context.Context, cfg *config.Config) (Repository, error) {
	if cfg.DirectusURL != "" {
		return directus.New(ctx, cfg.Prod, cfg.DirectusURL, cfg.DirectusToken), nil
	}
	if cfg.Static {
		return static.New(), nil
	}
	return nil, errors.New("no valid repository configuration found")
}
```

- [ ] **Step 6: Update main.go to handle repository error**

In `cmd/serve/main.go`, change the `repository.New` call from:

```go
repo := repository.New(ctx, cfg)
```

to:

```go
repo, err := repository.New(ctx, cfg)
if err != nil {
    log.Fatalf("failed to create repository: %v", err)
}
```

- [ ] **Step 7: Verify it compiles**

```bash
go build ./...
```

Expected: success, no errors.

- [ ] **Step 8: Run go mod tidy to remove orphaned dependencies**

```bash
go mod tidy
```

This removes `github.com/edgedb/edgedb-go` and its transitive deps (`certifi/gocertifi`, `sigurn/crc16`, `xdg/scram`, `xdg/stringprep`).

- [ ] **Step 9: Verify it still compiles after tidy**

```bash
go build ./...
```

Expected: success.

- [ ] **Step 10: Commit**

```bash
git add -A
git commit -m "Remove deprecated EdgeDB repository and related files"
```

---

### Task 2: Remove commented-out code

**Files:**
- Modify: `internal/http/server.go:89-92`
- Modify: `internal/projecter/projecter.go:61-64`

- [ ] **Step 1: Remove commented-out routes from server.go**

In `internal/http/server.go`, replace the `initRoutes` method. Change from:

```go
func (s *Server) initRoutes() {
	s.mux.HandleFunc("GET /", s.rootHandler())
	// s.mux.HandleFunc("GET /blog", s.BlogIndexHandler())
	// s.mux.HandleFunc("GET /blog/{slug}", s.BlogPostHandler())

	// s.mux.Handle("/", http.StripPrefix("/", http.FileServerFS(s.static)))
}
```

to:

```go
func (s *Server) initRoutes() {
	s.mux.HandleFunc("GET /", s.rootHandler())
}
```

- [ ] **Step 2: Remove commented-out fetcher calls from projecter.go**

In `internal/projecter/projecter.go`, in the `Build()` method, change from:

```go
	projects := p.fetcher.Projects()
	// contributions := p.fetcher.Contributions()
	// awards := p.fetcher.Awards()
	// interests := p.fetcher.Interests()
	// languages := p.fetcher.Languages()
	posts := p.fetcher.Posts()
```

to:

```go
	projects := p.fetcher.Projects()
	posts := p.fetcher.Posts()
```

- [ ] **Step 3: Verify it compiles**

```bash
go build ./...
```

Expected: success.

- [ ] **Step 4: Commit**

```bash
git add internal/http/server.go internal/projecter/projecter.go
git commit -m "Remove commented-out code from server and projecter"
```

---

### Task 3: Deduplicate Fetcher.Fetch() with generic helper

**Files:**
- Modify: `internal/fetcher/fetcher.go`

- [ ] **Step 1: Add the generic fetchEntity helper**

In `internal/fetcher/fetcher.go`, add this function after the `notify()` method (before the `Fetch()` method):

```go
func fetchEntity[T any](
	f *Fetcher,
	ctx context.Context,
	wg *sync.WaitGroup,
	name string,
	fetch func(context.Context) (T, error),
	store func(T),
) {
	wg.Add(1)
	f.pool.Submit(func() {
		defer wg.Done()
		slog.Debug("Fetching data...", "entity", name)
		result, err := fetch(ctx)
		if err != nil {
			slog.Error("failed to fetch", "entity", name, "error", err)
			return
		}
		f.dataMu.Lock()
		store(result)
		f.dataMu.Unlock()
	})
}
```

- [ ] **Step 2: Replace the Fetch() method body**

Replace the entire `Fetch()` method with:

```go
func (f *Fetcher) Fetch() {
	operation := func() error {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		var wg sync.WaitGroup

		fetchEntity(f, ctx, &wg, "meta",
			func(ctx context.Context) (*entity.Meta, error) {
				return f.repo.GetMeta(ctx, "bd99e066-440b-11ed-924c-9fd15527df84")
			},
			func(v *entity.Meta) { f.meta = v })
		fetchEntity(f, ctx, &wg, "links",
			func(ctx context.Context) ([]entity.Link, error) { return f.repo.GetLinks(ctx) },
			func(v []entity.Link) { f.links = v })
		fetchEntity(f, ctx, &wg, "intro",
			func(ctx context.Context) (*entity.Introduction, error) {
				return f.repo.GetIntroduction(ctx, "a4296eac-441b-11ed-924c-830c8fd1144c")
			},
			func(v *entity.Introduction) { f.intro = v })
		fetchEntity(f, ctx, &wg, "sections",
			func(ctx context.Context) ([]entity.Section, error) { return f.repo.GetSections(ctx) },
			func(v []entity.Section) { f.sections = v })
		fetchEntity(f, ctx, &wg, "experiences",
			func(ctx context.Context) ([]entity.Experience, error) { return f.repo.GetExperiences(ctx) },
			func(v []entity.Experience) { f.experiences = v })
		fetchEntity(f, ctx, &wg, "projects",
			func(ctx context.Context) ([]entity.Project, error) { return f.repo.GetProjects(ctx) },
			func(v []entity.Project) { f.projects = v })
		fetchEntity(f, ctx, &wg, "contributions",
			func(ctx context.Context) ([]entity.Contribution, error) { return f.repo.GetContributions(ctx) },
			func(v []entity.Contribution) { f.contributions = v })
		fetchEntity(f, ctx, &wg, "awards",
			func(ctx context.Context) ([]entity.Award, error) { return f.repo.GetAwards(ctx) },
			func(v []entity.Award) { f.awards = v })
		fetchEntity(f, ctx, &wg, "interests",
			func(ctx context.Context) ([]entity.Interest, error) { return f.repo.GetInterests(ctx) },
			func(v []entity.Interest) { f.interests = v })
		fetchEntity(f, ctx, &wg, "languages",
			func(ctx context.Context) ([]entity.Language, error) { return f.repo.GetLanguages(ctx) },
			func(v []entity.Language) { f.languages = v })
		fetchEntity(f, ctx, &wg, "posts",
			func(ctx context.Context) ([]entity.Post, error) { return f.repo.GetPosts(ctx) },
			func(v []entity.Post) { f.posts = v })

		wg.Wait()
		return nil
	}

	exp := backoff.NewExponentialBackOff()
	exp.InitialInterval = 5 * time.Second
	exp.Multiplier = 2.0
	exp.MaxInterval = 5 * time.Minute
	exp.MaxElapsedTime = 15 * time.Minute

	var bkf backoff.BackOff = exp
	if !f.fetchedAt.IsZero() {
		bkf = backoff.WithMaxRetries(bkf, maxFreshFetchAttempts)
	}

	err := backoff.Retry(operation, bkf)
	if err != nil {
		slog.Error("failed to fetched data", "error", err)
		return
	}

	f.dataMu.Lock()
	f.fetchedAt = time.Now()
	f.dataMu.Unlock()

	f.notify()
}
```

- [ ] **Step 3: Remove the `errors` import if no longer used**

Check `internal/fetcher/fetcher.go` imports. The `errors` package was only used by the old `Fetch()`. Remove it from the import block if present.

- [ ] **Step 4: Verify it compiles**

```bash
go build ./...
```

Expected: success.

- [ ] **Step 5: Commit**

```bash
git add internal/fetcher/fetcher.go
git commit -m "Deduplicate Fetcher.Fetch with generic fetchEntity helper"
```

---

### Task 4: Remove DeepCopy and copier dependency

**Files:**
- Delete: `internal/fetcher/copier.go`
- Modify: `internal/fetcher/fetcher.go`

- [ ] **Step 1: Delete copier.go**

```bash
rm internal/fetcher/copier.go
```

- [ ] **Step 2: Update accessor methods to return directly**

In `internal/fetcher/fetcher.go`, replace all accessor methods (from `Meta()` through `FetchedAt()`) with:

```go
func (f *Fetcher) Meta() entity.Meta {
	f.dataMu.RLock()
	defer f.dataMu.RUnlock()
	if f.meta == nil {
		return entity.Meta{}
	}
	return *f.meta
}

func (f *Fetcher) Links() []entity.Link {
	f.dataMu.RLock()
	defer f.dataMu.RUnlock()
	return f.links
}

func (f *Fetcher) Introduction() entity.Introduction {
	f.dataMu.RLock()
	defer f.dataMu.RUnlock()
	if f.intro == nil {
		return entity.Introduction{}
	}
	return *f.intro
}

func (f *Fetcher) Sections() []entity.Section {
	f.dataMu.RLock()
	defer f.dataMu.RUnlock()
	return f.sections
}

func (f *Fetcher) Experiences() []entity.Experience {
	f.dataMu.RLock()
	defer f.dataMu.RUnlock()
	return f.experiences
}

func (f *Fetcher) Projects() []entity.Project {
	f.dataMu.RLock()
	defer f.dataMu.RUnlock()
	return f.projects
}

func (f *Fetcher) Contributions() []entity.Contribution {
	f.dataMu.RLock()
	defer f.dataMu.RUnlock()
	return f.contributions
}

func (f *Fetcher) Awards() []entity.Award {
	f.dataMu.RLock()
	defer f.dataMu.RUnlock()
	return f.awards
}

func (f *Fetcher) Interests() []entity.Interest {
	f.dataMu.RLock()
	defer f.dataMu.RUnlock()
	return f.interests
}

func (f *Fetcher) Languages() []entity.Language {
	f.dataMu.RLock()
	defer f.dataMu.RUnlock()
	return f.languages
}

func (f *Fetcher) Posts() []entity.Post {
	f.dataMu.RLock()
	defer f.dataMu.RUnlock()
	return f.posts
}

func (f *Fetcher) FetchedAt() time.Time {
	f.dataMu.RLock()
	defer f.dataMu.RUnlock()
	return f.fetchedAt
}
```

- [ ] **Step 3: Remove copier import from fetcher.go if present**

The `copier` import should already be gone since the accessor methods no longer call `DeepCopy`. Verify the import block in `internal/fetcher/fetcher.go` does not reference `jinzhu/copier`.

- [ ] **Step 4: Run go mod tidy**

```bash
go mod tidy
```

This removes `github.com/jinzhu/copier` from go.mod.

- [ ] **Step 5: Verify it compiles**

```bash
go build ./...
```

Expected: success.

- [ ] **Step 6: Commit**

```bash
git add -A
git commit -m "Remove DeepCopy helper and copier dependency"
```

---

### Task 5: Fix fetcher notify() race condition

**Files:**
- Modify: `internal/fetcher/fetcher.go`

- [ ] **Step 1: Add read lock to notify()**

In `internal/fetcher/fetcher.go`, replace the `notify()` method. Change from:

```go
func (f *Fetcher) notify() {
	for _, s := range f.subscribers {
		if s == nil {
			continue
		}

		f.pool.Submit(s)
	}
}
```

to:

```go
func (f *Fetcher) notify() {
	f.subscribersMu.RLock()
	defer f.subscribersMu.RUnlock()
	for _, s := range f.subscribers {
		if s == nil {
			continue
		}
		f.pool.Submit(s)
	}
}
```

- [ ] **Step 2: Verify it compiles**

```bash
go build ./...
```

Expected: success.

- [ ] **Step 3: Commit**

```bash
git add internal/fetcher/fetcher.go
git commit -m "Fix race condition in fetcher notify()"
```

---

### Task 6: Decouple projecter from concrete fetcher and convert all constructors to error returns

**Files:**
- Modify: `internal/projecter/projecter.go`
- Modify: `internal/auth/auth.go`
- Modify: `internal/fetcher/fetcher.go`
- Modify: `internal/http/server.go`
- Modify: `cmd/serve/main.go`

These changes must land together because the projecter's `New()` signature change (accepting `DataProvider` instead of `*fetcher.Fetcher`) and the `DataProvider` interface definition must exist in the same commit to compile.

- [ ] **Step 1: Add DataProvider interface, rename field, and update New() in projecter**

Replace the entire `internal/projecter/projecter.go` with:

```go
package projecter

import (
	"bytes"
	"compress/gzip"
	"context"
	"errors"
	"log/slog"
	"slices"
	"strings"
	"sync"

	"github.com/a-h/templ"
	"github.com/mgjules/mgjules-go/internal/entity"
	"github.com/mgjules/mgjules-go/internal/templates/minimal"
	"github.com/panjf2000/ants/v2"
)

const (
	seperator string = "/"
)

type DataProvider interface {
	Meta() entity.Meta
	Links() []entity.Link
	Introduction() entity.Introduction
	Sections() []entity.Section
	Experiences() []entity.Experience
	Projects() []entity.Project
	Contributions() []entity.Contribution
	Awards() []entity.Award
	Interests() []entity.Interest
	Languages() []entity.Language
	Posts() []entity.Post
}

type Projecter struct {
	prod bool
	pool *ants.Pool
	data DataProvider

	mu          sync.RWMutex
	projections map[string][]byte
}

func New(
	prod bool,
	pool *ants.Pool,
	data DataProvider,
) (*Projecter, error) {
	if pool == nil {
		return nil, errors.New("pool cannot be nil")
	}
	if data == nil {
		return nil, errors.New("data provider cannot be nil")
	}

	return &Projecter{
		prod:        prod,
		pool:        pool,
		data:        data,
		projections: make(map[string][]byte),
	}, nil
}

func (p *Projecter) Build() {
	slog.Debug("Building projections...")
	defer slog.Debug("Built projections successfully")

	meta := p.data.Meta()
	links := p.data.Links()
	intro := p.data.Introduction()
	sections := p.data.Sections()
	experiences := p.data.Experiences()
	projects := p.data.Projects()
	posts := p.data.Posts()

	ctx := context.Background()
	b := &bytes.Buffer{}
	w, err := gzip.NewWriterLevel(b, gzip.BestCompression)
	if err != nil {
		slog.Error("failed to create new gzip writer", "error", err)
		return
	}
	defer w.Close()

	p.render(ctx, w, b, buildKey("index"), minimal.Index(meta, links, intro, sections, experiences, projects))
	p.render(ctx, w, b, buildKey("blog"), minimal.BlogIndex(meta, links, intro, posts))
	for _, post := range posts {
		p.render(ctx, w, b, buildKey("blog", post.Slug), minimal.BlogPost(meta, links, intro, post))
	}
	p.render(ctx, w, b, buildKey("404"), minimal.NotFound(meta, links, intro))
}

func (p *Projecter) Get(keys ...string) ([]byte, bool) {
	if len(keys) == 0 {
		return nil, false
	}

	if !p.prod {
		p.Build()
	}

	p.mu.RLock()
	out, found := p.projections[buildKey(keys...)]
	p.mu.RUnlock()

	return out, found
}

func (p *Projecter) render(
	ctx context.Context,
	w *gzip.Writer,
	b *bytes.Buffer,
	key string,
	comp templ.Component,
) {
	if err := comp.Render(ctx, w); err != nil {
		slog.Error("failed to create %q component", "key", key, "error", err)
	} else {
		p.mu.Lock()
		w.Flush()
		p.projections[key] = slices.Clone(b.Bytes())
		b.Reset()
		w.Reset(b)
		p.mu.Unlock()
	}
}

func buildKey(keys ...string) string {
	return strings.Join(keys, seperator)
}
```

- [ ] **Step 2: Update auth.New() to return error**

Replace `internal/auth/auth.go` with:

```go
package auth

import "errors"

type Auth struct {
	token string
}

func New(token string) (*Auth, error) {
	if token == "" {
		return nil, errors.New("token cannot be empty")
	}
	return &Auth{token: token}, nil
}

func (a *Auth) Validate(token string) bool {
	return a.token == token
}
```

- [ ] **Step 3: Update fetcher.New() to return error**

In `internal/fetcher/fetcher.go`, replace the `New` function. Change from:

```go
func New(repo repository.Repository, pool *ants.Pool) *Fetcher {
	if repo == nil {
		panic("repo cannot be nil")
	}
	if pool == nil {
		panic("pool cannot be nil")
	}

	f := &Fetcher{
		repo: repo,
		pool: pool,
		cron: cron.New(),
	}

	f.cron.AddFunc("@hourly", func() {
		f.Fetch()
	})

	return f
}
```

to:

```go
func New(repo repository.Repository, pool *ants.Pool) (*Fetcher, error) {
	if repo == nil {
		return nil, errors.New("repo cannot be nil")
	}
	if pool == nil {
		return nil, errors.New("pool cannot be nil")
	}

	f := &Fetcher{
		repo: repo,
		pool: pool,
		cron: cron.New(),
	}

	f.cron.AddFunc("@hourly", func() {
		f.Fetch()
	})

	return f, nil
}
```

Note: re-add `"errors"` to the import block since it was removed in Task 3.

- [ ] **Step 4: Update http.NewServer() to return error**

In `internal/http/server.go`, replace the `NewServer` function. Change from:

```go
func NewServer(
	prod bool,
	host string,
	port int,
	tlsDomain string,
	auth *auth.Auth,
	fetcher *fetcher.Fetcher,
	projection *projecter.Projecter,
) *Server {
	if auth == nil {
		panic("auth cannot be nil")
	}
	if fetcher == nil {
		panic("fetcher cannot be nil")
	}
	if projection == nil {
		panic("projection cannot be nil")
	}

	s := &Server{
		mux:       http.NewServeMux(),
		prod:      prod,
		host:      host,
		port:      port,
		tlsDomain: tlsDomain,
		auth:      auth,
		fetcher:   fetcher,
		projecter: projection,
		static:    static.FS,
	}

	s.initRoutes()

	return s
}
```

to:

```go
func NewServer(
	prod bool,
	host string,
	port int,
	tlsDomain string,
	auth *auth.Auth,
	fetcher *fetcher.Fetcher,
	projection *projecter.Projecter,
) (*Server, error) {
	if auth == nil {
		return nil, errors.New("auth cannot be nil")
	}
	if fetcher == nil {
		return nil, errors.New("fetcher cannot be nil")
	}
	if projection == nil {
		return nil, errors.New("projection cannot be nil")
	}

	s := &Server{
		mux:       http.NewServeMux(),
		prod:      prod,
		host:      host,
		port:      port,
		tlsDomain: tlsDomain,
		auth:      auth,
		fetcher:   fetcher,
		projecter: projection,
		static:    static.FS,
	}

	s.initRoutes()

	return s, nil
}
```

Note: add `"errors"` to the import block.

- [ ] **Step 5: Update main.go to handle all constructor errors**

Replace `cmd/serve/main.go` with:

```go
package main

import (
	"context"
	"log"

	"github.com/mgjules/mgjules-go/internal/auth"
	"github.com/mgjules/mgjules-go/internal/config"
	"github.com/mgjules/mgjules-go/internal/fetcher"
	"github.com/mgjules/mgjules-go/internal/http"
	"github.com/mgjules/mgjules-go/internal/logger"
	"github.com/mgjules/mgjules-go/internal/projecter"
	"github.com/mgjules/mgjules-go/internal/repository"
	"github.com/panjf2000/ants/v2"
)

//go:generate npm run build

func main() {
	cfg, err := config.Parse()
	if err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}

	logger.Init(cfg.Prod)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	repo, err := repository.New(ctx, cfg)
	if err != nil {
		log.Fatalf("failed to create repository: %v", err)
	}
	defer repo.Cleanup()

	pool, err := ants.NewPool(100)
	if err != nil {
		log.Fatalf("failed to create pool: %v", err)
	}
	defer pool.Release()

	f, err := fetcher.New(repo, pool)
	if err != nil {
		log.Fatalf("failed to create fetcher: %v", err)
	}
	f.Start()
	defer f.Stop()

	p, err := projecter.New(cfg.Prod, pool, f)
	if err != nil {
		log.Fatalf("failed to create projecter: %v", err)
	}
	f.AddSubscriber(p.Build)

	pool.Submit(f.Fetch)

	a, err := auth.New(cfg.AuthToken)
	if err != nil {
		log.Fatalf("failed to create auth: %v", err)
	}

	server, err := http.NewServer(
		cfg.Prod,
		cfg.ServerHost,
		cfg.ServerPort,
		cfg.ServerTLSDomain,
		a,
		f,
		p,
	)
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}
	if err = server.Start(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
```

- [ ] **Step 6: Verify it compiles**

```bash
go build ./...
```

Expected: success.

- [ ] **Step 7: Commit**

```bash
git add internal/auth/auth.go internal/fetcher/fetcher.go internal/projecter/projecter.go internal/http/server.go cmd/serve/main.go
git commit -m "Decouple projecter via DataProvider interface and replace panic with error returns"
```

---

### Task 7: Final verification and ARCHITECTURE.md update

**Files:**
- Modify: `ARCHITECTURE.md`

- [ ] **Step 1: Full build verification**

```bash
go build ./...
```

Expected: success.

- [ ] **Step 2: Run go vet**

```bash
go vet ./...
```

Expected: no issues.

- [ ] **Step 3: Update ARCHITECTURE.md**

In `ARCHITECTURE.md`, make these changes:

1. In the Repository section, remove the EdgeDB row from the table. Change from:

```
| Implementation | Description |
|---------------|-------------|
| directus/ | REST API client using resty, configurable via env |
| edgedb/ | Graph database client (deprecated, takes priority if DSN set) |
| static/ | Hardcoded fixture data for local dev without CMS |
```

to:

```
| Implementation | Description |
|---------------|-------------|
| directus/ | REST API client using resty, configurable via env |
| static/ | Hardcoded fixture data for local dev without CMS |
```

2. Update the heading text. Change from:

```
Three implementations with priority: EdgeDB > Directus > Static
```

to:

```
Two implementations with priority: Directus > Static
```

3. In the Dependencies table, remove the edgedb-go row if present.

4. In the ASCII pipeline diagram, no changes needed (it already shows `Repository (Directus/EdgeDB/Static)` — update to `Repository (Directus/Static)`). Change from:

```
Repository (Directus/EdgeDB/Static)
```

to:

```
Repository (Directus/Static)
```

5. In the Mermaid diagram, remove the EdgeDB node. Change from:

```
    subgraph Data Sources
        Directus["Directus CMS<br/>(REST API)"]
        EdgeDB["EdgeDB<br/>(deprecated)"]
        Static["Static Data<br/>(hardcoded)"]
    end
```

to:

```
    subgraph Data Sources
        Directus["Directus CMS<br/>(REST API)"]
        Static["Static Data<br/>(hardcoded)"]
    end
```

And remove the line:

```
    EdgeDB --> Repository
```

- [ ] **Step 4: Commit**

```bash
git add ARCHITECTURE.md
git commit -m "Update architecture docs to reflect EdgeDB removal"
```
