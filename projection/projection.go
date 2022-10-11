package projection

import (
	"context"
	"embed"
	"strings"
	"sync"
	"time"

	"github.com/flosch/pongo2/v6"
	"github.com/mgjules/mgjules-go/entity"
	"github.com/mgjules/mgjules-go/logger"
	"github.com/mgjules/mgjules-go/repository"
	loader "github.com/nathan-osman/pongo2-embed-loader"
	"github.com/robfig/cron/v3"
)

const (
	seperator string = ":"
)

type Projection struct {
	prod bool

	dataMu sync.Mutex // guards the data
	meta   entity.Meta
	links  []entity.Link
	intro  entity.Introduction

	projectionsMu sync.Mutex // guards the projections
	projections   map[string][]byte

	repo        repository.Repository
	templates   embed.FS
	templateSet *pongo2.TemplateSet
	cron        *cron.Cron
}

func New(prod bool, repo repository.Repository, templates embed.FS) *Projection {
	p := &Projection{
		prod:        prod,
		repo:        repo,
		cron:        cron.New(),
		projections: make(map[string][]byte),
		templates:   templates,
	}

	if p.prod {
		p.templateSet = pongo2.NewSet("", &loader.Loader{Content: templates})
	} else {
		p.templateSet = pongo2.NewSet("", pongo2.MustNewLocalFileSystemLoader("./"))
	}

	p.cron.AddFunc("@hourly", func() {
		p.FetchData()
		p.BuildProjections()
	})

	return p
}

func (p *Projection) Start() {
	p.cron.Start()
}

func (p *Projection) Stop() {
	p.cron.Stop()
}

func (p *Projection) FetchData() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	p.dataMu.Lock()
	defer p.dataMu.Unlock()

	logger.Logger.Debug("Fetching meta data...")
	meta, err := p.repo.GetMeta(ctx, "bd99e066-440b-11ed-924c-9fd15527df84")
	if err != nil {
		logger.Logger.Errorf("failed to get meta: %v", err)
	} else {
		p.meta = *meta
	}

	logger.Logger.Debug("Fetching links data...")
	links, err := p.repo.GetLinks(ctx)
	if err != nil {
		logger.Logger.Errorf("failed to get links: %v", err)
	} else {
		p.links = links
	}

	logger.Logger.Debug("Fetching introduction data...")
	intro, err := p.repo.GetIntroduction(ctx, "a4296eac-441b-11ed-924c-830c8fd1144c")
	if err != nil {
		logger.Logger.Errorf("failed to get introduction: %v", err)
	} else {
		p.intro = *intro
	}
}

func (p *Projection) BuildProjections() {
	logger.Logger.Debug("Building index projection...")
	index, err := p.BuildIndex()
	if err != nil {
		logger.Logger.Errorf("failed to build index projection: %v", err)
	} else {
		p.projectionsMu.Lock()
		p.projections[buildKey("index")] = index
		p.projectionsMu.Unlock()
	}
}

func (p *Projection) Get(keys ...string) ([]byte, bool) {
	if len(keys) == 0 {
		return nil, false
	}

	if !p.prod {
		p.BuildProjections()
	}

	out, found := p.projections[buildKey(keys...)]
	return out, found
}

func buildKey(keys ...string) string {
	return strings.Join(keys, seperator)
}
