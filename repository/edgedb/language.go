package edgedb

import (
	"context"
	"fmt"

	"github.com/mgjules/mgjules-go/entity"
)

func (db *EdgeDB) GetLanguages(ctx context.Context) ([]entity.Language, error) {
	var languages []entity.Language
	err := db.client.Query(ctx, `
		select CVLanguage {
			id,
			name,
			icon,
			level
		} order by .sort
	`, &languages)
	if err != nil {
		return nil, fmt.Errorf("failed to query languages: %w", err)
	}

	return languages, nil
}
