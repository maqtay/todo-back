package services

import (
	"ToDo/models"
	"ToDo/repository"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAll(t *testing.T)  {
	controller := gomock.NewController(t)
	defer controller.Finish()

	var todos []models.ToDo
	todos = append(todos, models.ToDo{
		Id: 1,
		Note: "Hi Guys!",
	})
	repo := repository.NewMockRepo(controller)
	repo.EXPECT().GetALl().Return(todos).Times(1)
	service := ToDoService{repo}

	getall := service.GetAll().([]models.ToDo)
	assert.NotNil(t, getall)
	assert.NotEmpty(t, getall)
	assert.Equal(t, 1, len(getall))
	assert.Equal(t, 1, getall[0].Id)
	assert.Equal(t, "Hi Guys", getall[0].Note)
}

func TestAdd(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repo := repository.NewMockRepo(controller)
	repo.EXPECT().Add("Hi Guys").Return(models.ToDo{
		Id:        1,
		Note:   "Hi Guys",
	}).Times(1)

	service := NewTodoService(repo)
	createdTodo := service.Add("Hi Guys")

	assert.NotNil(t, createdTodo)
	todo := createdTodo.(models.ToDo)
	assert.Equal(t, 1, todo.Id)
	assert.Equal(t, "Buy some milk", todo.Note)
}

func TestDelete(t *testing.T) {
	controller := gomock.NewController(t)

	repo := repository.NewMockRepo(controller)
	repo.EXPECT().Delete(1).Times(1)

	service := ToDoService{repo}
	service.Delete(1)
}
