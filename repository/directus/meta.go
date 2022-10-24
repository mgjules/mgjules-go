package directus

import (
	"context"
	"fmt"
	"net/url"

	"github.com/mgjules/mgjules-go/entity"
)

type Meta struct {
	ID          string   `json:"id"`
	BaseURL     string   `json:"base_url"`
	Lang        string   `json:"lang"`
	Description string   `json:"description"`
	FirstName   string   `json:"first_name"`
	LastName    string   `json:"last_name"`
	Keywords    []string `json:"keywords"`
	Github      string   `json:"github"`
	Username    string   `json:"username"`
	Gender      string   `json:"gender"`
	Avatar      string   `json:"avatar"`
}

func (m Meta) ToEntity(directusURL string) entity.Meta {
	return entity.Meta{
		ID:          m.ID,
		BaseURL:     m.BaseURL,
		Lang:        m.Lang,
		Description: m.Description,
		FirstName:   m.FirstName,
		LastName:    m.LastName,
		FullName:    m.FirstName + " " + m.LastName,
		Keywords:    m.Keywords,
		Github:      m.Github,
		Username:    m.Username,
		Gender:      m.Gender,
		Avatar:      directusURL + "/assets/" + m.Avatar + "/avatar.webp?key=meta",
	}
}

func (db *Directus) GetMeta(ctx context.Context, id string) (*entity.Meta, error) {
	var result Result[Meta]
	resp, err := db.client.R().
		SetQueryParamsFromValues(url.Values{
			"fields": []string{
				"base_url",
				"lang",
				"description",
				"first_name",
				"last_name",
				"keywords",
				"github",
				"username",
				"gender",
				"avatar",
			},
			"limit": []string{"1"},
		}).
		SetResult(&result).
		Get("/items/meta")
	if err != nil {
		return nil, fmt.Errorf("failed to get meta: %w", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("failed to get meta: response code %d", resp.StatusCode())
	}

	meta := result.Data.ToEntity(db.directusURL)

	return &meta, nil
}
