package custom_middleware

import (
	"os"

	"github.com/labstack/echo/middleware"
)

var JwtAuth = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte(os.Getenv("JWT_SECRET")),
})
