package projecter

import (
	"fmt"

	"github.com/mgjules/mgjules-go/entity"
	"github.com/mgjules/mgjules-go/mapstruct"
)

func (p *Projecter) Build404() ([]byte, error) {
	currentTab := entity.Tab{
		Name:      "Not Found",
		Icon:      "ooui:article-not-found-ltr",
		Extension: "404",
	}

	tabs := []entity.Tab{
		currentTab,
	}

	values := map[string]any{
		"title":       p.fetcher.Meta().FullName + " - " + currentTab.Name + "." + currentTab.Extension,
		"tabs":        mapstruct.FromSlice(tabs),
		"current_tab": mapstruct.FromSingle(currentTab),
	}

	out, err := p.render(values, "", "templates/404.dhtml")
	if err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}

	return out, nil
}
