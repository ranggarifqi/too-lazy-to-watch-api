package main

import (
	"net/http"
	"too-lazy-to-watch-api/helper"
	"too-lazy-to-watch-api/routes"
	v1 "too-lazy-to-watch-api/routes/v1"
	"too-lazy-to-watch-api/src/auth"
	"too-lazy-to-watch-api/src/summary"

	"github.com/go-playground/validator"
	"github.com/kkdai/youtube/v2"
	"github.com/labstack/echo/v4"
)

func main() {
	helper.InitializeEnv("./.env")

	supabaseClient, adminClient := helper.GetSupabaseClient()
	_ = helper.GetRabbitMQChannel()

	ytClient := youtube.Client{}

	/* Dependency Setup */
	authRepository := auth.NewSupabaseAuthRepository(supabaseClient, adminClient)

	summaryRepository := summary.NewSupabaseSummaryRepository(supabaseClient)
	summaryService := summary.NewSummaryService(summaryRepository, ytClient)

	e := echo.New()
	e.Validator = &routes.CustomValidator{Validator: validator.New()}

	v1Route := e.Group("/v1")
	v1.NewV1Handler(v1Route, v1.V1Dependencies{AuthRepository: authRepository, SummaryService: summaryService})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":3000"))
}
