package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func init() {

}

type (
	Handler struct{}
)

func (h *Handler) Index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Kitty")
}
