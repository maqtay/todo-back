package services

import "ToDo/repository"

type Service interface {
	GetAll() interface{}
	Delete(id int)
	Add(note string) interface{}
}

type ToDoService struct {
	repository repository.Repository
}

func NewTodoService(repository repository.Repository) ToDoService {
	return ToDoService{repository}
}

func (t ToDoService) GetAll() interface{} {
	return t.repository.GetAll()
}

func (t ToDoService) Delete(id int) {
	t.repository.Delete(id)
}

func (t ToDoService) Add(note string) interface{} {
	return t.repository.Add(note)
}