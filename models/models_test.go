package models

import (
    "testing"
)

var (
    m Model = Model{}
)

// Setup
func init() {
}

func TestCreateUser(t *testing.T) {

    user := m.CreateUser("Kitty")

    if user.NickName != "Kitty" {
        t.Errorf("Test CreateUser: Got %s - want %s", user.NickName, "Kitty")
    }
}

func TestUser(t *testing.T) {
    tests := []struct {
        name, got, want string
    }{
        {"GetUser", "Kitty", "Kitty"},
    }

    for _, tt := range tests {
        if m.GetUser(tt.got).NickName != tt.want {
            t.Errorf("Test %s: Got %s - want %s", tt.name, m.GetUser(tt.got).NickName, tt.want)
        }
    }
    // m.TearDown()
}

func TestCreateGroceryList(t *testing.T) {
    t.Skip("Not implemented yet")
}

func TestUpdateGroceryList(t *testing.T) {
    t.Skip("Not implemented yet")
}

func TestGetGroceryList(t *testing.T) {
    t.Skip("Not implemented yet")
}
