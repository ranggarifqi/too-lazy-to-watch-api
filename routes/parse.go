package routes

import (
	custom_error "too-lazy-to-watch-api/src/error"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

func ParseAndValidatePayload[T comparable](payload T, c echo.Context) error {
	err := c.Bind(payload)
	if err != nil {
		return err
	}

	err = c.Validate(payload)
	if err != nil {
		return err
	}

	return nil
}

func ParseAndGetUserClaim(c echo.Context) (jwt.MapClaims, custom_error.Error) {
	token, ok := c.Get("user").(*jwt.Token) // by default token is stored under `user` key
	if !ok {
		return nil, custom_error.NewUnauthorizedError("JWT token missing or invalid")
	}
	claims, ok := token.Claims.(jwt.MapClaims) // by default claims is of type `jwt.MapClaims`
	if !ok {
		return nil, custom_error.NewUnauthorizedError("failed to cast claims as jwt.MapClaims")
	}

	return claims, nil
}
