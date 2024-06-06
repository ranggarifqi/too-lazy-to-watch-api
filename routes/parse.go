package routes

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
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

func GetUserClaim(c echo.Context) jwt.MapClaims {
	return c.Get("userClaim").(jwt.MapClaims)
}
