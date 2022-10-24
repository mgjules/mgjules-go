package directus

type User struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Description string `json:"description"`
	Avatar      string `json:"avatar"`
	Github      string `json:"github"`
	Username    string `json:"username"`
	Gender      string `json:"gender"`
}
