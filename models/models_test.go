package models

import (
    "testing"
)

// Setup
func init() {
    
}

func TestCreateUser(t *testing.T) {
    m := Model{}
    m.SetUp()

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

    m := Model{}
    m.SetUp()

    for _, tt := range tests {
        if m.GetUser(tt.got).NickName != tt.want {
            t.Errorf("Test %s: Got %s - want %s", tt.name, m.GetUser(tt.got).NickName, tt.want)
        }
    }
    // m.TearDown()
}
