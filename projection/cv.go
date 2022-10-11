package projection

import (
	"errors"
	"fmt"

	"github.com/mgjules/mgjules-go/entity"
	"github.com/mgjules/mgjules-go/mapstruct"
	"github.com/samber/lo"
)

func (p *Projection) BuildCV(section string) ([]byte, error) {
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

	indexCSS, err := p.parseSCSS("templates/index.scss")
	if err != nil {
		return nil, fmt.Errorf("failed to parse editor.scss: %w", err)
	}

	values := map[string]any{
		"title":       p.meta.FullName + " - " + currentTab.Name + "." + currentTab.Extension,
		"current_tab": mapstruct.FromSingle(currentTab),
		"intro":       mapstruct.FromSingle(p.intro),
		"index_css":   indexCSS,
	}

	out, err := p.render(values, "Home", "templates/index.dhtml")
	if err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}

	return out, nil
}
