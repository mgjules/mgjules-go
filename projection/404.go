package projection

import (
	"fmt"

	"github.com/mgjules/mgjules-go/entity"
	"github.com/mgjules/mgjules-go/mapstruct"
)

func (p *Projection) Build404() ([]byte, error) {
	currentTab := entity.Tab{
		Name:      "Not Found",
		Icon:      "ooui:article-not-found-ltr",
		Extension: "404",
	}

	tabs := []entity.Tab{
		currentTab,
	}

	pCtx := map[string]any{
		"title":       p.meta.FullName + " - " + currentTab.Name + "." + currentTab.Extension,
		"tabs":        mapstruct.FromSlice(tabs),
		"current_tab": mapstruct.FromSingle(currentTab),
	}

	out, err := p.render(pCtx, "", "templates/404.dhtml")
	if err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}

	return out, nil
}
