package handler

import (
	"github.com/golang-jwt/jwt"
	"github.com/kristiannissen/grocery-monkey-golang/models"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

// TODO: Move to handler
type (
	jwtCustomClaims struct {
		NickName string `json:"nickname"`
		Uuid     string `json:"useruuid"`
		jwt.StandardClaims
	}
)

// TODO: Move to handler
const (
	secret string = "hello-kitty"
)

func (h *Handler) Authenticate(c echo.Context) error {
	var err error
	var msg Message

	m := models.Model{}
	u := m.NewUser()

	if err = c.Bind(u); err != nil {
		log.Printf("Request Error %s", err)
		msg.Text = err.Error()

		return c.JSON(http.StatusUnauthorized, m)
	}

	user := new(models.User)

	user, err = m.GetUser(u.NickName)
	if err != nil {
		// User was not found and will be created
		user, err = m.CreateUser(u.NickName)
		if err != nil {
			log.Printf("User could not be created %s", err)
			msg.Text = err.Error()

			return c.JSON(http.StatusInternalServerError, msg)
		}
	}

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
		msg.Text = err.Error()

		return c.JSON(http.StatusInternalServerError, msg)
	}

	return c.JSONPretty(http.StatusOK, echo.Map{
		"token": t,
		"useruuid":  u.Uuid,
	}, "  ")
}
