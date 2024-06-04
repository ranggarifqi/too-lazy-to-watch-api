package auth

type IAuthRepository interface {
	SignUpByEmail(email string, password string) error
	// LoginByPassword(email string, password string) (string, error)
}
