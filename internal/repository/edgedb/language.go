package edgedb

import (
	"context"
	"fmt"

	"github.com/edgedb/edgedb-go"
	"github.com/mgjules/mgjules-go/internal/entity"
)

type Language struct {
	ID    edgedb.UUID `edgedb:"id"`
	Name  string      `edgedb:"name"`
	Icon  string      `edgedb:"icon"`
	Level string      `edgedb:"level"`
}

func (l Language) ToEntity() entity.Language {
	return entity.Language{
		ID:    l.ID.String(),
		Name:  l.Name,
		Icon:  l.Icon,
		Level: l.Level,
	}
}

func (db *EdgeDB) GetLanguages(ctx context.Context) ([]entity.Language, error) {
	var results []Language
	err := db.client.Query(ctx, `
		select CVLanguage {
			id,
			name,
			icon,
			level
		} order by .sort
	`, &results)
	if err != nil {
		return nil, fmt.Errorf("failed to query languages: %w", err)
	}

	languages := make([]entity.Language, len(results))
	for i, result := range results {
		languages[i] = result.ToEntity()
	}

	return languages, nil
}
