package static

import (
	"context"

	"github.com/mgjules/mgjules-go/internal/entity"
)

func (db *static) GetIntroduction(ctx context.Context, id string) (*entity.Introduction, error) {
	return &entity.Introduction{
		ID: id,
		Introduction: `Highly accomplished Senior Software Engineer with over 10 years of experience in 
    backend development of which 5 years specializing in Go. Also, a definite cat lover.`,
		Avatar: "/img/avatar.webp",
	}, nil
}
