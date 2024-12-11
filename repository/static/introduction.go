package static

import (
	"context"

	"github.com/mgjules/mgjules-go/entity"
)

func (db *Static) GetIntroduction(ctx context.Context, id string) (*entity.Introduction, error) {
	return &entity.Introduction{
		ID:           id,
		Introduction: "Hello kitty!",
		Avatar:       "/img/avatar.webp",
	}, nil
}
