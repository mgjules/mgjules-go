package entity

import "github.com/edgedb/edgedb-go"

type Project struct {
	ID           edgedb.UUID  `edgedb:"id" json:"id"`
	Name         string       `edgedb:"name" json:"name"`
	Description  string       `edgedb:"description" json:"description"`
	Link         string       `edgedb:"link" json:"link"`
	Technologies []Technology `edgedb:"technologies" json:"technologies"`
}
