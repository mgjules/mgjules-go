package edgedb

import (
	"context"
	"fmt"

	"github.com/mgjules/mgjules-go/entity"
)

func (db *EdgeDB) GetLinks(ctx context.Context) ([]entity.Link, error) {
	var links []entity.Link
	err := db.client.Query(ctx, `
		select SiteLink {
			id,
			name,
			url,
			alternate_url,
			new_window,
			icon
		} order by .sort
	`, &links)
	if err != nil {
		return nil, fmt.Errorf("failed to query links: %w", err)
	}

	return links, nil
}
