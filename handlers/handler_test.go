package handlers

import (
	"encoding/json"
	"github.com/kristiannissen/grocery-monkey-golang/models"
	"github.com/labstack/echo/v4"
	_ "log"
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

	if rec.Code != 201 {
		t.Errorf("Status %d Body %q", rec.Code, body)
	}
}

func TestAPICreateGroceryList(t *testing.T) {
	t.Skip("Not implemented yet")
}

func TestAPIUpdateGroceryList(t *testing.T) {
	t.Skip("Not implemented yet")
}
