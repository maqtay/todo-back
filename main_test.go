package main

import (
	"ToDo/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	mockDB = map[string]*models.ToDo{
		"Hi Guys": &models.ToDo{Id:1, Note:"Hi Guys", Date:"2021-04-16 14:30:33.495523"},
	}

 	todoJSON = `{"Id":1,"todo":"Hi Guys","Date":"2021-04-16 14:30:33.495523"}`

 	expectedString = "\"Hi Guys\"\n"

 	deleteTodoId = "83"
)

func TestAdd(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest(http.MethodPost, "/addtodo",strings.NewReader(todoJSON))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Errorf("The request could not be created because of: %v", err)
	}
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	res := rec.Result()
	defer res.Body.Close()

	if assert.NoError(t, Add(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, expectedString, rec.Body.String())
	}
}

func TestGetAll(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/getalltodo", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, GetAll(c)){
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), todoJSON)
	}
}

func TestDeleteToDo(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest(http.MethodDelete, "/deletetodo", nil)
	q := req.URL.Query()
	q.Add("id", deleteTodoId)
	req.URL.RawQuery = q.Encode()

	if err != nil {
		t.Errorf("The request could not be created because of: %v", err)
	}
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, DeleteToDo(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "\"Todo has been deleted!\"\n", rec.Body.String())
	}
}