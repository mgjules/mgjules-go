package entity

import (
	"time"

	"github.com/edgedb/edgedb-go"
)

type Post struct {
	ID         edgedb.UUID `edgedb:"id" json:"id"`
	Title      string      `edgedb:"title" json:"title"`
	Slug       string      `edgedb:"slug" json:"slug"`
	Summary    string      `edgedb:"summary" json:"summary"`
	CoverImage string      `edgedb:"cover_image" json:"cover_image"`
	Content    string      `edgedb:"content" json:"content"`
	Tags       []Tag       `edgedb:"tags" json:"tags"`
	CreatedAt  time.Time   `edgedb:"created_at" json:"created_at"`
	CreatedBy  User        `edgedb:"created_by" json:"created_by"`
}
