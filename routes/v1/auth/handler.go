package v1_auth

import (
	"net/http"
	"too-lazy-to-watch-api/routes"
	"too-lazy-to-watch-api/src/auth"
	custom_error "too-lazy-to-watch-api/src/error"

	"github.com/labstack/echo"
)

type authHandler struct {
	authRepository auth.IAuthRepository
}

func NewAuthHandler(g *echo.Group, authRepository auth.IAuthRepository) {
	h := &authHandler{
		authRepository: authRepository,
	}

	// TODO: Use basic auth to protect this endpoint
	g.POST("/admin/signup", h.SignUp)
}

func (h *authHandler) SignUp(c echo.Context) error {
	payload := new(SignUpDTO)
	err := c.Bind(payload)
	if err != nil {
		apiError := routes.ConstructApiError(err)
		return c.JSON(apiError.Code, apiError)
	}

	err = c.Validate(payload)
	if err != nil {
		apiError := routes.ConstructApiError(custom_error.NewBadRequestError(err.Error()))
		return c.JSON(apiError.Code, apiError)
	}

	res, err := h.authRepository.SignUpByEmail(auth.ISignupPayload{
		Email:    payload.Email,
		Password: payload.Password,
		Name:     payload.Name,
	})
	if err != nil {
		apiError := routes.ConstructApiError(custom_error.NewBadRequestError(err.Error()))
		return c.JSON(apiError.Code, apiError)
	}

	return c.JSON(http.StatusOK, res)
}
