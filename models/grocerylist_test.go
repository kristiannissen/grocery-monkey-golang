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

	grocerylist, err := m.CreateGroceryList(grocerylist)

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
	t.Skip("Not implemented yet")
}