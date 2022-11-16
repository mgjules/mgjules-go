package projecter

import (
	"errors"
	"fmt"

	"github.com/mgjules/mgjules-go/entity"
	"github.com/mgjules/mgjules-go/mapstruct"
	"github.com/samber/lo"
)

func (p *Projecter) BuildIndex(meta *entity.Meta, links []entity.Link, introduction *entity.Introduction) ([]byte, error) {
	link, found := lo.Find(links, func(link entity.Link) bool {
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

	values := map[string]any{
		"title":       meta.FullName + " - " + currentTab.Name + "." + currentTab.Extension,
		"tabs":        mapstruct.FromSlice(tabs),
		"current_tab": mapstruct.FromSingle(currentTab),
		"intro":       mapstruct.FromSingle(introduction),
		"index_css":   indexCSS,
	}

	out, err := p.render(meta, links, "Home", "templates/index.dhtml", values)
	if err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}

	return out, nil
}
