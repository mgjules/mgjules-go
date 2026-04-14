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
