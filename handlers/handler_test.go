package handlers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	_ "log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	userJSON = `{"nickname": "Hello Kitty"}`
)

type (
	Response struct {
		Token string `json:"token"`
	}
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

func TestAPIAuthenticateFailed(t *testing.T) {

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/authenticate", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := &Handler{}
	h.Authenticate(c)

	res := Response{}
	json.Unmarshal([]byte(rec.Body.String()), &res)

	if rec.Code != 401 {
		t.Error("JWT token not returned")
	}
}

func TestAPIAuthenticateSuccess(t *testing.T) {
	// time.Now().UnixNano() to create a random number
	t.Skip("Not implemented yet")
}

func TestAPICreateGroceryList(t *testing.T) {
	t.Skip("Not implemented yet")
}

func TestAPIUpdateGroceryList(t *testing.T) {
	t.Skip("Not implemented yet")
}
