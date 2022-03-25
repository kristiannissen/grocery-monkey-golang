package models

import (
	_ "log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	model := Model{}
	model.UserSetUp()
	model.GroceryListSetUp()

	code := m.Run()

	os.Exit(code)

	model.UserTearDown()
	model.GroceryListTearDown()
}
