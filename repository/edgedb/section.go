package edgedb

import (
	"context"
	"fmt"

	"github.com/edgedb/edgedb-go"
	"github.com/mgjules/mgjules-go/entity"
)

type Section struct {
	ID   edgedb.UUID `edgedb:"id"`
	Name string      `edgedb:"name"`
	Icon string      `edgedb:"icon"`
}

func (s Section) ToEntity() entity.Section {
	return entity.Section{
		ID:   s.ID.String(),
		Name: s.Name,
		Icon: s.Icon,
	}
}

func (db *EdgeDB) GetSections(ctx context.Context) ([]entity.Section, error) {
	var results []Section
	err := db.client.Query(ctx, `
		select CVSection {
			id,
			name,
			icon,
		} order by .sort
	`, &results)
	if err != nil {
		return nil, fmt.Errorf("failed to query sections: %w", err)
	}

	sections := make([]entity.Section, len(results))
	for i, result := range results {
		sections[i] = result.ToEntity()
	}

	return sections, nil
}
