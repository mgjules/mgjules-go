package entity

import "time"

type Award struct {
	ID          string    `json:"id"`
	Event       string    `json:"event"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Link        string    `json:"link"`
	Result      string    `json:"result"`
	Icon        string    `json:"icon"`
}
