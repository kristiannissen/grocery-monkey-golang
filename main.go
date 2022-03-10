package main

import (
	"database/sql"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"net/http"
	"os"
	"time"
)

// Constant secret
const (
	secret string = "pussysecret"
)

type (
	User struct {
		UserName string `json:"username" form:"username"`
		Uuid     string `json:"uuid" form:"uuid"`
	}

	jwtCustomClaims struct {
		UserName string `json:"username"`
		Uid      string `json:"uid"`
		jwt.StandardClaims
	}

	Grocery struct {
		Name string `json:"name"`
		Id   string `json:"id"`
	}

	GroceryList struct {
		User        string    `json:"user"`
		Subscribers []string  `json:"subscribers"`
		Id          string    `json:"id" param:"id"`
		Groceries   []Grocery `json:"groceries"`
	}

	Handler struct {
		DB *sql.DB
	}
)

// Handlers
func (h *Handler) Home(c echo.Context) error {
	// Close DB
	defer h.DB.Close()

	return c.HTML(http.StatusOK, "Hello Kitty")
}

func (h *Handler) DBFuncs(c echo.Context) error {
    defer h.DB.Close()

    if _, err := h.DB.Exec("CREATE TABLE users (id serial NOT NULL, nickname varchar, uuid varchar)"); err != nil {
        log.Fatal("DB Error %q", err)
        return err
    }

    return c.HTML(http.StatusOK, "Done")
}

func authenticate(c echo.Context) error {
	// Create a new user
	user := new(User)
	if err := c.Bind(user); err != nil {
		log.Warn("User not in request")
		return echo.ErrUnauthorized
	}
	// Check if username is empty
	if user.UserName == "" {
		log.Warn("Empty username in request")
		return echo.ErrUnauthorized
	}

	// Add Uuid
	user.Uuid = uuid.New().String()

	// Create Claims
	claims := &jwtCustomClaims{
		user.UserName,
		user.Uuid,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign token
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return err
	}
	// Faking a new list
	groceries := GroceryList{
		User:        user.Uuid,
		Id:          uuid.New().String(),
		Subscribers: []string{},
		Groceries:   []Grocery{},
	}

	return c.JSONPretty(http.StatusOK, echo.Map{
		"token":     t,
		"groceries": groceries,
	}, "  ")
}

// Restricted handlers
func newGroceryList(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	// We can now check the user using claims.UserName

	// TODO: Check if user has a list already,
	// if not, create a new empty list
	// Create UUID for the list
	listUid := uuid.New()

	groceries := GroceryList{
		User:        claims.Uid,
		Id:          listUid.String(),
		Subscribers: []string{"1", "2"},
		Groceries:   []Grocery{},
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
	// Iterate over groceries and assign uuid
	g := []Grocery{}
	for _, v := range groceries.Groceries {
		v.Id = uuid.New().String()
		g = append(g, v)
	}
	groceries.Groceries = g

	return c.JSONPretty(http.StatusOK, groceries, "  ")
}

// Show specific list
func showGroceryList(c echo.Context) error {
	id := c.Param("id")
	// Get the user
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	// Return the list
	groceries := new(GroceryList)
	groceries.User = claims.Uid
	groceries.Id = id

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
	// Logging for debug
	e.Logger.SetLevel(log.DEBUG)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Config cors
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// DB setup
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("DB ERROR %q", err)
	}

	// Initialize handlers
	h := Handler{DB: db}

	// Default route
	e.GET("/", h.Home)
    e.GET("/dbfuncs", h.DBFuncs)
	// Post to get token
	e.POST("/api/authenticate", authenticate)

	// Join grocerylist
	e.GET("/api/join/:id", joinGroceryList)

	// Group that requires jwt token
	r := e.Group("/api/groceries")

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

	// Show grocery list
	r.GET("/:id", showGroceryList)

	// Delete grocery list
	r.DELETE("/:id/delete", deleteGroceryList)

	e.Logger.Fatal(e.Start(":" + port))
}
