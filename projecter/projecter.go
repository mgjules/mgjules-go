package projecter

import (
	"bytes"
	"compress/gzip"
	"context"
	"log/slog"
	"strings"
	"sync"

	"github.com/a-h/templ"
	"github.com/mgjules/mgjules-go/fetcher"
	"github.com/mgjules/mgjules-go/templates/minimal"
	"github.com/panjf2000/ants/v2"
)

const (
	seperator string = ":"
)

type Projecter struct {
	prod    bool
	pool    *ants.Pool
	fetcher *fetcher.Fetcher

	mu          sync.RWMutex // guards the projections
	projections map[string][]byte
}

func New(
	prod bool,
	pool *ants.Pool,
	fetcher *fetcher.Fetcher,
) *Projecter {
	if pool == nil {
		panic("pool cannot be empty")
	}
	if fetcher == nil {
		panic("fetcher cannot be nil")
	}

	return &Projecter{
		prod:        prod,
		pool:        pool,
		fetcher:     fetcher,
		projections: make(map[string][]byte),
	}
}

func (p *Projecter) Build() {
	meta := p.fetcher.Meta()
	links := p.fetcher.Links()
	intro := p.fetcher.Introduction()
	sections := p.fetcher.Sections()
	experiences := p.fetcher.Experiences()
	projects := p.fetcher.Projects()
	// contributions := p.fetcher.Contributions()
	// awards := p.fetcher.Awards()
	// interests := p.fetcher.Interests()
	// languages := p.fetcher.Languages()
	// posts := p.fetcher.Posts()

	ctx := context.Background()
	b := &bytes.Buffer{}
	w, err := gzip.NewWriterLevel(b, gzip.BestCompression)
	if err != nil {
		slog.Error("failed to create new gzip writer", "error", err)
		return
	}
	defer w.Close()

	p.render(ctx, w, b, buildKey("index"), minimal.Index(meta, links, intro, sections, experiences, projects))
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
		w.Flush()
		p.mu.Lock()
		p.projections[key] = b.Bytes()
		p.mu.Unlock()
		w.Reset(b)
	}
}

func buildKey(keys ...string) string {
	return strings.Join(keys, seperator)
}
