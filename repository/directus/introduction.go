package directus

import (
	"context"
	"fmt"
	"net/url"

	"github.com/mgjules/mgjules-go/entity"
)

type Introduction struct {
	ID           string `json:"id"`
	Introduction string `json:"introduction"`
	Avatar       string `json:"avatar"`
}

func (i Introduction) ToEntity(directusURL string) entity.Introduction {
	return entity.Introduction{
		ID:           i.ID,
		Introduction: i.Introduction,
		Avatar:       directusURL + "/assets/" + i.Avatar + "/avatar.webp?key=intro",
	}
}

func (db *Directus) GetIntroduction(ctx context.Context, id string) (*entity.Introduction, error) {
	var result Result[Introduction]
	resp, err := db.client.R().
		SetQueryParamsFromValues(url.Values{
			"fields": []string{
				"id",
				"introduction",
				"avatar",
			},
			"limit": []string{"1"},
		}).
		SetResult(&result).
		Get("/items/introduction")
	if err != nil {
		return nil, fmt.Errorf("failed to get introduction: %w", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("failed to get introduction: response code %d", resp.StatusCode())
	}

	intro := result.Data.ToEntity(db.directusURL)

	return &intro, nil
}
