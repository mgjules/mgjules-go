package edgedb

import (
	"context"
	"fmt"

	"github.com/mgjules/mgjules-go/entity"
)

func (db *EdgeDB) GetInterests(ctx context.Context) ([]entity.Interest, error) {
	var interests []entity.Interest
	err := db.client.Query(ctx, `
		select CVInterest {
			id,
			name,
			image
		} order by .sort
	`, &interests)
	if err != nil {
		return nil, fmt.Errorf("failed to query interests: %w", err)
	}

	return interests, nil
}
