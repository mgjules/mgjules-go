package edgedb

import (
	"github.com/edgedb/edgedb-go"
	"github.com/mgjules/mgjules-go/entity"
)

type Technology struct {
	ID   edgedb.UUID `edgedb:"id"`
	Name string      `edgedb:"name"`
}

func (t Technology) ToEntity() entity.Technology {
	return entity.Technology{
		ID:   t.ID.String(),
		Name: t.Name,
	}
}
