package projecter

import (
	"fmt"

	"github.com/mgjules/mgjules-go/entity"
	"github.com/mgjules/mgjules-go/mapstruct"
)

func (p *Projecter) Build404(meta *entity.Meta, links []entity.Link) ([]byte, error) {
	currentTab := entity.Tab{
		Name:      "Not Found",
		Icon:      "ooui:article-not-found-ltr",
		Extension: "404",
	}

	tabs := []entity.Tab{
		currentTab,
	}

	values := map[string]any{
		"title":       meta.FullName + " - " + currentTab.Name + "." + currentTab.Extension,
		"tabs":        mapstruct.FromSlice(tabs),
		"current_tab": mapstruct.FromSingle(currentTab),
	}

	out, err := p.render(meta, links, "", "templates/404.dhtml", values)
	if err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}

	return out, nil
}
