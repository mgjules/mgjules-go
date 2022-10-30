package directus

import (
	"context"
	"fmt"
	"net/url"

	"github.com/mgjules/mgjules-go/entity"
)

type Project struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Link         string `json:"link"`
	Technologies []struct {
		Technology Technology `json:"technology"`
	} `json:"technologies"`
}

func (p Project) ToEntity() entity.Project {
	technologies := make([]entity.Technology, len(p.Technologies))
	for i, technology := range p.Technologies {
		technologies[i] = technology.Technology.ToEntity()
	}

	return entity.Project{
		ID:           p.ID,
		Name:         p.Name,
		Description:  p.Description,
		Link:         p.Link,
		Technologies: technologies,
	}
}

func (db *Directus) GetProjects(ctx context.Context) ([]entity.Project, error) {
	var result Result[[]Project]
	resp, err := db.client.R().
		SetContext(ctx).
		SetQueryParamsFromValues(url.Values{
			"fields": []string{
				"id",
				"name",
				"description",
				"link",
				"technologies.technology.id",
				"technologies.technology.name",
				"technologies.technology.link",
			},
			"status": []string{"published"},
			"sort":   []string{"sort"},
		}).
		SetResult(&result).
		Get("/items/project")
	if err != nil {
		return nil, fmt.Errorf("failed to get projects: %w", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("failed to get projects: response code %d", resp.StatusCode())
	}

	projects := make([]entity.Project, len(result.Data))
	for i, project := range result.Data {
		projects[i] = project.ToEntity()
	}

	return projects, nil
}
