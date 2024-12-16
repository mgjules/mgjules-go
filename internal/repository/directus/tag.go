package directus

import "github.com/mgjules/mgjules-go/internal/entity"

type Tag struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (t Tag) ToEntity() entity.Tag {
	return entity.Tag{
		ID:   t.ID,
		Name: t.Name,
		Slug: t.Slug,
	}
}
