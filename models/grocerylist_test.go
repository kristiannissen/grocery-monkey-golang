package models

import (
	_ "log"
	"testing"
)

func init() {
	m.GroceryListTearDown()
	m.GroceryListSetUp()
}

func TestCreateGroceryList(t *testing.T) {

	// Get a new struct
	grocerylist := m.NewGroceryList()
	// Get a user
	user := m.NewUser()
	grocerylist.UserUuid = user.Uuid
	grocerylist.Groceries = []Grocery{
		{"Beer", "2"},
	}
	// Add the user as a subscriber
	grocerylist.Subscribers = append(grocerylist.Subscribers, user.Uuid)

	grocerylist, err := m.CreateGroceryList(grocerylist)
	// log.Print(grocerylist)

	if err != nil {
		t.Errorf("Grocerylist could not be created %q", err)
	}
}

func TestGetGroceryList(t *testing.T) {
	// Get a new struct
	grocerylist := m.NewGroceryList()
	// Get a user
	user := m.NewUser()
	grocerylist.UserUuid = user.Uuid

	grocerylist, _ = m.CreateGroceryList(grocerylist)

	// Get grocerylist by user uuid
	_, err := m.GetGroceryList(user.Uuid)

	if err != nil {
		t.Errorf("Grocerylist could not be found %q", err)
	}
}

func TestUpdateGroceryList(t *testing.T) {
	// Get a new struct
	grocerylist := m.NewGroceryList()
	// Get a user
	user := m.NewUser()
	grocerylist.UserUuid = user.Uuid
	grocerylist.Subscribers = append(grocerylist.Subscribers, user.Uuid)

	grocerylist, _ = m.CreateGroceryList(grocerylist)
	// Update groceries
	grocerylist.Groceries = []Grocery{
		{"Beer", "1"},
		{"More Beer", "2"},
		{"Even More Beer", "3"},
	}

	var err error

	grocerylist, err = m.UpdateGroceryList(grocerylist)

	// Test we didn't get an error back
	if err != nil {
		t.Errorf("Grocerylist could not be found %q", err)
	}
	// Test length of the groceries
	if len(grocerylist.Groceries) == 0 {
		t.Error("boom")
	}
}
