package entity

import (
	"time"
)

type Post struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	Slug       string    `json:"slug"`
	Summary    string    `json:"summary"`
	CoverImage string    `json:"cover_image"`
	Content    string    `json:"content"`
	Tags       []Tag     `json:"tags"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  User      `json:"created_by"`
}
