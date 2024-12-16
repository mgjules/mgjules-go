package entity

type Link struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	URL          string  `json:"url"`
	Icon         string  `json:"icon"`
	AlternateURL *string `json:"alternate_url"`
	NewWindow    bool    `json:"new_window"`
	IsCurrent    bool    `json:"is_current"`
}
