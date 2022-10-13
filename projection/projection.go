package projection

import (
	"context"
	"embed"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/flosch/pongo2/v6"
	"github.com/mgjules/mgjules-go/entity"
	"github.com/mgjules/mgjules-go/logger"
	"github.com/mgjules/mgjules-go/repository"
	loader "github.com/nathan-osman/pongo2-embed-loader"
	"github.com/panjf2000/ants/v2"
	"github.com/robfig/cron/v3"
)

const (
	seperator string = ":"
)

type Projection struct {
	prod        bool
	pool        *ants.Pool
	repo        repository.Repository
	templates   embed.FS
	templateSet *pongo2.TemplateSet
	cron        *cron.Cron

	dataMu      sync.RWMutex // guards the data
	meta        entity.Meta
	links       []entity.Link
	intro       entity.Introduction
	sections    []entity.Section
	experiences []entity.Experience

	projectionsMu sync.RWMutex // guards the projections
	projections   map[string][]byte
}

func New(prod bool, repo repository.Repository, templates embed.FS) (*Projection, error) {
	p := &Projection{
		prod:        prod,
		repo:        repo,
		cron:        cron.New(),
		projections: make(map[string][]byte),
		templates:   templates,
	}

	pool, err := ants.NewPool(1000)
	if err != nil {
		return nil, fmt.Errorf("failed to create pool: %w", err)
	}
	p.pool = pool

	if p.prod {
		p.templateSet = pongo2.NewSet("", &loader.Loader{Content: templates})
	} else {
		p.templateSet = pongo2.NewSet("", pongo2.MustNewLocalFileSystemLoader("./"))
	}

	p.cron.AddFunc("@hourly", func() {
		p.FetchData()
		p.BuildProjections()
	})

	return p, nil
}

func (p *Projection) Start() {
	p.cron.Start()
}

func (p *Projection) Stop() {
	p.cron.Stop()
	p.pool.Release()
}

func (p *Projection) FetchData() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	p.pool.Submit(func() {
		defer wg.Done()

		logger.L.Debug("Fetching meta data...")
		meta, err := p.repo.GetMeta(ctx, "bd99e066-440b-11ed-924c-9fd15527df84")
		if err != nil {
			logger.L.Errorf("failed to get meta: %v", err)
		} else {
			p.dataMu.Lock()
			p.meta = *meta
			p.dataMu.Unlock()
		}
	})

	wg.Add(1)
	p.pool.Submit(func() {
		defer wg.Done()

		logger.L.Debug("Fetching links data...")
		links, err := p.repo.GetLinks(ctx)
		if err != nil {
			logger.L.Errorf("failed to get links: %v", err)
		} else {
			p.dataMu.Lock()
			p.links = links
			p.dataMu.Unlock()
		}
	})

	wg.Add(1)
	p.pool.Submit(func() {
		defer wg.Done()

		logger.L.Debug("Fetching introduction data...")
		intro, err := p.repo.GetIntroduction(ctx, "a4296eac-441b-11ed-924c-830c8fd1144c")
		if err != nil {
			logger.L.Errorf("failed to get introduction: %v", err)
		} else {
			p.dataMu.Lock()
			p.intro = *intro
			p.dataMu.Unlock()
		}
	})

	wg.Add(1)
	p.pool.Submit(func() {
		defer wg.Done()

		logger.L.Debug("Fetching section data...")
		sections, err := p.repo.GetSections(ctx)
		if err != nil {
			logger.L.Errorf("failed to get sections: %v", err)
		} else {
			p.dataMu.Lock()
			p.sections = sections
			p.dataMu.Unlock()
		}
	})

	wg.Add(1)
	p.pool.Submit(func() {
		defer wg.Done()

		logger.L.Debug("Fetching experience data...")
		experiences, err := p.repo.GetExperiences(ctx)
		if err != nil {
			logger.L.Errorf("failed to get experiences: %v", err)
		} else {
			p.dataMu.Lock()
			p.experiences = experiences
			p.dataMu.Unlock()
		}
	})

	wg.Wait()
}

func (p *Projection) BuildProjections() {
	var wg sync.WaitGroup

	wg.Add(1)
	p.pool.Submit(func() {
		defer wg.Done()

		logger.L.Debug("Building index projection...")
		p.dataMu.RLock()
		index, err := p.BuildIndex()
		p.dataMu.RUnlock()
		if err != nil {
			logger.L.Errorf("failed to build index projection: %v", err)
		} else {
			p.projectionsMu.Lock()
			p.projections[buildKey("index")] = index
			p.projectionsMu.Unlock()
		}
	})

	p.dataMu.RLock()
	sections := p.sections
	p.dataMu.RUnlock()
	for _, section := range sections {
		section := section

		wg.Add(1)
		p.pool.Submit(func() {
			defer wg.Done()

			logger.L.Debugf("Building cv '%s' projection...", section.Name)
			p.dataMu.RLock()
			cv, err := p.BuildCV(section)
			p.dataMu.RUnlock()
			if err != nil {
				logger.L.Errorf("failed to build cv '%s' projection: %v", section.Name, err)
			} else {
				p.projectionsMu.Lock()
				p.projections[buildKey("cv", strings.ToLower(section.Name))] = cv
				p.projectionsMu.Unlock()
			}
		})
	}

	wg.Wait()
}

func (p *Projection) Get(keys ...string) ([]byte, bool) {
	if len(keys) == 0 {
		return nil, false
	}

	if !p.prod {
		p.BuildProjections()
	}

	p.projectionsMu.RLock()
	out, found := p.projections[buildKey(keys...)]
	p.projectionsMu.RUnlock()

	return out, found
}

func buildKey(keys ...string) string {
	return strings.Join(keys, seperator)
}
