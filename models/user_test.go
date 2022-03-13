package models

import (
    "testing"
    _ "log"
)

var (
    m *Model
)

func init() {
   m.UserSetUp() 
}

func TestGetUserNotFound(t *testing.T) {

    m.CleanUserDatabase()

    _, err := m.GetUser("Pussy")
    if err == nil {
        t.Fatal("None existing user found")
    }

    t.Cleanup(func() {
        // m.UserTearDown()
    })
}

func TestCreateUser(t *testing.T) {
    // m.UserSetUp()

    user := m.CreateUser("Pussy")

    if user.NickName != "Pussy" {
        t.Fatalf("Test CreateUser: Wanted %s - Got %s", "Pussy", user.NickName)
    }
}

func TestGetUserFound(t *testing.T) {

    user, err := m.GetUser("Pussy")

    if err != nil {
        t.Fatalf("Found %q", user)
    }
}
