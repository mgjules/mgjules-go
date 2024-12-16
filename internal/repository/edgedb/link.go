package edgedb

import (
	"context"
	"fmt"

	"github.com/edgedb/edgedb-go"
	"github.com/mgjules/mgjules-go/internal/entity"
)

type Link struct {
	ID           edgedb.UUID        `edgedb:"id"`
	Name         string             `edgedb:"name"`
	URL          string             `edgedb:"url"`
	Icon         string             `edgedb:"icon"`
	AlternateURL edgedb.OptionalStr `edgedb:"alternate_url"`
	NewWindow    bool               `edgedb:"new_window"`
}

func (l Link) ToEntity() entity.Link {
	var alterateURL *string
	if val, ok := l.AlternateURL.Get(); ok {
		alterateURL = &val
	}

	return entity.Link{
		ID:           l.ID.String(),
		Name:         l.Name,
		URL:          l.URL,
		Icon:         l.Icon,
		AlternateURL: alterateURL,
		NewWindow:    l.NewWindow,
	}
}

func (db *edgeDB) GetLinks(ctx context.Context) ([]entity.Link, error) {
	var results []Link
	err := db.client.Query(ctx, `
		select SiteLink {
			id,
			name,
			url,
			alternate_url,
			new_window,
			icon
		} order by .sort
	`, &results)
	if err != nil {
		return nil, fmt.Errorf("failed to query links: %w", err)
	}

	links := make([]entity.Link, len(results))
	for i, result := range results {
		links[i] = result.ToEntity()
	}

	return links, nil
}
