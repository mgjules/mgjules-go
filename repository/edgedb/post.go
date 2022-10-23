package edgedb

import (
	"context"
	"fmt"
	"time"

	"github.com/araddon/dateparse"
	"github.com/edgedb/edgedb-go"
	"github.com/mgjules/mgjules-go/entity"
	"github.com/mgjules/mgjules-go/logger"
)

type Post struct {
	ID         edgedb.UUID             `edgedb:"id"`
	Title      string                  `edgedb:"title"`
	Slug       string                  `edgedb:"slug"`
	Summary    string                  `edgedb:"summary"`
	CoverImage string                  `edgedb:"cover_image"`
	Content    string                  `edgedb:"content"`
	Tags       []Tag                   `edgedb:"tags"`
	CreatedAt  edgedb.OptionalDateTime `edgedb:"created_at"`
	CreatedBy  User                    `edgedb:"created_by"`
}

func (p Post) ToEntity() entity.Post {
	tags := make([]entity.Tag, len(p.Tags))
	for i, tag := range p.Tags {
		tags[i] = tag.ToEntity()
	}

	var createdAt time.Time
	if val, ok := p.CreatedAt.Get(); ok {
		if parsed, err := dateparse.ParseAny(val.String()); err != nil {
			logger.L.Debugf("failed to parse `created_at` for `%s`: %v", p.ID, err)
		} else {
			createdAt = parsed
		}
	}

	return entity.Post{
		ID:         p.ID.String(),
		Title:      p.Title,
		Slug:       p.Slug,
		Summary:    p.Summary,
		CoverImage: p.CoverImage,
		Content:    p.Content,
		Tags:       tags,
		CreatedAt:  createdAt,
		CreatedBy:  p.CreatedBy.ToEntity(),
	}
}

func (db *EdgeDB) GetPosts(ctx context.Context) ([]entity.Post, error) {
	var results []Post
	err := db.client.Query(ctx, `
		select BlogPost {
			id,
			title,
			slug,
			summary,
			cover_image,
			content,
			tags: {
				name,
				slug
			},
			created_by: {
				first_name,
				last_name,
				avatar
			},
			created_at
		} order by .created_at desc
	`, &results)
	if err != nil {
		return nil, fmt.Errorf("failed to query posts: %w", err)
	}

	posts := make([]entity.Post, len(results))
	for i, result := range results {
		posts[i] = result.ToEntity()
	}

	return posts, nil
}
