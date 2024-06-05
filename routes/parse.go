package routes

import (
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
