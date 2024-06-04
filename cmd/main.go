package main

import (
	"net/http"
	"too-lazy-to-watch-api/helper"

	"github.com/labstack/echo"
)

func main() {
	helper.InitializeEnv("./.env")

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":3000"))
}
