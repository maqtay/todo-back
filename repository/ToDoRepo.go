package repository

import (
	"ToDo/models"
	"context"
	"database/sql"
	_ "database/sql"
	"fmt"
	"time"
)

type Repository interface {
	GetAll() interface{}
	Add(note string) interface{}
	Delete(id int) interface{}
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
	todo := new(models.ToDo)
	stmt, err := ts.db.Prepare("INSERT INTO ToDoList(todo, createdDate) VALUES(?,?)")
	if err != nil {
		panic(err.Error())
	}
	res, err2 :=  stmt.Exec(note, time.Now())
	if err2 != nil {
		panic(err2.Error())
	}
	fmt.Println(res)
	return todo
}

func (ts *ToDoStore) Delete(id int) interface{} {
	rmvDB, err := ts.db.Prepare("DELETE FROM todolist WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	res, _ := rmvDB.Exec(id)
	i, err := res.RowsAffected()
	if i == 0 {
		return nil
	}
	return true
}