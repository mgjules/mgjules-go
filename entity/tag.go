package entity

import "github.com/edgedb/edgedb-go"

type Tag struct {
	ID   edgedb.UUID `edgedb:"id" json:"id"`
	Name string      `edgedb:"name" json:"name"`
	Slug string      `edgedb:"slug" json:"slug"`
}
