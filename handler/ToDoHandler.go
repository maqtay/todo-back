package handler

import (
	"ToDo/models"
	"ToDo/services"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type TodoHandler struct {
	service services.Service
}

type Handler interface {
	GetAll(c echo.Context) error
	Add(c echo.Context) error
	Delete(c echo.Context) error
}

func NewTodoHandler(service services.Service) Handler {
	return TodoHandler{service}
}

func (handler TodoHandler) GetAll(c echo.Context) error {
	return c.JSON(http.StatusOK, handler.service.GetAll())
}

func (handler TodoHandler) Add(c echo.Context) error {
	t := new(models.ToDo)
	if err := c.Bind(t); err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusCreated, handler.service.Add(t.Note))
}

func (handler TodoHandler) Delete(c echo.Context) error {
	idStr := c.QueryParams().Get("id")
	id, _ := strconv.Atoi(idStr)
	handler.service.Delete(id)

	return c.JSON(http.StatusOK, "")
}
