package projecter

import (
	"fmt"
	"strings"

	"github.com/mgjules/mgjules-go/entity"
	"github.com/mgjules/mgjules-go/mapstruct"
)

func (p *Projecter) BuildCV(section entity.Section) ([]byte, error) {
	var tabs []entity.Tab
	for _, section := range p.fetcher.Sections() {
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
		"title":       p.fetcher.Meta().FullName + " - " + currentTab.Name + "." + currentTab.Extension,
		"tabs":        mapstruct.FromSlice(tabs),
		"current_tab": mapstruct.FromSingle(currentTab),
		"cv_css":      cvCSS,
	}

	switch section.Name {
	case "Experiences":
		values["experiences"] = mapstruct.FromSlice(p.fetcher.Experiences())
	case "Projects":
		values["projects"] = mapstruct.FromSlice(p.fetcher.Projects())
	case "Contributions":
		values["contributions"] = mapstruct.FromSlice(p.fetcher.Contributions())
	case "Awards":
		values["awards"] = mapstruct.FromSlice(p.fetcher.Awards())
	case "Interests":
		values["interests"] = mapstruct.FromSlice(p.fetcher.Interests())
	case "Languages":
		values["languages"] = mapstruct.FromSlice(p.fetcher.Languages())
	}

	out, err := p.render(values, "Curriculum Vitae", "templates/cv/"+strings.ToLower(section.Name)+".dhtml")
	if err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}

	return out, nil
}
