package directus

import (
	"context"
	"fmt"
	"net/url"

	"github.com/mgjules/mgjules-go/internal/entity"
)

type Language struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Icon  string `json:"icon"`
	Level string `json:"level"`
}

func (l Language) ToEntity() entity.Language {
	return entity.Language{
		ID:    l.ID,
		Name:  l.Name,
		Icon:  l.Icon,
		Level: l.Level,
	}
}

func (db *directus) GetLanguages(ctx context.Context) ([]entity.Language, error) {
	var result Result[[]Language]
	resp, err := db.client.R().
		SetContext(ctx).
		SetQueryParamsFromValues(url.Values{
			"fields": []string{
				"id",
				"name",
				"icon",
				"level",
			},
			"status": []string{"published"},
			"sort":   []string{"sort"},
		}).
		SetResult(&result).
		Get("/items/language")
	if err != nil {
		return nil, fmt.Errorf("failed to get languages: %w", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("failed to get languages: response code %d", resp.StatusCode())
	}

	languages := make([]entity.Language, len(result.Data))
	for i, language := range result.Data {
		languages[i] = language.ToEntity()
	}

	return languages, nil
}
