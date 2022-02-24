package main

import (
	"database/sql"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"net/http"
	"os"
	"time"
)

// Constant secret
const (
	secret string = "pussysecret"
	DB_DNS        = os.Getenv("DATABASE_URL")
)

type (
	jwtCustomClaims struct {
		UserName string `json:"username"`
		Uid      string `json:"uid"`
		jwt.StandardClaims
	}

	Grocery struct {
		Name     string `json:"name"`
		Quantity int    `json:"qty"`
		Unit     string `json:"unit"`
		Store    string `json:"store"`
	}

	GroceryList struct {
		User        string    `json:"user"`
		Subscribers []string  `json:"subscribers"`
		Id          string    `json:"id" param:"id"`
		Groceries   []Grocery `json:"groceries"`
	}
)

// Handlers
func sign(c echo.Context) error {
	username := c.FormValue("username")
	// Generate a new UUID
	uuid := uuid.New()

	// Check form value
	if username == "" {
		return echo.ErrUnauthorized
	}
	// Set custom claims
	claims := &jwtCustomClaims{
		username,
		uuid.String(),
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
	return c.JSONPretty(http.StatusOK, echo.Map{
		"token": t,
	}, "  ")
}

// Restricted handlers
func newGroceryList(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	// We can now check the user using claims.UserName

	// Create UUID for the list
	listUid := uuid.New()

	groceries := GroceryList{
		User:        claims.Uid,
		Id:          listUid.String(),
		Subscribers: []string{"1", "2"},
		Groceries: []Grocery{
			Grocery{},
		},
	}

	return c.JSONPretty(http.StatusOK, groceries, "  ")
}

// Update grocerylist
func updateGroceryList(c echo.Context) error {
	// Create a new grocerylist
	groceries := new(GroceryList)

	if err := c.Bind(groceries); err != nil {
		return err
	}

	return c.JSONPretty(http.StatusOK, groceries, "  ")
}

// Delete grocerylist
func deleteGroceryList(c echo.Context) error {
	return c.String(http.StatusOK, "Deleted")
}

// Share list
func joinGroceryList(c echo.Context) error {
	return c.String(http.StatusOK, "Joined")
}

func main() {
	var port string = os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	e := echo.New()

	e.Use(middleware.Logger())
	// Default route
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello Kitty")
	})
	// Post to get token
	e.POST("/sign", sign)

	// DB func
	e.GET("/dbfunc", func(c echo.Context) error {
		db, err := sql.Open("postgres", DB_DNS)
		if err != nil {
			return err
		}
		defer db.Close()

		return c.String(http.StatusOK, "Hello Pussy")
	})

	// Join grocerylist
	e.GET("/join/:id", joinGroceryList)

	// Group that requires jwt token
	r := e.Group("/groceries")

	// Configure jwt
	config := middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte(secret),
	}

	// Use middleware with config
	r.Use(middleware.JWTWithConfig(config))

	// Restricted routes
	// Create new grocery list
	r.POST("", newGroceryList)

	// Update grocery list
	r.PUT("/:id/update", updateGroceryList)

	// Delete grocery list
	r.DELETE("/:id/delete", deleteGroceryList)

	e.Logger.Fatal(e.Start(":" + port))
}
