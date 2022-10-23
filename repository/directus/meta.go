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
	FullName    string   `json:"full_name"`
	Keywords    []string `json:"keywords"`
	Github      string   `json:"github"`
	Username    string   `json:"username"`
	Gender      string   `json:"gender"`
	Avatar      string   `json:"avatar"`
}

func (m Meta) ToEntity() entity.Meta {
	return entity.Meta{
		ID:          m.ID,
		BaseURL:     m.BaseURL,
		Lang:        m.Lang,
		Description: m.Description,
		FirstName:   m.FirstName,
		LastName:    m.LastName,
		FullName:    m.FullName,
		Keywords:    m.Keywords,
		Github:      m.Github,
		Username:    m.Username,
		Gender:      m.Gender,
		Avatar:      m.Avatar,
	}
}

func (db *Directus) GetMeta(ctx context.Context, id string) (*entity.Meta, error) {
	var result Result[Meta]
	resp, err := db.client.R().
		SetQueryParamsFromValues(url.Values{
			"fields": []string{
				"base_url",
				"first_name",
				"last_name",
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

	meta := result.Data.ToEntity()

	return &meta, nil
}
