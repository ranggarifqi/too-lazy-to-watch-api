package custom_middleware

import (
	"os"
	custom_error "too-lazy-to-watch-api/src/error"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var AdminAuth = middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
	if username != os.Getenv("ADMIN_USERNAME") || password != os.Getenv("ADMIN_PASSWORD") {
		return false, custom_error.NewUnauthorizedError("Invalid credentials")
	}
	return true, nil
})
