package v1_auth

type SignUpDTO struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}
