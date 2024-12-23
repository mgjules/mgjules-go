package edgedb

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/araddon/dateparse"
	"github.com/edgedb/edgedb-go"
	"github.com/mgjules/mgjules-go/internal/entity"
)

type Experience struct {
	ID           edgedb.UUID              `edgedb:"id"`
	Company      string                   `edgedb:"company"`
	Position     string                   `edgedb:"position"`
	From         edgedb.LocalDate         `edgedb:"from"`
	To           edgedb.OptionalLocalDate `edgedb:"to"`
	Link         string                   `edgedb:"link"`
	Technologies []Technology             `edgedb:"technologies"`
	Tasks        []string                 `edgedb:"tasks"`
}

func (e Experience) ToEntity() entity.Experience {
	from, err := dateparse.ParseAny(e.From.String())
	if err != nil {
		slog.Debug("failed to parse `from`", "id", e.ID, "error", err)
	}

	var to *time.Time
	if val, ok := e.To.Get(); ok {
		if parsed, err := dateparse.ParseAny(val.String()); err != nil {
			slog.Debug("failed to parse `to`", "id", e.ID, "error", err)
		} else {
			to = &parsed
		}
	}

	technologies := make([]entity.Technology, len(e.Technologies))
	for i, techology := range e.Technologies {
		technologies[i] = techology.ToEntity()
	}

	return entity.Experience{
		ID:           e.ID.String(),
		Company:      e.Company,
		Position:     e.Position,
		From:         from,
		To:           to,
		Link:         e.Link,
		Technologies: technologies,
		Tasks:        e.Tasks,
	}
}

func (db *edgeDB) GetExperiences(ctx context.Context) ([]entity.Experience, error) {
	var results []Experience
	err := db.client.Query(ctx, `
		select CVExperience {
			id,
			company,
			position,
			from,
			to,
			link,
			tasks,
			technologies: {
				name,
				link
			} order by @sort
		} order by .from desc
	`, &results)
	if err != nil {
		return nil, fmt.Errorf("failed to query experiences: %w", err)
	}

	experiences := make([]entity.Experience, len(results))
	for i, result := range results {
		experiences[i] = result.ToEntity()
	}

	return experiences, nil
}
