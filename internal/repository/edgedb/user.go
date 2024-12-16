package edgedb

import (
	"github.com/edgedb/edgedb-go"
	"github.com/mgjules/mgjules-go/internal/entity"
)

type User struct {
	ID        edgedb.UUID `edgedb:"id"`
	FirstName string      `edgedb:"first_name"`
	LastName  string      `edgedb:"last_name"`
	Avatar    string      `edgedb:"avatar"`
}

func (u User) ToEntity() entity.User {
	return entity.User{
		ID:        u.ID.String(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Avatar:    u.Avatar,
	}
}
