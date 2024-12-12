package static

import (
	"context"
	"time"

	"github.com/mgjules/mgjules-go/entity"
)

func (db *Static) GetExperiences(ctx context.Context) ([]entity.Experience, error) {
	return []entity.Experience{
		{
			Company:      "Livestorm",
			Position:     "Senior Software Engineer",
			From:         time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC),
			Link:         "https://livestorm.co",
			Technologies: getTechnologies("Go"),
			Tasks: []string{
				`Something *markdown*`, "Task 2",
			},
		},
		{
			Company:      "Atellio",
			Position:     "Senior Software Engineer",
			From:         time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC),
			Link:         "https://livestorm.co",
			Technologies: getTechnologies("Go"),
			Tasks:        []string{"Task 1", "Task 2"},
		},
	}, nil
}
