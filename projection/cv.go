package projection

import (
	"fmt"
	"strings"

	"github.com/mgjules/mgjules-go/entity"
	"github.com/mgjules/mgjules-go/mapstruct"
)

func (p *Projection) BuildCV(section entity.Section) ([]byte, error) {
	var tabs []entity.Tab
	for _, section := range p.sections {
		tabs = append(tabs, entity.Tab{
			Name:      section.Name,
			Icon:      section.Icon,
			Extension: "cv",
			URL:       "/cv/" + strings.ToLower(section.Name),
		})
	}

	currentTab := entity.Tab{
		Name:      section.Name,
		Icon:      section.Icon,
		Extension: "cv",
		URL:       "/cv/" + strings.ToLower(section.Name),
	}

	cvCSS, err := p.parseSCSS("templates/cv/" + strings.ToLower(section.Name) + ".scss")
	if err != nil {
		return nil, fmt.Errorf("failed to parse editor.scss: %w", err)
	}

	values := map[string]any{
		"title":       p.meta.FullName + " - " + currentTab.Name + "." + currentTab.Extension,
		"tabs":        mapstruct.FromSlice(tabs),
		"current_tab": mapstruct.FromSingle(currentTab),
		"intro":       mapstruct.FromSingle(p.intro),
		"cv_css":      cvCSS,
	}

	out, err := p.render(values, "Curriculum Vitae", "templates/cv/"+strings.ToLower(section.Name)+".dhtml")
	if err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}

	return out, nil
}
