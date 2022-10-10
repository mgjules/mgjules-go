package edgedb

import (
	"context"
	"fmt"

	"github.com/edgedb/edgedb-go"
	"github.com/mgjules/mgjules-go/entity"
)

func (db *EdgeDB) GetIntroduction(ctx context.Context, id string) (*entity.Introduction, error) {
	uuid, err := edgedb.ParseUUID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to parse id: %w", err)
	}

	var intro entity.Introduction
	err = db.client.QuerySingle(ctx, `
		select Introduction {
			id,
			introduction,
			avatar
		} filter .id = <uuid>$0
	`, &intro, uuid)
	if err != nil {
		return nil, fmt.Errorf("failed to query single introduction: %w", err)
	}

	return &intro, nil
}
