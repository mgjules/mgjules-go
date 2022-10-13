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
}
