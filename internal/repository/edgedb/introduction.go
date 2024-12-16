package edgedb

import (
	"context"
	"fmt"

	"github.com/edgedb/edgedb-go"
	"github.com/mgjules/mgjules-go/internal/entity"
)

type Introduction struct {
	ID           edgedb.UUID `edgedb:"id"`
	Introduction string      `edgedb:"introduction"`
	Avatar       string      `edgedb:"avatar"`
}

func (i Introduction) ToEntity() entity.Introduction {
	return entity.Introduction{
		ID:           i.ID.String(),
		Introduction: i.Introduction,
		Avatar:       i.Avatar,
	}
}

func (db *edgeDB) GetIntroduction(ctx context.Context, id string) (*entity.Introduction, error) {
	uuid, err := edgedb.ParseUUID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to parse id: %w", err)
	}

	var result Introduction
	err = db.client.QuerySingle(ctx, `
		select Introduction {
			id,
			introduction,
			avatar
		} filter .id = <uuid>$0
	`, &result, uuid)
	if err != nil {
		return nil, fmt.Errorf("failed to query single introduction: %w", err)
	}

	intro := result.ToEntity()

	return &intro, nil
}
