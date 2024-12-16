package directus

import (
	"context"
	"fmt"
	"net/url"

	"github.com/mgjules/mgjules-go/internal/entity"
)

type Section struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

func (s Section) ToEntity() entity.Section {
	return entity.Section{
		ID:   s.ID,
		Name: s.Name,
		Icon: s.Icon,
	}
}

func (db *directus) GetSections(ctx context.Context) ([]entity.Section, error) {
	var result Result[[]Section]
	resp, err := db.client.R().
		SetContext(ctx).
		SetQueryParamsFromValues(url.Values{
			"fields": []string{
				"id",
				"name",
				"icon",
			},
			"status": []string{"published"},
			"sort":   []string{"sort"},
		}).
		SetResult(&result).
		Get("/items/section")
	if err != nil {
		return nil, fmt.Errorf("failed to get sections: %w", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("failed to get sections: response code %d", resp.StatusCode())
	}

	sections := make([]entity.Section, len(result.Data))
	for i, section := range result.Data {
		sections[i] = section.ToEntity()
	}

	return sections, nil
}
