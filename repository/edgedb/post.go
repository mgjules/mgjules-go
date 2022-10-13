package edgedb

import (
	"context"
	"fmt"

	"github.com/mgjules/mgjules-go/entity"
)

func (db *EdgeDB) GetPosts(ctx context.Context) ([]entity.Post, error) {
	var posts []entity.Post
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
	`, &posts)
	if err != nil {
		return nil, fmt.Errorf("failed to query posts: %w", err)
	}

	return posts, nil
}
