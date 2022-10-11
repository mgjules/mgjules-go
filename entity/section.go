package entity

import "github.com/edgedb/edgedb-go"

type Section struct {
	ID   edgedb.UUID `edgedb:"id" json:"id"`
	Name string      `edgedb:"name" json:"name"`
	Icon string      `edgedb:"icon" json:"icon"`
}
