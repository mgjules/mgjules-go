package static

import (
	"context"

	"github.com/mgjules/mgjules-go/entity"
)

func (db *Static) GetLinks(ctx context.Context) ([]entity.Link, error) {
	return []entity.Link{
		{
			ID:   "home",
			Name: "Home",
			URL:  "/",
			Icon: "material-symbols:waving-hand-sharp",
		},
		{
			ID:   "blog",
			Name: "Blog",
			URL:  "/blog",
			Icon: "material-symbols:newspaper-sharp",
		},
		{
			ID:        "github",
			Name:      "Github",
			Icon:      "mdi:github",
			NewWindow: true,
		},
		{
			ID:        "linkedin",
			Name:      "LinkedIn",
			Icon:      "mdi:linkedin",
			NewWindow: true,
		},
	}, nil
}
