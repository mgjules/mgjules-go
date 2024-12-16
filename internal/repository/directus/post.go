package directus

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/avelino/slugify"
	"github.com/mgjules/mgjules-go/internal/entity"
)

type Post struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	Summary    string `json:"summary"`
	CoverImage string `json:"cover_image"`
	Content    string `json:"content"`
	Tags       []struct {
		Tag Tag `json:"tag"`
	} `json:"tags"`
	CreatedAt time.Time `json:"date_created"`
	CreatedBy User      `json:"user_created"`
}

func (p Post) ToEntity(directusURL string) entity.Post {
	tags := make([]entity.Tag, len(p.Tags))
	for i, tag := range p.Tags {
		tags[i] = tag.Tag.ToEntity()
	}

	author := p.CreatedBy.ToEntity(directusURL)
	author.Avatar += "?key=post-avatar"

	return entity.Post{
		ID:         p.ID,
		Title:      p.Title,
		Slug:       p.Slug,
		Summary:    p.Summary,
		CoverImage: directusURL + "/assets/" + p.CoverImage + "/" + slugify.Slugify(p.Title) + ".webp?key=post-cover",
		Content:    p.Content,
		Tags:       tags,
		CreatedAt:  p.CreatedAt,
		CreatedBy:  author,
	}
}

func (db *directus) GetPosts(ctx context.Context) ([]entity.Post, error) {
	var result Result[[]Post]
	resp, err := db.client.R().
		SetContext(ctx).
		SetQueryParamsFromValues(url.Values{
			"fields": []string{
				"id",
				"title",
				"slug",
				"summary",
				"cover_image",
				"content",
				"tags.tag.id",
				"tags.tag.name",
				"tags.tag.slug",
				"date_created",
				"user_created.first_name",
				"user_created.last_name",
				"user_created.description",
				"user_created.avatar",
				"user_created.github",
				"user_created.username",
				"user_created.gender",
			},
			"status": []string{"published"},
			"sort":   []string{"-date_created"},
		}).
		SetResult(&result).
		Get("/items/post")
	if err != nil {
		return nil, fmt.Errorf("failed to get posts: %w", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("failed to get posts: response code %d", resp.StatusCode())
	}

	posts := make([]entity.Post, len(result.Data))
	for i, post := range result.Data {
		posts[i] = post.ToEntity(db.directusURL)
	}

	return posts, nil
}
