package repository

import (
	"ToDo/models"
	"context"
	"database/sql"
	_ "database/sql"
	"time"
)

type Repository interface {
	GetAll() interface{}
	Add(note string) interface{}
	Delete(id int)
}

type ToDoStore struct {
	db *sql.DB
}

func NewToDoStore(db *sql.DB) *ToDoStore {
	return &ToDoStore{
		db: db,
	}
}

func (ts *ToDoStore) GetAll() interface{} {
	var sql = "SELECT * FROM todoList ORDER BY createdDate"
	var todos []models.ToDo
	res, _ := ts.db.QueryContext(context.Background(), sql)
	for res.Next() {
		var todo models.ToDo
		err := res.Scan(&todo.Id, &todo.Note, &todo.Date)
		todos = append(todos, todo)
		if err != nil {
			panic(err.Error())
		}
	}
	return todos
}

func (ts *ToDoStore) Add(note string) interface{} {
	todo := models.ToDo{
		Note: note,
	}

	stmt, err := ts.db.PrepareContext(context.Background(),"INSERT INTO ToDoList(todo, createdDate) VALUES(?,?)")
	if err != nil {
		panic(err.Error())
	}
	_, err2 :=  stmt.Exec(todo.Note, time.Now())
	if err2 != nil {
		panic(err2.Error())
	}
	return todo
}

func (ts *ToDoStore) Delete(id int) {
	rmvDB, err := ts.db.Prepare("DELETE FROM todolist WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	rmvDB.Exec(id)
}