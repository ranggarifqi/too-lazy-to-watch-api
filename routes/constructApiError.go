package routes

import (
	"fmt"
	"net/http"
	custom_error "too-lazy-to-watch-api/src/error"

	"github.com/labstack/echo"
)

func ConstructApiError(err error) *echo.HTTPError {
	myError, ok := err.(custom_error.Error)

	if ok {
		statusCode := myError.GetStatusCode()
		return echo.NewHTTPError(statusCode, myError.Error())
	}

	errMessage := fmt.Sprintf("Error 500: %v", err.Error())
	return echo.NewHTTPError(http.StatusInternalServerError, errMessage)
}
