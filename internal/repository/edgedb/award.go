package edgedb

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/araddon/dateparse"
	"github.com/edgedb/edgedb-go"
	"github.com/mgjules/mgjules-go/internal/entity"
)

type Award struct {
	ID          edgedb.UUID      `edgedb:"id"`
	Event       string           `edgedb:"event"`
	Description string           `edgedb:"description"`
	Date        edgedb.LocalDate `edgedb:"date"`
	Link        string           `edgedb:"link"`
	Result      string           `edgedb:"result"`
	Icon        string           `edgedb:"icon"`
}

func (a Award) ToEntity() entity.Award {
	date, err := dateparse.ParseAny(a.Date.String())
	if err != nil {
		slog.Debug("failed to parse `date`", "id", a.ID, "error", err)
	}

	return entity.Award{
		ID:          a.ID.String(),
		Event:       a.Event,
		Description: a.Description,
		Date:        date,
		Link:        a.Link,
		Result:      a.Result,
		Icon:        a.Icon,
	}
}

func (db *edgeDB) GetAwards(ctx context.Context) ([]entity.Award, error) {
	var results []Award
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
	`, &results)
	if err != nil {
		return nil, fmt.Errorf("failed to query awards: %w", err)
	}

	awards := make([]entity.Award, len(results))
	for i, result := range results {
		awards[i] = result.ToEntity()
	}

	return awards, nil
}
