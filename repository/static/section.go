package static

import (
	"context"

	"github.com/mgjules/mgjules-go/entity"
)

func (db *Static) GetSections(ctx context.Context) ([]entity.Section, error) {
	return []entity.Section{
		{
			Name: "Experiences",
			Icon: "material-symbols:work-sharp",
		},
		{
			Name: "Projects",
			Icon: "material-symbols:work",
		},
	}, nil
}
