package directus

import (
	"context"
	"fmt"
	"log/slog"
	"net/url"

	"github.com/araddon/dateparse"
	"github.com/mgjules/mgjules-go/internal/entity"
)

type Award struct {
	ID          string `json:"id"`
	Event       string `json:"event"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Link        string `json:"link"`
	Result      string `json:"result"`
	Icon        string `json:"icon"`
}

func (a Award) ToEntity() entity.Award {
	date, err := dateparse.ParseAny(a.Date)
	if err != nil {
		slog.Debug("failed to parse `date`", "id", a.ID, "error", err)
	}

	return entity.Award{
		ID:          a.ID,
		Event:       a.Event,
		Description: a.Description,
		Date:        date,
		Link:        a.Link,
		Result:      a.Result,
		Icon:        a.Icon,
	}
}

func (db *directus) GetAwards(ctx context.Context) ([]entity.Award, error) {
	var result Result[[]Award]
	resp, err := db.client.R().
		SetContext(ctx).
		SetQueryParamsFromValues(url.Values{
			"fields": []string{
				"id",
				"event",
				"description",
				"date",
				"link",
				"result",
				"icon",
			},
			"status": []string{"published"},
			"sort":   []string{"-date"},
		}).
		SetResult(&result).
		Get("/items/award")
	if err != nil {
		return nil, fmt.Errorf("failed to get awards: %w", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("failed to get awards: response code %d", resp.StatusCode())
	}

	awards := make([]entity.Award, len(result.Data))
	for i, award := range result.Data {
		awards[i] = award.ToEntity()
	}

	return awards, nil
}
