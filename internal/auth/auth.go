package auth

import "errors"

type Auth struct {
	token string
}

func New(token string) (*Auth, error) {
	if token == "" {
		return nil, errors.New("token cannot be empty")
	}
	return &Auth{token: token}, nil
}

func (a *Auth) Validate(token string) bool {
	return a.token == token
}
