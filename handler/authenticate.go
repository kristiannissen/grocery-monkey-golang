package handler

import (
	"github.com/golang-jwt/jwt"
	"github.com/kristiannissen/grocery-monkey-golang/models"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
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

func (h *Handler) Authenticate(c echo.Context) error {
	var err error

	m := models.Model{}
	u := m.NewUser()

	if err = c.Bind(u); err != nil {
		log.Printf("Request Error %s", err)
		return c.String(http.StatusUnauthorized, "Request Error")
	}

	user := new(models.User)

	user, err = m.GetUser(u.NickName)
	if err != nil {
		// User was not found and will be created
		user, err = m.CreateUser(u.NickName)
		if err != nil {
			log.Printf("User could not be created %s", err)
			return c.String(http.StatusInternalServerError, "User could not be created")
		}
	}

	groceryList := m.NewGroceryList()
	groceryList.UserUuid = user.Uuid
	groceryList.Subscribers = append(groceryList.Subscribers, user.Uuid)

	claims := &jwtCustomClaims{
		user.NickName,
		user.Uuid,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Printf("JWT token error %s", err)
		return c.String(http.StatusInternalServerError, "JWT error")
	}

	return c.JSONPretty(http.StatusOK, echo.Map{
		"token":     t,
		"groceries": groceryList,
	}, "  ")
}
