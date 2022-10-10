package entity

import "github.com/edgedb/edgedb-go"

type Introduction struct {
	ID           edgedb.UUID `edgedb:"id" json:"id"`
	Introduction string      `edgedb:"introduction" json:"introduction"`
	Avatar       string      `edgedb:"avatar" json:"avatar"`
}
