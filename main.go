package main

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
)

type (
	Item struct {
		Name  string
		Key   string
		Store string
		QTY   string
		Unit  string
	}
)

var (
	items = []Item{
		{"Milk", "1", "Meny", "4", "stk"},
		{"Bread", "2", "Netto", "2", "kg"},
		{"Apple", "3", "Irma", "1", "pose"},
	}
)

func main() {
	var port string = os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello Sick Ass Pussy!")
	})

	e.GET("/items", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusOK)

		return json.NewEncoder(c.Response()).Encode(items)
	})
	e.Logger.Fatal(e.Start(":" + port))
}
