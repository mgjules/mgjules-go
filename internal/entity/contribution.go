package entity

import (
	"time"
)

type Contribution struct {
	ID    string     `json:"id"`
	Event string     `json:"event"`
	Title string     `json:"title"`
	From  time.Time  `json:"from"`
	To    *time.Time `json:"to"`
	Link  string     `json:"link"`
	Role  string     `json:"role"`
}
