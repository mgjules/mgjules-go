package repository

import (
	"context"
	"errors"

	"github.com/mgjules/mgjules-go/config"
	"github.com/mgjules/mgjules-go/entity"
	"github.com/mgjules/mgjules-go/repository/directus"
	"github.com/mgjules/mgjules-go/repository/edgedb"
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
}

func New(ctx context.Context, cfg *config.Config) (Repository, error) {
	var (
		repo Repository
		err  error
	)

	if cfg.EdgeDBDSN != "" {
		repo, err = edgedb.New(ctx, cfg.EdgeDBDSN)
	} else if cfg.DirectusURL != "" {
		repo = directus.New(ctx, cfg.Prod, cfg.DirectusURL, cfg.DirectusToken)
	} else {
		err = errors.New("no configuration for a valid repository found")
	}

	return repo, err
}
