package entity

import "github.com/edgedb/edgedb-go"

type Technology struct {
	ID   edgedb.UUID `edgedb:"id" json:"id"`
	Name string      `edgedb:"name" json:"name"`
}
