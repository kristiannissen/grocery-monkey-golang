package handler

import (
	"github.com/kristiannissen/grocery-monkey-golang/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func init() {

}

func (h *Handler) Index(c echo.Context) error {

	return c.String(http.StatusOK, "Hello Kitty")
}

func (h *Handler) SetUp(c echo.Context) error {
	m := models.Model{}
	m.UserSetUp()
	m.GroceryListSetUp()

	return c.String(http.StatusOK, "Setup working")
}
