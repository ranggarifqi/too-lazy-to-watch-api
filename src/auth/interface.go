package auth

import "too-lazy-to-watch-api/src/user"

type ISignupPayload struct {
	Email    string
	Password string
	Name     string
}

type IAuthRepository interface {
	SignUpByEmail(payload ISignupPayload) (*user.User, error)
	// LoginByPassword(email string, password string) (string, error)
}
