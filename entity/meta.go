package entity

import "github.com/edgedb/edgedb-go"

type Meta struct {
	ID          edgedb.UUID `edgedb:"id" json:"id"`
	BaseURL     string      `edgedb:"base_url" json:"base_url"`
	Lang        string      `edgedb:"lang" json:"lang"`
	Description string      `edgedb:"description" json:"description"`
	FirstName   string      `edgedb:"first_name" json:"first_name"`
	LastName    string      `edgedb:"last_name" json:"last_name"`
	Keywords    []string    `edgedb:"keywords" json:"keywords"`
	Github      string      `edgedb:"github" json:"github"`
	Username    string      `edgedb:"username" json:"username"`
	Gender      string      `edgedb:"gender" json:"gender"`
	Avatar      string      `edgedb:"avatar" json:"avatar"`
}
