package directus

import (
	"context"
	"fmt"
	"net/url"

	"github.com/mgjules/mgjules-go/entity"
)

type Link struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	URL          string  `json:"url"`
	Icon         string  `json:"icon"`
	AlternateURL *string `json:"alternate_url"`
	NewWindow    bool    `json:"new_window"`
}

func (l Link) ToEntity() entity.Link {
	return entity.Link{
		ID:           l.ID,
		Name:         l.Name,
		URL:          l.URL,
		Icon:         l.Icon,
		AlternateURL: l.AlternateURL,
		NewWindow:    l.NewWindow,
	}
}

func (db *Directus) GetLinks(ctx context.Context) ([]entity.Link, error) {
	var result Result[[]Link]
	resp, err := db.client.R().
		SetContext(ctx).
		SetQueryParamsFromValues(url.Values{
			"fields": []string{
				"id",
				"name",
				"url",
				"icon",
				"alternate_url",
				"new_window",
			},
			"status": []string{"published"},
			"sort":   []string{"sort"},
		}).
		SetResult(&result).
		Get("/items/link")
	if err != nil {
		return nil, fmt.Errorf("failed to get links: %w", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("failed to get links: response code %d", resp.StatusCode())
	}

	links := make([]entity.Link, len(result.Data))
	for i, link := range result.Data {
		links[i] = link.ToEntity()
	}

	return links, nil
}
