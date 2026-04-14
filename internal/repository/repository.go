package repository

import (
	"context"
	"errors"

	"github.com/mgjules/mgjules-go/internal/config"
	"github.com/mgjules/mgjules-go/internal/entity"
	"github.com/mgjules/mgjules-go/internal/repository/directus"
	"github.com/mgjules/mgjules-go/internal/repository/static"
)

type Repository interface {
	GetMeta(ctx context.Context, id string) (*entity.Meta, error)
	GetLinks(ctx context.Context) ([]entity.Link, error)
	GetIntroduction(ctx context.Context, id string) (*entity.Introduction, error)
	GetSections(ctx context.Context) ([]entity.Section, error)
	GetExperiences(ctx context.Context) ([]entity.Experience, error)
	GetProjects(ctx context.Context) ([]entity.Project, error)
	GetContributions(ctx context.Context) ([]entity.Contribution, error)
	GetAwards(ctx context.Context) ([]entity.Award, error)
	GetInterests(ctx context.Context) ([]entity.Interest, error)
	GetLanguages(ctx context.Context) ([]entity.Language, error)
	GetPosts(ctx context.Context) ([]entity.Post, error)
	Cleanup() error
}

func New(ctx context.Context, cfg *config.Config) (Repository, error) {
	if cfg.DirectusURL != "" {
		return directus.New(ctx, cfg.Prod, cfg.DirectusURL, cfg.DirectusToken), nil
	}
	if cfg.Static {
		return static.New(), nil
	}
	return nil, errors.New("no valid repository configuration found")
}
