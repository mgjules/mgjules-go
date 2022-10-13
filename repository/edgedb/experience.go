package edgedb

import (
	"context"
	"fmt"

	"github.com/mgjules/mgjules-go/entity"
)

func (db *EdgeDB) GetExperiences(ctx context.Context) ([]entity.Experience, error) {
	var experiences []entity.Experience
	err := db.client.Query(ctx, `
		select CVExperience {
			id,
			company,
			position,
			from,
			to,
			link,
			tasks,
			technologies: {
				name
			} order by @sort
		} order by .from desc
	`, &experiences)
	if err != nil {
		return nil, fmt.Errorf("failed to query experiences: %w", err)
	}

	return experiences, nil
}
