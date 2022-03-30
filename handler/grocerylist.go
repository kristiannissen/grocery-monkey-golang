package handler

import (
	"github.com/kristiannissen/grocery-monkey-golang/models"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func (h *Handler) CreateGroceryList(c echo.Context) error {
	var err error
	m := models.Model{}
	g := m.NewGroceryList()

	if err = c.Bind(g); err != nil {
		log.Printf("Request Error %s", err)
		return c.String(http.StatusInternalServerError, "Request Error")
	}

	// Store the grocerylist
	if g, err = m.CreateGroceryList(g); err != nil {
		log.Printf("GroceryList could not be created %q", err)
		return c.String(http.StatusInternalServerError, "Data Error")
	}
	// Add user as subscriber
	g.Subscribers = append(g.Subscribers, g.UserUuid)

	return c.JSONPretty(http.StatusCreated, g, "  ")
}

func (h *Handler) UpdateGroceryList(c echo.Context) error {
	var err error
	m := models.Model{}
	g := m.NewGroceryList()

	if err = c.Bind(g); err != nil {
		log.Printf("Request Error %s", err)
		return c.String(http.StatusInternalServerError, "Request Error")
	}

	// Store the grocerylist
	if g, err = m.UpdateGroceryList(g); err != nil {
		log.Printf("GroceryList could not be created %q", err)
		return c.String(http.StatusInternalServerError, "Data Error")
	}

	return c.JSONPretty(http.StatusOK, g, "  ")
}
