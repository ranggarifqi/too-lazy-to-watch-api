package main

import (
	"net/http"
	"too-lazy-to-watch-api/helper"
	"too-lazy-to-watch-api/routes"
	v1 "too-lazy-to-watch-api/routes/v1"
	"too-lazy-to-watch-api/src/auth"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

func main() {
	helper.InitializeEnv("./.env")

	supabaseClient, adminClient := helper.GetSupabaseClient()

	authRepository := auth.NewSupabaseAuthRepository(supabaseClient, adminClient)

	e := echo.New()
	e.Validator = &routes.CustomValidator{Validator: validator.New()}

	v1Route := e.Group("/v1")
	v1.NewV1Handler(v1Route, v1.V1Dependencies{AuthRepository: authRepository})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":3000"))
}
