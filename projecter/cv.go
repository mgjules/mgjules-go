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

	cvCSS, err := p.parseSCSS("cv/" + strings.ToLower(section.Name) + ".scss")
	if err != nil {
		return nil, fmt.Errorf("failed to parse %s.scss: %w", strings.ToLower(section.Name), err)
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

	out, err := p.render(meta, links, "Curriculum Vitae", "cv/"+strings.ToLower(section.Name)+".dhtml", values)
	if err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}

	return out, nil
}

func (p *Projecter) BuildCVPrint(
	meta *entity.Meta,
	links []entity.Link,
	sections []entity.Section,
	introduction *entity.Introduction,
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
		Name:      "Print",
		Icon:      "ic:baseline-local-printshop",
		Extension: "cv",
		URL:       "/cv/print",
	}

	indexCSS, err := p.parseSCSS("index.scss")
	if err != nil {
		return nil, fmt.Errorf("failed to parse index.scss: %w", err)
	}

	if len(projects) > 5 {
		projects = projects[:5]
	}

	if len(contributions) > 5 {
		contributions = contributions[:5]
	}

	values := map[string]any{
		"title":         meta.FullName + " - " + currentTab.Name + "." + currentTab.Extension,
		"tabs":          mapstruct.FromSlice(tabs),
		"current_tab":   mapstruct.FromSingle(currentTab),
		"intro":         mapstruct.FromSingle(introduction),
		"experiences":   mapstruct.FromSlice(experiences),
		"projects":      mapstruct.FromSlice(projects),
		"contributions": mapstruct.FromSlice(contributions),
		"awards":        mapstruct.FromSlice(awards),
		"interests":     mapstruct.FromSlice(interests),
		"languages":     mapstruct.FromSlice(languages),
		"index_css":     indexCSS,
	}

	for _, section := range sections {
		cvCSS, err := p.parseSCSS("cv/" + strings.ToLower(section.Name) + ".scss")
		if err != nil {
			continue
		}

		values[strings.ToLower(section.Name)+"_css"] = cvCSS
	}

	out, err := p.render(meta, links, "Curriculum Vitae", "cv/print.dhtml", values)
	if err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}

	return out, nil
}
