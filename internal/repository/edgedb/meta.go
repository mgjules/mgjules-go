package edgedb

import (
	"context"
	"fmt"

	"github.com/edgedb/edgedb-go"
	"github.com/mgjules/mgjules-go/internal/entity"
)

type Meta struct {
	ID          edgedb.UUID `edgedb:"id"`
	BaseURL     string      `edgedb:"base_url"`
	Lang        string      `edgedb:"lang"`
	Description string      `edgedb:"description"`
	FirstName   string      `edgedb:"first_name"`
	LastName    string      `edgedb:"last_name"`
	FullName    string      `edgedb:"full_name"`
	Keywords    []string    `edgedb:"keywords"`
	Github      string      `edgedb:"github"`
	Username    string      `edgedb:"username"`
	Gender      string      `edgedb:"gender"`
	Avatar      string      `edgedb:"avatar"`
}

func (m Meta) ToEntity() entity.Meta {
	return entity.Meta{
		ID:          m.ID.String(),
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

func (db *EdgeDB) GetMeta(ctx context.Context, id string) (*entity.Meta, error) {
	uuid, err := edgedb.ParseUUID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to parse id: %w", err)
	}

	var result Meta
	err = db.client.QuerySingle(ctx, `
		select Meta {
			id,
			base_url,
			lang,
			first_name,
			last_name,
			full_name := .first_name ++ ' ' ++ .last_name,
			gender,
			description,
			keywords,
			github,
			username,
			avatar
		} filter .id = <uuid>$0
	`, &result, uuid)
	if err != nil {
		return nil, fmt.Errorf("failed to query single meta: %w", err)
	}

	meta := result.ToEntity()

	return &meta, nil
}
