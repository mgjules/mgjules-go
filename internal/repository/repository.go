package repository

import (
	"context"
	"errors"

	"github.com/mgjules/mgjules-go/internal/config"
	"github.com/mgjules/mgjules-go/internal/entity"
	"github.com/mgjules/mgjules-go/internal/repository/directus"
	"github.com/mgjules/mgjules-go/internal/repository/edgedb"
	"github.com/mgjules/mgjules-go/internal/repository/static"
)

type Repository interface {
	// Meta
	GetMeta(ctx context.Context, id string) (*entity.Meta, error)

	// Link
	GetLinks(ctx context.Context) ([]entity.Link, error)

	// Introduction
	GetIntroduction(ctx context.Context, id string) (*entity.Introduction, error)

	// Section
	GetSections(ctx context.Context) ([]entity.Section, error)

	// Experience
	GetExperiences(ctx context.Context) ([]entity.Experience, error)

	// Project
	GetProjects(ctx context.Context) ([]entity.Project, error)

	// Contribution
	GetContributions(ctx context.Context) ([]entity.Contribution, error)

	// Award
	GetAwards(ctx context.Context) ([]entity.Award, error)

	// Interest
	GetInterests(ctx context.Context) ([]entity.Interest, error)

	// Language
	GetLanguages(ctx context.Context) ([]entity.Language, error)

	// Post
	GetPosts(ctx context.Context) ([]entity.Post, error)

	Cleanup() error
}

func New(ctx context.Context, cfg *config.Config) Repository {
	var (
		repo Repository
		err  error
	)

	if cfg.EdgeDBDSN != "" {
		repo, err = edgedb.New(ctx, cfg.EdgeDBDSN)
	} else if cfg.DirectusURL != "" {
		repo = directus.New(ctx, cfg.Prod, cfg.DirectusURL, cfg.DirectusToken)
	} else if cfg.Static {
		repo = static.New()
	} else {
		err = errors.New("no valid configuration for repository found")
	}
	if err != nil {
		panic(err)
	}

	return repo
}
