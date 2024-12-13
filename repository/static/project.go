package static

import (
	"context"

	"github.com/mgjules/mgjules-go/entity"
)

func (db *Static) GetProjects(ctx context.Context) ([]entity.Project, error) {
	return []entity.Project{
		{
			ID:           "harvit",
			Name:         "Harvit",
			Description:  "Harvit harvests data from different sources (e.g websites, APIs), converts and transforms it.",
			Link:         "https://github.com/mgjules/harvit",
			Technologies: getTechnologies("Go", "Docker"),
		},
		{
			ID:           "hap",
			Name:         "Hap",
			Description:  "hap, like in what's happening, is a generic event system aimed towards simplicity and performance.",
			Link:         "https://github.com/mgjules/hap",
			Technologies: getTechnologies("Go"),
		},
		{
			ID:           "flo",
			Name:         "Flo",
			Description:  "A flow-based programming thingy in Go. Can generate valid Go code from a node-based graph. Highly experimental.",
			Link:         "https://github.com/mgjules/flo",
			Technologies: getTechnologies("Go"),
		},
	}, nil
}
