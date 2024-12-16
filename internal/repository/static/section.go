package static

import (
	"context"

	"github.com/mgjules/mgjules-go/internal/entity"
)

func (db *Static) GetSections(ctx context.Context) ([]entity.Section, error) {
	return []entity.Section{
		{
			ID:   "experiences",
			Name: "Professional Experiences",
			Icon: "material-symbols:work-sharp",
		},
		{
			ID:   "projects",
			Name: "Pinned Personal Projects",
			Icon: "material-symbols:work",
		},
	}, nil
}
