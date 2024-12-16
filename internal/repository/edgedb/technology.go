package edgedb

import (
	"github.com/edgedb/edgedb-go"
	"github.com/mgjules/mgjules-go/internal/entity"
)

type Technology struct {
	ID   edgedb.UUID `edgedb:"id"`
	Name string      `edgedb:"name"`
	Link string      `edgedb:"link"`
}

func (t Technology) ToEntity() entity.Technology {
	return entity.Technology{
		ID:   t.ID.String(),
		Name: t.Name,
		Link: t.Link,
	}
}
