package edgedb

import (
	"context"
	"fmt"

	"github.com/mgjules/mgjules-go/entity"
)

func (db *EdgeDB) GetProjects(ctx context.Context) ([]entity.Project, error) {
	var projects []entity.Project
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
	`, &projects)
	if err != nil {
		return nil, fmt.Errorf("failed to query projects: %w", err)
	}

	return projects, nil
}
