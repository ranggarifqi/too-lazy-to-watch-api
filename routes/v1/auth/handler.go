package v1_auth

import (
	"net/http"
	"too-lazy-to-watch-api/routes"
	"too-lazy-to-watch-api/src/auth"
	custom_error "too-lazy-to-watch-api/src/error"

	"github.com/labstack/echo"
)

type handler struct {
	authRepository auth.IAuthRepository
}

func NewHandler(g *echo.Group, authRepository auth.IAuthRepository) {
	h := &handler{
		authRepository: authRepository,
	}

	g.POST("/signin", h.SignIn)
}

func (h *handler) SignIn(c echo.Context) error {
	payload := new(SignInDTO)
	if err := routes.ParseAndValidatePayload(payload, c); err != nil {
		return routes.HandleError(c, custom_error.NewBadRequestError(err.Error()))
	}

	token, err := h.authRepository.SignInWithEmailPassword(payload.Email, payload.Password)
	if err != nil {
		return routes.HandleError(c, custom_error.NewBadRequestError(err.Error()))
	}

	res := struct {
		Token string `json:"token"`
	}{
		Token: token,
	}

	return c.JSON(http.StatusOK, res)
}
