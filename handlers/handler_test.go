package handlers

import (
  "testing"
  "net/http"
  "net/http/httptest"
  _ "log"
  "github.com/labstack/echo/v4"
)

func TestIndex(t *testing.T) {

  e := echo.New()
  req := httptest.NewRequest(http.MethodGet, "/", nil)
  rec := httptest.NewRecorder()
  c := e.NewContext(req, rec)

  h := &Handler{}

  h.Index(c)

  if rec.Body.String() != "Hello Kitty" {
    t.Error("Index did not respond with Hello Kitty")
  }
}

func TestAPIAuthenticate(t *testing.T) {
  t.Skip("Not implemented yet")
}

func TestAPICreateUser(t *testing.T) {
  t.Skip("Not implemented yet")
}

func TestAPICreateGroceryList(t *testing.T) {
  t.Skip("Not implemented yet")
}

func TestAPIUpdateGroceryList(t *testing.T) {
  t.Skip("Not implemented yet")
}
