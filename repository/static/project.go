package static

import (
	"context"

	"github.com/mgjules/mgjules-go/entity"
)

func (db *Static) GetProjects(ctx context.Context) ([]entity.Project, error) {
	return []entity.Project{
		{
			Name:         "Flo",
			Description:  "A flow-based programming thingy in Go.",
			Link:         "https://github.com/mgjules/flo",
			Technologies: getTechnologies("Go"),
		},
	}, nil
}
