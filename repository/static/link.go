package static

import (
	"context"

	"github.com/mgjules/mgjules-go/entity"
)

func (db *Static) GetLinks(ctx context.Context) ([]entity.Link, error) {
	return []entity.Link{
		{
			Name: "Home",
			URL:  "/",
			Icon: "material-symbols:waving-hand-sharp",
		},
		{
			Name: "Blog",
			URL:  "/blog",
			Icon: "material-symbols:newspaper-sharp",
		},
		{
			Name:      "Github (mgjules)",
			URL:       "https://github.com/mgjules",
			Icon:      "mdi:github",
			NewWindow: true,
		},
		{
			Name:      "LinkedIn (mgjules)",
			URL:       "https://linkedin.com/in/mgjules",
			Icon:      "mdi:linkedin",
			NewWindow: true,
		},
	}, nil
}
