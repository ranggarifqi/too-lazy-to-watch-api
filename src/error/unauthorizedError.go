package custom_error

import (
	"fmt"
	"net/http"
)

type unauthorizedError CustomError

func NewUnauthorizedError(errMessage string) Error {
	return &unauthorizedError{
		statusCode: http.StatusUnauthorized,
		errMessage: errMessage,
	}
}

func (e *unauthorizedError) Error() string {
	return fmt.Sprintf("Error %v: %v", e.statusCode, e.errMessage)
}

func (e *unauthorizedError) GetStatusCode() int {
	return e.statusCode
}
