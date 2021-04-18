package services

import "ToDo/repository"

type Service interface {
	GetAll() interface{}
	Delete(id int) interface{}
	Add(note string) interface{}
}

type ToDoService struct {
	repository repository.Repository
}
func NewTodoService(repository repository.Repository) Service {
	return ToDoService{repository}
}

func (t ToDoService) GetAll() interface{} {
	return t.repository.GetAll()
}

func (t ToDoService) Delete(id int) interface{} {
	return t.repository.Delete(id)
}

func (t ToDoService) Add(note string) interface{} {
	return t.repository.Add(note)
}