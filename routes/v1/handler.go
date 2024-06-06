package v1

import (
	custom_middleware "too-lazy-to-watch-api/routes/middleware"
	v1_admin "too-lazy-to-watch-api/routes/v1/admin"
	v1_auth "too-lazy-to-watch-api/routes/v1/auth"
	"too-lazy-to-watch-api/src/auth"

	"github.com/labstack/echo/v4"
)

type V1Dependencies struct {
	AuthRepository auth.IAuthRepository
}

func NewV1Handler(g *echo.Group, dep V1Dependencies) {

	adminRoute := g.Group("/admin")
	authRoute := g.Group("/auth")

	adminRoute.Use(custom_middleware.AdminAuth)

	v1_admin.NewHandler(adminRoute, dep.AuthRepository)
	v1_auth.NewHandler(authRoute, dep.AuthRepository)
}
