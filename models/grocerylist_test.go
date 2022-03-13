package models

import (
    "testing"
    _ "log"
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
    t.Skip("Not implemented yet")
}

func TestUpdateGroceryList(t *testing.T) {
    t.Skip("Not implemented yet")
}
