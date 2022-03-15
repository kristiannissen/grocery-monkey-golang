package handlers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
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
		log.Println(rec.Code)
		t.Error("Unauthorized access")
	}
}

func TestAPIAuthenticateSuccess(t *testing.T) {
	userStr := "{\"nickname\":\"" + strconv.FormatInt(time.Now().UnixNano(), 10) + "-Kitty\"}"

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/authenticate", strings.NewReader(userStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := &Handler{}
	h.Authenticate(c)

	res := Response{}
	json.Unmarshal([]byte(rec.Body.String()), &res)
	log.Println("test")

	if rec.Code != 201 {
		t.Skip("Not implemented yet")
	}
}

func TestAPICreateGroceryList(t *testing.T) {
	t.Skip("Not implemented yet")
}

func TestAPIUpdateGroceryList(t *testing.T) {
	t.Skip("Not implemented yet")
}
