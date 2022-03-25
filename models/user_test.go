package models

import (
	_ "log"
	"testing"
)

var (
	m *Model
)

func init() {
}

func TestGetUserNotFound(t *testing.T) {
	_, err := m.GetUser("Pussy")

	if err != nil {
		t.Error("Not existing user found")
	}
}

func TestCreateUser(t *testing.T) {

	user, _ := m.CreateUser("Pussy")

	if user.NickName != "Pussy" {
		t.Errorf("Test CreateUser: Wanted %s - Got %s", "Pussy", user.NickName)
	}
}

func TestGetUserFound(t *testing.T) {
	// Create user
	m.CreateUser("Pussy")

	user, err := m.GetUser("Pussy")

	if err != nil {
		t.Errorf("Found %q", user)
	}
}
