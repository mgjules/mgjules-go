package entity

import "github.com/edgedb/edgedb-go"

type Contribution struct {
	ID    edgedb.UUID              `edgedb:"id" json:"id"`
	Event string                   `edgedb:"event" json:"event"`
	Title string                   `edgedb:"title" json:"title"`
	From  edgedb.LocalDate         `edgedb:"from" json:"from"`
	To    edgedb.OptionalLocalDate `edgedb:"to" json:"to"`
	Link  string                   `edgedb:"link" json:"link"`
	Role  string                   `edgedb:"role" json:"role"`
}
