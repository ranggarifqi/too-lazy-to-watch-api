package custom_middleware

import (
	"os"
	"too-lazy-to-watch-api/routes"
	custom_error "too-lazy-to-watch-api/src/error"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func GetJWTAuth() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
		ErrorHandler: func(c echo.Context, err error) error {
			return routes.HandleError(c, custom_error.NewUnauthorizedError("Invalid or missing JWT token"))
		},
		SuccessHandler: func(c echo.Context) {
			claim, _ := ParseAndGetUserClaim(c)
			c.Set("userClaim", claim)
		},
	})
}

func ParseAndGetUserClaim(c echo.Context) (jwt.MapClaims, custom_error.Error) {
	token, ok := c.Get("user").(*jwt.Token) // by default token is stored under `user` key
	if !ok {
		return nil, custom_error.NewUnauthorizedError("JWT token missing or invalid")
	}
	claims, ok := token.Claims.(jwt.MapClaims) // by default claims is of type `jwt.MapClaims`
	if !ok {
		return nil, custom_error.NewUnauthorizedError("failed to cast claims as jwt.MapClaims")
	}

	return claims, nil
}
