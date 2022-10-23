package entity

import (
	"time"
)

type Experience struct {
	ID           string       `json:"id"`
	Company      string       `json:"company"`
	Position     string       `json:"position"`
	From         time.Time    `json:"from"`
	To           *time.Time   `json:"to"`
	Link         string       `json:"link"`
	Technologies []Technology `json:"technologies"`
	Tasks        []string     `json:"tasks"`
}
