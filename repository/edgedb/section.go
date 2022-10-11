package edgedb

import (
	"context"
	"fmt"

	"github.com/mgjules/mgjules-go/entity"
)

func (db *EdgeDB) GetSections(ctx context.Context) ([]entity.Section, error) {
	var sections []entity.Section
	err := db.client.Query(ctx, `
		select CVSection {
			id,
			name,
			icon,
		} order by .sort
	`, &sections)
	if err != nil {
		return nil, fmt.Errorf("failed to query sections: %w", err)
	}

	return sections, nil
}
