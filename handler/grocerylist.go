package handler

import (
	"github.com/kristiannissen/grocery-monkey-golang/models"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func (h *Handler) CreateGroceryList(c echo.Context) error {
	var err error
	var msg Message

	m := models.Model{}
	g := m.NewGroceryList()

	if err = c.Bind(g); err != nil {
		log.Printf("Request Error %q", err)
		msg.Text = err.Error()

		return c.JSON(http.StatusInternalServerError, msg)
	}

	// Store the grocerylist
	if g, err = m.CreateGroceryList(g); err != nil {
		log.Printf("GroceryList could not be created %s", err)
		msg.Text = err.Error()

		return c.JSON(http.StatusInternalServerError, msg)
	}
	// Add user as subscriber
	g.Subscribers = append(g.Subscribers, g.UserUuid)

	return c.JSONPretty(http.StatusCreated, g, "  ")
}

func (h *Handler) UpdateGroceryList(c echo.Context) error {
	var err error
	var msg Message
	m := models.Model{}
	g := m.NewGroceryList()

	if err = c.Bind(g); err != nil {
		log.Printf("Request Error %s", err)
		msg.Text = err.Error()

		return c.JSON(http.StatusInternalServerError, msg)
	}

	// Get the grocerylist
	if g, err = m.GetGroceryList(g.UserUuid); err != nil {
		log.Printf("Could not find grocerylist %s", err)
		msg.Text = err.Error()

		return c.JSON(http.StatusOK, msg)
	}

	// Store the grocerylist
	if g, err = m.UpdateGroceryList(g); err != nil {
		log.Printf("GroceryList could not be created %s", err)
		msg.Text = err.Error()

		return c.JSON(http.StatusInternalServerError, msg)
	}

	return c.JSONPretty(http.StatusOK, g, "  ")
}
