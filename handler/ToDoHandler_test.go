package handler

import (
	"ToDo/models"
	"ToDo/services"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetAll(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	var todos []models.ToDo
	todos = append(todos, models.ToDo{
		Id:        1,
		Note:   "Hi Guys",
	})

	service := services.NewMockService(controller)
	service.EXPECT().GetAll().Return(todos).Times(1)
	handler := TodoHandler{service}

	todoJSON := `[{"Id":1,"note":"Hi Guys","Date":""}]
`
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/getalltodo", nil)
	req.Header.Set(echo.HeaderAccept, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handler.GetAll(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, todoJSON, rec.Body.String())
	}
}

func TestAdd(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	requestAsString := `{"Note":"Hi Guys"}`
	expected := models.ToDo{
		Id:        1,
		Note:   "Hi Guys",
	}
	todoJSON := `{"Id":1,"note":"Hi Guys","Date":""}
`

	service := services.NewMockService(controller)
	service.EXPECT().Add("Hi Guys").Return(expected).Times(1)

	handler := TodoHandler{service}
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/addtodo", strings.NewReader(requestAsString))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handler.Add(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, todoJSON, rec.Body.String())
	}
}

func TestDelete(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	service := services.NewMockService(controller)
	service.EXPECT().Delete(1)
	handler := TodoHandler{service}

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/deletetodo", nil)
	q := req.URL.Query()
	q.Add("id", "1")
	req.URL.RawQuery = q.Encode()

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handler.Delete(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

