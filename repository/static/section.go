package static

import (
	"context"

	"github.com/mgjules/mgjules-go/entity"
)

func (db *Static) GetSections(ctx context.Context) ([]entity.Section, error) {
	return []entity.Section{
		{
			ID:   "experiences",
			Name: "Work Experiences",
			Icon: "material-symbols:work-sharp",
		},
		{
			ID:   "projects",
			Name: "Pinned Projects",
			Icon: "material-symbols:work",
		},
	}, nil
}
