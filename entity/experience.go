package entity

import "github.com/edgedb/edgedb-go"

type Experience struct {
	ID           edgedb.UUID              `edgedb:"id" json:"id"`
	Company      string                   `edgedb:"company" json:"company"`
	Position     string                   `edgedb:"position" json:"position"`
	From         edgedb.LocalDate         `edgedb:"from" json:"from"`
	To           edgedb.OptionalLocalDate `edgedb:"to" json:"to"`
	Link         string                   `edgedb:"link" json:"link"`
	Technologies []Technology             `edgedb:"technologies" json:"technologies"`
	Tasks        []string                 `edgedb:"tasks" json:"tasks"`
}
