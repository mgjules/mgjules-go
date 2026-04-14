# Architecture Cleanup Design

## Goal

Clean up the project architecture by removing dead code, reducing repetition in the fetcher, and improving structural boundaries — without rearchitecting what already works.

## Scope

Three areas of work, ordered by risk (lowest first).

---

## 1. Dead Code Removal

### Delete entirely

| Path | Reason |
|------|--------|
| `internal/repository/edgedb/` (all 15 files) | Deprecated EdgeDB repository implementation |
| `internal/db/edgedb/` (schema + 28 migrations) | EdgeDB schema and migrations |
| `edgedb.toml` | EdgeDB project config |
| `internal/fetcher/copier.go` | DeepCopy helper, no longer needed (see Section 2) |

### Config changes

- Remove `EdgeDBDSN` field from `internal/config/config.go`
- Simplify `repository.New()` — remove EdgeDB branch, keep Directus > Static fallback

Resulting factory:

```go
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

### Commented-out code

- `internal/http/server.go:89-92` — remove commented-out route registrations
- `internal/projecter/projecter.go:61-64` — remove commented-out fetcher calls (contributions, awards, interests, languages)

### Dependencies removed from go.mod

Direct:
- `github.com/edgedb/edgedb-go`
- `github.com/jinzhu/copier`

Note: `github.com/samber/lo` is still used in templates and static repository — it stays.

Transitive deps will be cleaned up automatically by `go mod tidy`.

---

## 2. Fetcher Deduplication

### Problem

`Fetcher.Fetch()` is ~180 lines of 11 copy-pasted goroutine blocks that differ only in entity name, repo method, and storage field.

### Solution

Introduce a generic `fetchEntity` helper:

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

`Fetch()` becomes a series of `fetchEntity` calls:

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
    // backoff retry logic unchanged
}
```

### Error handling change

Current: `errors.Join` collects all errors, returns them, triggering a full retry of all 11 entities.

New: Per-entity log-and-skip. A failed entity keeps its previous cached value. No retry of the full batch for partial failures. This is better behavior — if "awards" fails, no reason to re-fetch all 11.

The backoff retry is kept but only triggers on context errors (timeout), not partial entity failures.

### DeepCopy removal

Entities are plain structs. The only reference types are slices (`[]entity.Link`, `[]string`, `[]Technology`). The fetcher replaces fields atomically under a write lock — it never mutates existing slice contents. The projecter reads all data at once under read locks via `Build()`. There is no concurrent mutation scenario that requires deep copying.

Accessor methods become direct returns under read lock:

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
```

Remove `internal/fetcher/copier.go` and the `jinzhu/copier` dependency.

---

## 3. Structural Improvements

### 3a. Constructor error handling

Replace `panic` with returned errors in all constructors:

| Package | Function | Change |
|---------|----------|--------|
| `repository` | `New()` | Return `(Repository, error)` instead of panicking |
| `fetcher` | `New()` | Return `(*Fetcher, error)` |
| `projecter` | `New()` | Return `(*Projecter, error)` |
| `http` | `NewServer()` | Return `(*Server, error)` |
| `auth` | `New()` | Return `(*Auth, error)` |

`main.go` handles all errors with `log.Fatalf` consistently.

### 3b. Projecter interface decoupling

The projecter currently imports `*fetcher.Fetcher` concretely. Define a `DataProvider` interface in the projecter package:

```go
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
```

`Projecter.New()` takes `DataProvider` instead of `*fetcher.Fetcher`. The fetcher already satisfies this interface — no changes needed there.

### 3c. Fetcher notify() lock fix

Current `notify()` reads `f.subscribers` without holding `subscribersMu`. Fix:

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

---

## Out of Scope

- Not splitting fetcher into DataStore + Syncer
- Not adding an App lifecycle struct
- Not touching templates, entities, or Directus/Static repository implementations
- Not adding tests (covered by separate test design spec)

## Files Modified

| File | Change |
|------|--------|
| `internal/config/config.go` | Remove `EdgeDBDSN` field |
| `internal/repository/repository.go` | Remove EdgeDB branch, return error |
| `internal/fetcher/fetcher.go` | Generic fetchEntity, remove DeepCopy usage, fix notify lock |
| `internal/fetcher/copier.go` | Delete |
| `internal/projecter/projecter.go` | DataProvider interface, return error from New, remove commented code |
| `internal/http/server.go` | Return error from NewServer, remove commented routes |
| `internal/auth/auth.go` | Return error from New |
| `cmd/serve/main.go` | Handle all constructor errors, remove copier/edgedb imports |
| `go.mod` / `go.sum` | Remove unused dependencies |
| `edgedb.toml` | Delete |
| `internal/repository/edgedb/` | Delete directory |
| `internal/db/edgedb/` | Delete directory |
