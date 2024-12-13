package static

import (
	"context"

	"github.com/mgjules/mgjules-go/entity"
)

func (db *Static) GetIntroduction(ctx context.Context, id string) (*entity.Introduction, error) {
	return &entity.Introduction{
		ID:           id,
		Introduction: "A Senior Software Engineer specializing in building modern and scalable systems.",
		Avatar:       "/img/avatar.webp",
	}, nil
}
