package entity

import "github.com/edgedb/edgedb-go"

type Interest struct {
	ID    edgedb.UUID `edgedb:"id" json:"id"`
	Name  string      `edgedb:"name" json:"name"`
	Image string      `edgedb:"image" json:"image"`
}
