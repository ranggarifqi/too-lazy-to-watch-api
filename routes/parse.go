package routes

import (
	custom_error "too-lazy-to-watch-api/src/error"

	"github.com/labstack/echo"
)

func ParseAndValidatePayload[T comparable](payload T, c echo.Context) error {
	err := c.Bind(payload)
	if err != nil {
		apiError := ConstructApiError(err)
		return c.JSON(apiError.Code, apiError)
	}

	err = c.Validate(payload)
	if err != nil {
		apiError := ConstructApiError(custom_error.NewBadRequestError(err.Error()))
		return c.JSON(apiError.Code, apiError)
	}

	return nil
}
