package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

func main() {
	var port string = os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello Pussy!")
	})
	e.Logger.Fatal(e.Start(":" + port))
}
