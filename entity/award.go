package entity

import "github.com/edgedb/edgedb-go"

type Award struct {
	ID          edgedb.UUID      `edgedb:"id" json:"id"`
	Event       string           `edgedb:"event" json:"event"`
	Description string           `edgedb:"description" json:"description"`
	Date        edgedb.LocalDate `edgedb:"date" json:"date"`
	Link        string           `edgedb:"link" json:"link"`
	Result      string           `edgedb:"result" json:"result"`
	Icon        string           `edgedb:"icon" json:"icon"`
}
