package edgedb

import (
	"context"
	"fmt"

	"github.com/mgjules/mgjules-go/entity"
)

func (db *EdgeDB) GetContributions(ctx context.Context) ([]entity.Contribution, error) {
	var contributions []entity.Contribution
	err := db.client.Query(ctx, `
		select CVContribution {
			id,
			event,
			title,
			role,
			from,
			to,
			link
		} order by .from desc
	`, &contributions)
	if err != nil {
		return nil, fmt.Errorf("failed to query contributions: %w", err)
	}

	return contributions, nil
}
