package edgedb

import (
	"context"
	"fmt"
	"time"

	"github.com/araddon/dateparse"
	"github.com/edgedb/edgedb-go"
	"github.com/mgjules/mgjules-go/entity"
	"github.com/mgjules/mgjules-go/logger"
)

type Contribution struct {
	ID    edgedb.UUID              `edgedb:"id"`
	Event string                   `edgedb:"event"`
	Title string                   `edgedb:"title"`
	From  edgedb.LocalDate         `edgedb:"from"`
	To    edgedb.OptionalLocalDate `edgedb:"to"`
	Link  string                   `edgedb:"link"`
	Role  string                   `edgedb:"role"`
}

func (c Contribution) ToEntity() entity.Contribution {
	from, err := dateparse.ParseAny(c.From.String())
	if err != nil {
		logger.L.Debugf("failed to parse `from` for `%s`: %v", c.ID, err)
	}

	var to *time.Time
	if val, ok := c.To.Get(); ok {
		if parsed, err := dateparse.ParseAny(val.String()); err != nil {
			logger.L.Debugf("failed to parse `to` for `%s`: %v", c.ID, err)
		} else {
			to = &parsed
		}
	}

	return entity.Contribution{
		ID:    c.ID.String(),
		Event: c.Event,
		Title: c.Title,
		From:  from,
		To:    to,
		Link:  c.Link,
		Role:  c.Role,
	}
}

func (db *EdgeDB) GetContributions(ctx context.Context) ([]entity.Contribution, error) {
	var results []Contribution
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
	`, &results)
	if err != nil {
		return nil, fmt.Errorf("failed to query contributions: %w", err)
	}

	contributions := make([]entity.Contribution, len(results))
	for i, result := range results {
		contributions[i] = result.ToEntity()
	}

	return contributions, nil
}
