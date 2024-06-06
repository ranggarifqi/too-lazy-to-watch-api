package custom_error

import (
	"fmt"
	"net/http"
)

type badRequestError CustomError

func NewBadRequestError(errMessage string) Error {
	return &badRequestError{
		statusCode: http.StatusBadRequest,
		errMessage: errMessage,
	}
}

func (e *badRequestError) Error() string {
	return fmt.Sprintf("Error %v: %v", e.statusCode, e.errMessage)
}

func (e *badRequestError) GetStatusCode() int {
	return e.statusCode
}
