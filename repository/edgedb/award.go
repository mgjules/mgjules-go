package edgedb

import (
	"context"
	"fmt"

	"github.com/mgjules/mgjules-go/entity"
)

func (db *EdgeDB) GetAwards(ctx context.Context) ([]entity.Award, error) {
	var awards []entity.Award
	err := db.client.Query(ctx, `
		select CVAward {
			id,
			event,
			description,
			result,
			date,
			link,
			icon
		} order by .date desc
	`, &awards)
	if err != nil {
		return nil, fmt.Errorf("failed to query awards: %w", err)
	}

	return awards, nil
}
