package main

import (
	"github.com/kristiannissen/grocery-monkey-golang/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"os"
)

func main() {
	var port string = os.Getenv("PORT")
	// If no env PORT is passed used 8080
	if port == "" {
		port = "8080"
	}

	// Init echo
	e := echo.New()
	// Configuration
	e.Logger.SetLevel(log.DEBUG)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Init the handlers
	h := &handler.Handler{}
	// Say Hello!
	e.GET("/", h.Index)
	// DB setup
	e.GET("/setup", h.SetUp)
	// Authenticate user
	e.POST("/api/authenticate", h.Authenticate)
	// Groups that require token
	g := e.Group("/api")
	// Create groceries
	g.POST("/groceries", h.CreateGroceryList)

	// Listen & Serve
	e.Logger.Fatal(e.Start(":" + port))
}
