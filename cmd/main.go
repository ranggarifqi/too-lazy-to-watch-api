package main

import (
	"fmt"
	"net/http"
	"too-lazy-to-watch-api/helper"
	"too-lazy-to-watch-api/routes"
	custom_middleware "too-lazy-to-watch-api/routes/middleware"
	v1 "too-lazy-to-watch-api/routes/v1"
	"too-lazy-to-watch-api/src/auth"
	custom_error "too-lazy-to-watch-api/src/error"
	"too-lazy-to-watch-api/src/storage"
	"too-lazy-to-watch-api/src/summary"
	"too-lazy-to-watch-api/src/taskPublisher"

	"github.com/go-playground/validator"
	"github.com/kkdai/youtube/v2"
	"github.com/labstack/echo/v4"
)

func main() {
	helper.InitializeEnv("./.env")

	supabaseClient, adminClient := helper.GetSupabaseClient()
	rabbitMQChannel := helper.GetRabbitMQChannel()

	ytClient := youtube.Client{}

	/* Dependency Setup */
	authRepository := auth.NewSupabaseAuthRepository(supabaseClient, adminClient)

	supabaseStorageRepository := storage.NewSupabaseStorageRepository(supabaseClient)
	rabbitMQPublisherRepository := taskPublisher.NewRabbitMQPublisherRepository(rabbitMQChannel)

	summaryRepository := summary.NewSupabaseSummaryRepository(supabaseClient)
	summaryService := summary.NewSummaryService(summaryRepository, ytClient, rabbitMQPublisherRepository, supabaseStorageRepository)

	e := echo.New()
	e.Validator = &routes.CustomValidator{Validator: validator.New()}

	v1Route := e.Group("/v1")
	v1.NewV1Handler(v1Route, v1.V1Dependencies{AuthRepository: authRepository, SummaryService: summaryService})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/admin/test-publish", func(c echo.Context) error {
		type Payload struct {
			Channel string `json:"channel"`
			Payload string `json:"payload"`
		}

		payload := new(Payload)
		if err := routes.ParseAndValidatePayload(payload, c); err != nil {
			return routes.HandleError(c, custom_error.NewBadRequestError(err.Error()))
		}

		rabbitMQPublisherRepository.Publish(payload.Channel, taskPublisher.PublishPayload{
			ContentType: "text/plain",
			Body:        []byte(payload.Payload),
		})

		return c.String(http.StatusOK, fmt.Sprintf("Published %s to channel %s", payload.Payload, payload.Channel))
	}, custom_middleware.AdminAuth)

	e.Logger.Fatal(e.Start(":3000"))
}
