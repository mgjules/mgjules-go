package entity

import "github.com/edgedb/edgedb-go"

type Link struct {
	ID           edgedb.UUID        `edgedb:"id" json:"id"`
	Name         string             `edgedb:"name" json:"name"`
	URL          string             `edgedb:"url" json:"url"`
	Icon         string             `edgedb:"icon" json:"icon"`
	AlternateURL edgedb.OptionalStr `edgedb:"alternate_url" json:"alternate_url"`
	NewWindow    bool               `edgedb:"new_window" json:"new_window"`
	IsCurrent    bool               `edgedb:"-" json:"is_current"`
}
