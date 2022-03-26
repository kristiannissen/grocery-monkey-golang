package main

import (
	"github.com/golang-jwt/jwt"
	"github.com/kristiannissen/grocery-monkey-golang/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"net/http"
	"os"
)

type (
	jwtCustomClaims struct {
		NickName string `json:"nickname"`
		Uuid     string `json:"useruuid"`
		jwt.StandardClaims
	}
)

const (
	secret string = "hello-kitty"
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
	// Add CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// Init the handlers
	h := &handler.Handler{}
	// Say Hello!
	e.GET("/", h.Index)
	// DB setup
	e.GET("/setup", h.SetUp)
	// Authenticate user
	e.POST("/api/authenticate", h.Authenticate)

	// Config JWT
	config := middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte(secret),
	}

	// Groups that require token
	g := e.Group("/api")
	// Add jwt to group
	g.Use(middleware.JWTWithConfig(config))
	// Create groceries
	g.POST("/groceries", h.CreateGroceryList)
	// Update groceries
	g.PUT("/groceries", h.UpdateGroceryList)
	// JWT Test
	g.GET("/jwttest", func(c echo.Context) error {
		log.Info("JWT is working")
		return c.String(http.StatusOK, "Hello")
	})

	// Listen & Serve
	e.Logger.Fatal(e.Start(":" + port))
}
