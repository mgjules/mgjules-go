package projection

import (
	"errors"
	"fmt"

	"github.com/mgjules/mgjules-go/entity"
	"github.com/mgjules/mgjules-go/mapstruct"
	"github.com/samber/lo"
)

func (p *Projection) BuildIndex() ([]byte, error) {
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

	indexCSS, err := p.parseSCSS("templates/index.scss")
	if err != nil {
		return nil, fmt.Errorf("failed to parse editor.scss: %w", err)
	}

	pCtx := map[string]any{
		"title":       p.meta.FirstName + " " + p.meta.LastName + " - " + currentTab.Name + "." + currentTab.Extension,
		"current_tab": mapstruct.FromSingle(currentTab),
		"tabs":        mapstruct.FromSlice(tabs),
		"intro":       mapstruct.FromSingle(p.intro),
		"print":       false,
		"indexCSS":    indexCSS,
	}

	out, err := p.render(pCtx, "Home", "templates/index.html")
	if err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}

	return out, nil
}
