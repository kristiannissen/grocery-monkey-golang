package main

import (
	//"encoding/json"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
	"time"
)

// Constant secret
const secret string = "pussysecret"

type (
	jwtCustomClaims struct {
        UserName string `json:"username"`
		jwt.StandardClaims
	}
)

// Handlers
func sign(c echo.Context) error {
    username := c.FormValue("username")

    // Check form value
    if username == "" {
        return echo.ErrUnauthorized
    }
    // Set custom claims
    claims := &jwtCustomClaims{
        username,
        jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
        },
    }

    // Create token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    // Generate token
    t, err := token.SignedString([]byte(secret))
    if err != nil {
        return err
    }

    // Response with token
    return c.JSON(http.StatusOK, echo.Map{
        "token": t,
    })
}
// Restricted handlers
func list(c echo.Context) error {
    return c.String(http.StatusOK, "Hello")
}

func main() {
	var port string = os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	e := echo.New()

    // Default route
    e.GET("/", func(c echo.Context) error {
        return c.HTML(http.StatusOK, "Hello Kitty")
    })
    // Post to get token
    e.POST("/sign", sign)

    // Group that requires jwt token
    r := e.Group("/list")

    // Configure jwt
    config := middleware.JWTConfig{
        Claims: &jwtCustomClaims{},
        SigningKey: []byte(secret),
    }

    // Use middleware with config
    r.Use(middleware.JWTWithConfig(config))

    // Restricted routes
    r.GET("/new", list)

	e.Logger.Fatal(e.Start(":" + port))
}
