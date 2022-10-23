package entity

type Project struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	Link         string       `json:"link"`
	Technologies []Technology `json:"technologies"`
}
