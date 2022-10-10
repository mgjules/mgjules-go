package repository

import (
	"context"

	"github.com/mgjules/mgjules-go/entity"
)

type Repository interface {
	// Meta
	GetMeta(ctx context.Context, id string) (*entity.Meta, error)

	// Link
	GetLinks(ctx context.Context) ([]entity.Link, error)

	// Introduction
	GetIntroduction(ctx context.Context, id string) (*entity.Introduction, error)
}
