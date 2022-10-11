package edgedb

import (
	"context"
	"fmt"

	"github.com/edgedb/edgedb-go"
	"github.com/mgjules/mgjules-go/entity"
)

func (db *EdgeDB) GetMeta(ctx context.Context, id string) (*entity.Meta, error) {
	uuid, err := edgedb.ParseUUID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to parse id: %w", err)
	}

	var meta entity.Meta
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
	`, &meta, uuid)
	if err != nil {
		return nil, fmt.Errorf("failed to query single meta: %w", err)
	}

	return &meta, nil
}
