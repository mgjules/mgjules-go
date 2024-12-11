package projecter

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/bep/godartsass/v2"
	"github.com/flosch/pongo2/v6"
	"github.com/mgjules/mgjules-go/fetcher"
	"github.com/mgjules/mgjules-go/logger"
	loader "github.com/nathan-osman/pongo2-embed-loader"
	"github.com/panjf2000/ants/v2"
)

const (
	seperator string = ":"
)

type Projecter struct {
	prod        bool
	pool        *ants.Pool
	fetcher     *fetcher.Fetcher
	templates   fs.FS
	templateSet *pongo2.TemplateSet
	transpiler  *godartsass.Transpiler

	projectionsMu sync.RWMutex // guards the projections
	projections   map[string][]byte
	projectedAt   time.Time
}

func New(
	prod bool,
	theme string,
	pool *ants.Pool,
	fetcher *fetcher.Fetcher,
	templates fs.FS,
	transpiler *godartsass.Transpiler,
) (*Projecter, error) {
	p := &Projecter{
		prod:        prod,
		pool:        pool,
		fetcher:     fetcher,
		projections: make(map[string][]byte),
		transpiler:  transpiler,
	}

	if p.prod {
		templates, err := fs.Sub(templates, "templates/"+theme)
		if err != nil {
			return nil, fmt.Errorf("failed to get sub fs: %w", err)
		}
		p.templates = templates
	} else {
		p.templates = os.DirFS("./templates/" + theme)
	}

	p.templateSet = pongo2.NewSet("", &loader.Loader{Content: p.templates})

	return p, nil
}

func (p *Projecter) Build() {
	var wg sync.WaitGroup

	p.projectionsMu.Lock()
	p.projectedAt = time.Now()
	p.projectionsMu.Unlock()

	meta := p.fetcher.Meta()
	links := p.fetcher.Links()
	sections := p.fetcher.Sections()
	introduction := p.fetcher.Introduction()
	experiences := p.fetcher.Experiences()
	projects := p.fetcher.Projects()
	contributions := p.fetcher.Contributions()
	awards := p.fetcher.Awards()
	interests := p.fetcher.Interests()
	languages := p.fetcher.Languages()
	posts := p.fetcher.Posts()

	wg.Add(1)
	p.pool.Submit(func() {
		defer wg.Done()

		logger.L.Debug("Building index projection...")
		index, err := p.BuildIndex(&meta, links, &introduction)
		if err != nil {
			logger.L.Errorf("failed to build index projection: %v", err)
		} else {
			p.projectionsMu.Lock()
			p.projections[buildKey("index")] = index
			p.projectionsMu.Unlock()
		}
	})

	for _, section := range sections {
		section := section

		wg.Add(1)
		p.pool.Submit(func() {
			defer wg.Done()

			logger.L.Debugf("Building cv '%s' projection...", section.Name)
			cv, err := p.BuildCV(
				&meta,
				links,
				sections,
				&section,
				experiences,
				projects,
				contributions,
				awards,
				interests,
				languages,
			)
			if err != nil {
				logger.L.Errorf("failed to build cv '%s' projection: %v", section.Name, err)
			} else {
				p.projectionsMu.Lock()
				p.projections[buildKey("cv", strings.ToLower(section.Name))] = cv
				p.projectionsMu.Unlock()
			}
		})
	}

	wg.Add(1)
	p.pool.Submit(func() {
		defer wg.Done()

		logger.L.Debugf("Building cv print projection...")
		cvPrint, err := p.BuildCVPrint(
			&meta,
			links,
			sections,
			&introduction,
			experiences,
			projects,
			contributions,
			awards,
			interests,
			languages,
		)
		if err != nil {
			logger.L.Errorf("failed to build cv print projection: %v", err)
		} else {
			p.projectionsMu.Lock()
			p.projections[buildKey("cv", "print")] = cvPrint
			p.projectionsMu.Unlock()
		}
	})

	wg.Add(1)
	p.pool.Submit(func() {
		defer wg.Done()

		logger.L.Debug("Building blog index projection...")
		blogIndex, err := p.BuildBlogIndex(&meta, links, posts)
		if err != nil {
			logger.L.Errorf("failed to build blog index projection: %v", err)
		} else {
			p.projectionsMu.Lock()
			p.projections[buildKey("blog", "index")] = blogIndex
			p.projectionsMu.Unlock()
		}
	})

	for _, post := range posts {
		post := post

		wg.Add(1)
		p.pool.Submit(func() {
			defer wg.Done()

			logger.L.Debugf("Building blog '%s' projection...", post.Slug)
			blogPost, err := p.BuildBlogPost(&meta, links, posts, &post)
			if err != nil {
				logger.L.Errorf("failed to build blog '%s' projection: %v", post.Slug, err)
			} else {
				p.projectionsMu.Lock()
				p.projections[buildKey("blog", strings.ToLower(post.Slug))] = blogPost
				p.projectionsMu.Unlock()
			}
		})
	}

	wg.Add(1)
	p.pool.Submit(func() {
		defer wg.Done()

		logger.L.Debug("Building 404 projection...")
		notFound, err := p.Build404(&meta, links)
		if err != nil {
			logger.L.Errorf("failed to build 404 projection: %v", err)
		} else {
			p.projectionsMu.Lock()
			p.projections[buildKey("404")] = notFound
			p.projectionsMu.Unlock()
		}
	})

	wg.Wait()
}

func (p *Projecter) Get(keys ...string) ([]byte, bool) {
	if len(keys) == 0 {
		return nil, false
	}

	if !p.prod {
		p.Build()
	}

	p.projectionsMu.RLock()
	out, found := p.projections[buildKey(keys...)]
	p.projectionsMu.RUnlock()

	return out, found
}

func buildKey(keys ...string) string {
	return strings.Join(keys, seperator)
}
