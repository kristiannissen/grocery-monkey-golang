package main

import (
"github.com/kristiannissen/grocery-monkey-golang/handlers"
  "os"
  "github.com/labstack/echo/v4/middleware"
  "github.com/labstack/echo/v4"
  "github.com/labstack/gommon/log"
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
	h := &Handler{}
  // Say Hello!
  e.GET("/", h.Index)

  // Listen & Serve
  e.Logger.Fatal(e.Start(":" + port))
}
