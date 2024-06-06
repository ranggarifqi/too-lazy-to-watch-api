package v1_admin

import (
	"net/http"
	"too-lazy-to-watch-api/routes"
	"too-lazy-to-watch-api/src/auth"
	custom_error "too-lazy-to-watch-api/src/error"

	"github.com/labstack/echo/v4"
)

type handler struct {
	authRepository auth.IAuthRepository
}

func NewHandler(g *echo.Group, authRepository auth.IAuthRepository) {
	h := &handler{
		authRepository: authRepository,
	}

	// TODO: Use basic auth to protect this endpoint
	g.POST("/signup", h.SignUp)
}

// TODO: Protect this by using BASIC auth
func (h *handler) SignUp(c echo.Context) error {
	payload := new(SignUpDTO)
	if err := routes.ParseAndValidatePayload(payload, c); err != nil {
		return routes.HandleError(c, custom_error.NewBadRequestError(err.Error()))
	}

	res, err := h.authRepository.SignUpByEmail(auth.ISignupPayload{
		Email:    payload.Email,
		Password: payload.Password,
		Name:     payload.Name,
	})
	if err != nil {
		return routes.HandleError(c, custom_error.NewBadRequestError(err.Error()))
	}

	return c.JSON(http.StatusOK, res)
}
