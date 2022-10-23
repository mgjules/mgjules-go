package entity

type Meta struct {
	ID          string   `json:"id"`
	BaseURL     string   `json:"base_url"`
	Lang        string   `json:"lang"`
	Description string   `json:"description"`
	FirstName   string   `json:"first_name"`
	LastName    string   `json:"last_name"`
	FullName    string   `json:"full_name"`
	Keywords    []string `json:"keywords"`
	Github      string   `json:"github"`
	Username    string   `json:"username"`
	Gender      string   `json:"gender"`
	Avatar      string   `json:"avatar"`
}
