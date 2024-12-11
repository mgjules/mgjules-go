package static

import (
	"context"

	"github.com/mgjules/mgjules-go/entity"
	"github.com/samber/lo"
)

func (db *Static) GetLinks(ctx context.Context) ([]entity.Link, error) {
	return []entity.Link{
		{
			Name: "Home",
			URL:  "/",
			Icon: "material-symbols:waving-hand-sharp",
		},
		{
			Name:         "Curriculum Vitae",
			URL:          "/cv",
			AlternateURL: lo.ToPtr("/cv/experiences"),
			Icon:         "material-symbols:lab-profile-sharp",
		},
		{
			Name: "Blog",
			URL:  "/blog",
			Icon: "material-symbols:newspaper-sharp",
		},
		{
			Name:      "Github",
			URL:       "https://github.com/mgjules",
			Icon:      "mdi:github",
			NewWindow: true,
		},
	}, nil
}
