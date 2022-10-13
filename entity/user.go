package entity

import "github.com/edgedb/edgedb-go"

type User struct {
	ID        edgedb.UUID `edgedb:"id" json:"id"`
	FirstName string      `edgedb:"first_name" json:"first_name"`
	LastName  string      `edgedb:"last_name" json:"last_name"`
	Avatar    string      `edgedb:"avatar" json:"avatar"`
}
