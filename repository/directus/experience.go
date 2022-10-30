package directus

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/araddon/dateparse"
	"github.com/mgjules/mgjules-go/entity"
	"github.com/mgjules/mgjules-go/logger"
)

type Experience struct {
	ID           string  `json:"id"`
	Company      string  `json:"company"`
	Position     string  `json:"position"`
	From         string  `json:"from"`
	To           *string `json:"to"`
	Link         string  `json:"link"`
	Technologies []struct {
		Technology Technology `json:"technology"`
	} `json:"technologies"`
	Tasks []struct {
		Task string `json:"task"`
	} `json:"tasks"`
}

func (e Experience) ToEntity() entity.Experience {
	from, err := dateparse.ParseAny(e.From)
	if err != nil {
		logger.L.Debugf("failed to parse `from` for `%s`: %v", e.ID, err)
	}

	var to *time.Time
	if e.To != nil {
		if parsed, err := dateparse.ParseAny(*e.To); err != nil {
			logger.L.Debugf("failed to parse `to` for `%s`: %v", e.ID, err)
		} else {
			to = &parsed
		}
	}

	technologies := make([]entity.Technology, len(e.Technologies))
	for i, technology := range e.Technologies {
		technologies[i] = technology.Technology.ToEntity()
	}

	tasks := make([]string, len(e.Tasks))
	for i, task := range e.Tasks {
		tasks[i] = task.Task
	}

	return entity.Experience{
		ID:           e.ID,
		Company:      e.Company,
		Position:     e.Position,
		From:         from,
		To:           to,
		Link:         e.Link,
		Technologies: technologies,
		Tasks:        tasks,
	}
}

func (db *Directus) GetExperiences(ctx context.Context) ([]entity.Experience, error) {
	var result Result[[]Experience]
	resp, err := db.client.R().
		SetContext(ctx).
		SetQueryParamsFromValues(url.Values{
			"fields": []string{
				"id",
				"company",
				"position",
				"from",
				"to",
				"link",
				"technologies.technology.id",
				"technologies.technology.name",
				"technologies.technology.link",
				"tasks",
			},
			"status": []string{"published"},
			"sort":   []string{"-from"},
		}).
		SetResult(&result).
		Get("/items/experience")
	if err != nil {
		return nil, fmt.Errorf("failed to get experiences: %w", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("failed to get experiences: response code %d", resp.StatusCode())
	}

	experiences := make([]entity.Experience, len(result.Data))
	for i, experience := range result.Data {
		experiences[i] = experience.ToEntity()
	}

	return experiences, nil
}
