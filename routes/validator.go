package routes

import (
	custom_error "too-lazy-to-watch-api/src/error"

	"github.com/go-playground/validator"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		return ConstructApiError(custom_error.NewBadRequestError(err.Error()))
	}
	return nil
}
