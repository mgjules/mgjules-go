package projection

import (
	"bytes"
	"errors"
	"fmt"
	"time"

	"github.com/flosch/pongo2/v6"
	"github.com/mgjules/mgjules-go/entity"
	"github.com/mgjules/mgjules-go/mapstruct"
	"github.com/samber/lo"
	"github.com/wellington/go-libsass"
)

func (p *Projection) BuildIndex() ([]byte, error) {
	tpl, err := p.templateSet.FromFile("templates/index.html")
	if err != nil {
		return nil, fmt.Errorf("failed to load templates from file: %w", err)
	}

	link, found := lo.Find(p.links, func(link entity.Link) bool {
		return link.Name == "Home"
	})
	if !found {
		return nil, errors.New("missing current link")
	}

	currentTab := entity.Tab{
		Name:      link.Name,
		Icon:      link.Icon,
		Extension: "index",
	}

	tabs := []entity.Tab{
		currentTab,
	}

	links := lo.Map(p.links, func(link entity.Link, _ int) entity.Link {
		if link.Name == "Home" {
			link.IsCurrent = true
		}

		return link
	})

	scss, err := p.templates.Open("templates/layouts/editor.scss")
	if err != nil {
		return nil, fmt.Errorf("failed to open editor.scss: %w", err)
	}

	buf := bytes.NewBuffer(nil)
	comp, err := libsass.New(buf, scss)
	if err != nil {
		return nil, fmt.Errorf("failed to create libsass compiler: %w", err)
	}

	if err := comp.Run(); err != nil {
		return nil, fmt.Errorf("failed to compile editor.scss: %w", err)
	}

	p.dataMu.RLock()
	pCtx := pongo2.Context{
		"title":        p.meta.FirstName + " " + p.meta.LastName + " - " + currentTab.Name + "." + currentTab.Extension,
		"meta":         mapstruct.FromSingle(p.meta),
		"current_tab":  mapstruct.FromSingle(currentTab),
		"tabs":         mapstruct.FromSlice(tabs),
		"links":        mapstruct.FromSlice(links),
		"intro":        mapstruct.FromSingle(p.intro),
		"print":        false,
		"css":          buf.String(),
		"current_year": time.Now().Year(),
	}
	p.dataMu.RUnlock()

	out, err := tpl.ExecuteBytes(pCtx)
	if err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}

	return out, nil
}
