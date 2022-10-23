package edgedb

import (
	"context"
	"fmt"

	"github.com/edgedb/edgedb-go"
	"github.com/mgjules/mgjules-go/entity"
)

type Project struct {
	ID           edgedb.UUID  `edgedb:"id"`
	Name         string       `edgedb:"name"`
	Description  string       `edgedb:"description"`
	Link         string       `edgedb:"link"`
	Technologies []Technology `edgedb:"technologies"`
}

func (p Project) ToEntity() entity.Project {
	technologies := make([]entity.Technology, len(p.Technologies))
	for i, techology := range p.Technologies {
		technologies[i] = techology.ToEntity()
	}

	return entity.Project{
		ID:           p.ID.String(),
		Name:         p.Name,
		Description:  p.Description,
		Link:         p.Link,
		Technologies: technologies,
	}
}

func (db *EdgeDB) GetProjects(ctx context.Context) ([]entity.Project, error) {
	var results []Project
	err := db.client.Query(ctx, `
		select CVProject {
			id,
			name,
			link,
			description,
			technologies: {
				name
			} order by @sort
		} order by .sort
	`, &results)
	if err != nil {
		return nil, fmt.Errorf("failed to query projects: %w", err)
	}

	projects := make([]entity.Project, len(results))
	for i, result := range results {
		projects[i] = result.ToEntity()
	}

	return projects, nil
}
