package directus

import (
	"context"
	"fmt"
	"net/url"

	"github.com/avelino/slugify"
	"github.com/mgjules/mgjules-go/entity"
)

type Interest struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

func (i Interest) ToEntity(directusURL string) entity.Interest {
	return entity.Interest{
		ID:    i.ID,
		Name:  i.Name,
		Image: directusURL + "/assets/" + i.Image + "/" + slugify.Slugify(i.Name) + ".webp?key=interest",
	}
}

func (db *Directus) GetInterests(ctx context.Context) ([]entity.Interest, error) {
	var result Result[[]Interest]
	resp, err := db.client.R().
		SetQueryParamsFromValues(url.Values{
			"fields": []string{
				"id",
				"name",
				"image",
			},
			"status": []string{"published"},
			"sort":   []string{"sort"},
		}).
		SetResult(&result).
		Get("/items/interest")
	if err != nil {
		return nil, fmt.Errorf("failed to get interests: %w", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("failed to get interests: response code %d", resp.StatusCode())
	}

	interests := make([]entity.Interest, len(result.Data))
	for i, interest := range result.Data {
		interests[i] = interest.ToEntity(db.directusURL)
	}

	return interests, nil
}
