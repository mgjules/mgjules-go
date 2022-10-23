package edgedb

import (
	"github.com/edgedb/edgedb-go"
	"github.com/mgjules/mgjules-go/entity"
)

type Tag struct {
	ID   edgedb.UUID `edgedb:"id"`
	Name string      `edgedb:"name"`
	Slug string      `edgedb:"slug"`
}

func (t Tag) ToEntity() entity.Tag {
	return entity.Tag{
		ID:   t.ID.String(),
		Name: t.Name,
		Slug: t.Slug,
	}
}
