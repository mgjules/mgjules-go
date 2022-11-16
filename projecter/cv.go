package projecter

import (
	"fmt"
	"strings"

	"github.com/mgjules/mgjules-go/entity"
	"github.com/mgjules/mgjules-go/mapstruct"
)

func (p *Projecter) BuildCV(
	meta *entity.Meta,
	links []entity.Link,
	sections []entity.Section,
	section *entity.Section,
	experiences []entity.Experience,
	projects []entity.Project,
	contributions []entity.Contribution,
	awards []entity.Award,
	interests []entity.Interest,
	languages []entity.Language,
) ([]byte, error) {
	var tabs []entity.Tab
	for _, section := range sections {
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
		"title":       meta.FullName + " - " + currentTab.Name + "." + currentTab.Extension,
		"tabs":        mapstruct.FromSlice(tabs),
		"current_tab": mapstruct.FromSingle(currentTab),
		"cv_css":      cvCSS,
	}

	switch section.Name {
	case "Experiences":
		values["experiences"] = mapstruct.FromSlice(experiences)
	case "Projects":
		values["projects"] = mapstruct.FromSlice(projects)
	case "Contributions":
		values["contributions"] = mapstruct.FromSlice(contributions)
	case "Awards":
		values["awards"] = mapstruct.FromSlice(awards)
	case "Interests":
		values["interests"] = mapstruct.FromSlice(interests)
	case "Languages":
		values["languages"] = mapstruct.FromSlice(languages)
	}

	out, err := p.render(meta, links, "Curriculum Vitae", "templates/cv/"+strings.ToLower(section.Name)+".dhtml", values)
	if err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}

	return out, nil
}
