package directus

import (
	"context"
	"fmt"
	"net/url"

	"github.com/mgjules/mgjules-go/internal/entity"
)

type Meta struct {
	ID       string   `json:"id"`
	BaseURL  string   `json:"base_url"`
	Lang     string   `json:"lang"`
	Keywords []string `json:"keywords"`
	User     User     `json:"user"`
}

func (m Meta) ToEntity(directusURL string) entity.Meta {
	return entity.Meta{
		ID:          m.ID,
		BaseURL:     m.BaseURL,
		Lang:        m.Lang,
		Keywords:    m.Keywords,
		Description: m.User.Description,
		FirstName:   m.User.FirstName,
		LastName:    m.User.LastName,
		FullName:    m.User.FirstName + " " + m.User.LastName,
		Github:      m.User.Github,
		Username:    m.User.Username,
		Gender:      m.User.Gender,
		Avatar:      directusURL + "/assets/" + m.User.Avatar + "/avatar.webp?key=meta",
	}
}

func (db *Directus) GetMeta(ctx context.Context, id string) (*entity.Meta, error) {
	var result Result[Meta]
	resp, err := db.client.R().
		SetContext(ctx).
		SetQueryParamsFromValues(url.Values{
			"fields": []string{
				"id",
				"base_url",
				"lang",
				"keywords",
				"user.first_name",
				"user.last_name",
				"user.description",
				"user.avatar",
				"user.github",
				"user.username",
				"user.gender",
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
