package directus

import (
	"context"
	"fmt"
	"log/slog"
	"net/url"
	"time"

	"github.com/araddon/dateparse"
	"github.com/mgjules/mgjules-go/entity"
)

type Contribution struct {
	ID    string  `json:"id"`
	Event string  `json:"event"`
	Title string  `json:"title"`
	From  string  `json:"from"`
	To    *string `json:"to"`
	Link  string  `json:"link"`
	Role  string  `json:"role"`
}

func (c Contribution) ToEntity() entity.Contribution {
	from, err := dateparse.ParseAny(c.From)
	if err != nil {
		slog.Debug("failed to parse `from`", "id", c.ID, "error", err)
	}

	var to *time.Time
	if c.To != nil {
		if parsed, err := dateparse.ParseAny(*c.To); err != nil {
			slog.Debug("failed to parse `to`", "id", c.ID, "error", err)
		} else {
			to = &parsed
		}
	}

	return entity.Contribution{
		ID:    c.ID,
		Event: c.Event,
		Title: c.Title,
		From:  from,
		To:    to,
		Link:  c.Link,
		Role:  c.Role,
	}
}

func (db *Directus) GetContributions(ctx context.Context) ([]entity.Contribution, error) {
	var result Result[[]Contribution]
	resp, err := db.client.R().
		SetContext(ctx).
		SetQueryParamsFromValues(url.Values{
			"fields": []string{
				"id",
				"event",
				"title",
				"link",
				"from",
				"to",
				"role",
			},
			"status": []string{"published"},
			"sort":   []string{"-from"},
		}).
		SetResult(&result).
		Get("/items/contribution")
	if err != nil {
		return nil, fmt.Errorf("failed to get contributions: %w", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("failed to get contributions: response code %d", resp.StatusCode())
	}

	contributions := make([]entity.Contribution, len(result.Data))
	for i, contribution := range result.Data {
		contributions[i] = contribution.ToEntity()
	}

	return contributions, nil
}
