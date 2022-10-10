package auth

type Auth struct {
	token string
}

func New(token string) *Auth {
	return &Auth{token: token}
}

func (a *Auth) Validate(token string) bool {
	return a.token == token
}
