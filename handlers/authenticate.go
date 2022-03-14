package handlers

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
		Uuid     string `json:"uuid"`
		jwt.StandardClaims
	}
)

const (
	secret string = "hellokitty"
)

func (h *Handler) Authenticate(c echo.Context) error {

	m := models.Model{}
	user := m.NewUser()

	if err := c.Bind(user); err != nil {
		log.Println("Could not bind")
		return echo.ErrUnauthorized
	}

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
		return err
	}

	return c.JSONPretty(http.StatusOK, echo.Map{
		"token":     t,
		"groceries": groceryList,
	}, "  ")
}
