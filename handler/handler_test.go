package handler

import (
	"encoding/json"
	"github.com/kristiannissen/grocery-monkey-golang/models"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
)

type (
	Response struct {
		Token string `json:"token"`
		Body  string
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

func TestAPIAuthenticateNewUser(t *testing.T) {
	m := &models.Model{}
	m.CleanUserTable()

	userStr := "{\"nickname\":\"" + strconv.FormatInt(time.Now().UnixNano(), 10) + "-Kitty\"}"

	e := echo.New()
	req := httptest.NewRequest(
		http.MethodPost, "/api/authenticate", strings.NewReader(userStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := &Handler{}
	h.Authenticate(c)

	res := Response{}
	body := rec.Body.String()
	json.Unmarshal([]byte(body), &res)

	if rec.Code != 200 {
		t.Errorf("Status %d Body %q", rec.Code, body)
	}
}

func TestAPICreateGroceryList(t *testing.T) {
	groceryData := `{
    "uuid": "e3357dac-a275-41f7-87ef-069d91de3c9e",
    "subscribers": [
      "bc21cf88-3a4b-4e9d-a34e-35090af0d165"
    ],
    "groceries": [],
    "useruuid": "bc21cf88-3a4b-4e9d-a34e-35090af0d165"
  },
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuaWNrbmFtZSI6IjE2NDc0MzEyMjcyOTE5NTc3MDAtS2l0dHkiLCJ1dWlkIjoiYmMyMWNmOD
gtM2E0Yi00ZTlkLWEzNGUtMzUwOTBhZjBkMTY1IiwiZXhwIjoxNjQ3NjkwNDI3fQ.J2lTWXd6P0Cfk8lJmMXFK1uKVuLUHsiibmGVJf6T-LI"
}`

	e := echo.New()
	req := httptest.NewRequest(
		http.MethodPost, "/api/groceries", strings.NewReader(groceryData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := &Handler{}
	h.CreateGroceryList(c)
	body := rec.Body.String()

	if rec.Code != 201 {
		log.Println(body)
		t.Errorf("Want 201 - got %d", rec.Code)
	}
}

func TestAPIUpdateGroceryList(t *testing.T) {
	groceryData := `{
    "uuid": "e3357dac-a275-41f7-87ef-069d91de3c9e",
    "subscribers": [
      "bc21cf88-3a4b-4e9d-a34e-35090af0d165"
    ],
    "groceries": [{"name": "Beer", "id": "1"}],
    "useruuid": "bc21cf88-3a4b-4e9d-a34e-35090af0d165"
  }`

	e := echo.New()
	req := httptest.NewRequest(
		http.MethodPut, "/api/groceries", strings.NewReader(groceryData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := &Handler{}
	h.UpdateGroceryList(c)
	body := rec.Body.String()

	if rec.Code != 200 {
		log.Println(body)
		t.Errorf("Want 200 - got %d", rec.Code)
	}
}
