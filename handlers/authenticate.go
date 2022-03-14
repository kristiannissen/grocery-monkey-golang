package handlers

import (
	"github.com/golang-jwt/jwt"
	"github.com/kristiannissen/grocery-monkey-golang/models"
	"github.com/labstack/echo/v4"
	_ "log"
	"net/http"
	"time"
)

type (
	jwtCustomClaims struct {
		NickName string `json:"nickname"`
		Uuid     string `json:"uuid"`
		jwt.StandardClaims
	}
)

const (
	secret string = "hellokitty"
)

func (h *Handler) Authenticate(c echo.Context) error {

	m := models.Model{}
	u := m.NewUser()

	if err := c.Bind(u); err != nil {
		return c.String(http.StatusUnauthorized, "Request Error")
	}

	user, err := m.GetUser(u.NickName)
	if err != nil {
		return c.String(http.StatusUnauthorized, "User exists")
	}

	user = m.CreateUser(u.NickName)

	groceryList := m.NewGroceryList()
	groceryList.UserUuid = user.Uuid

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
		return c.String(http.StatusUnauthorized, "JWT error")
	}

	return c.JSONPretty(http.StatusCreated, echo.Map{
		"token":     t,
		"groceries": groceryList,
	}, "  ")
}
