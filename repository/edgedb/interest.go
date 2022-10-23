package edgedb

import (
	"context"
	"fmt"

	"github.com/edgedb/edgedb-go"
	"github.com/mgjules/mgjules-go/entity"
)

type Interest struct {
	ID    edgedb.UUID `edgedb:"id"`
	Name  string      `edgedb:"name"`
	Image string      `edgedb:"image"`
}

func (i Interest) ToEntity() entity.Interest {
	return entity.Interest{
		ID:    i.ID.String(),
		Name:  i.Name,
		Image: i.Image,
	}
}

func (db *EdgeDB) GetInterests(ctx context.Context) ([]entity.Interest, error) {
	var results []Interest
	err := db.client.Query(ctx, `
		select CVInterest {
			id,
			name,
			image
		} order by .sort
	`, &results)
	if err != nil {
		return nil, fmt.Errorf("failed to query interests: %w", err)
	}

	interests := make([]entity.Interest, len(results))
	for i, result := range results {
		interests[i] = result.ToEntity()
	}

	return interests, nil
}
