package static

import (
	"context"

	"github.com/mgjules/mgjules-go/entity"
)

func (db *Static) GetMeta(ctx context.Context, id string) (*entity.Meta, error) {
	return &entity.Meta{
		BaseURL:     "https://mgjules.dev",
		Lang:        "en",
		Description: "Professional Software Engineer",
		FirstName:   "Michael Giovanni",
		LastName:    "Jules",
		FullName:    "Michael Giovanni Jules",
		Keywords: []string{
			"Go", "Developer", "Software Engineer", "Backend",
		},
		Github:   "https://github.com/mgjules",
		Username: "mgjules",
		Gender:   "male",
		Avatar:   "img/avatar.webp",
	}, nil
}
