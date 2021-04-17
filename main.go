package main

import (
	"ToDo/models"
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/labstack/echo/v4"
	_ "github.com/labstack/echo/v4/middleware"
)

func dbConn() (db *sql.DB)  {
	dbDriver := "mysql"
	dbUser := "todo"
	dbPass := "todoback"
	dbName := "todoDB"
	db, err := sql.Open(dbDriver, dbUser + ":" + dbPass + "@/" + dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func createTable() {
	db := dbConn()
	createDatabase := "CREATE DATABASE IF NOT EXISTS todoDB;"
	_, err1 := db.Exec(createDatabase)
	if err1 != nil {
		panic(err1.Error())
	}
	sql := "CREATE TABLE IF NOT EXISTS `todolist` (" +
		"  `id` int(11) NOT NULL AUTO_INCREMENT," +
		"  `todo` text COLLATE utf8mb4_unicode_520_ci NOT NULL," +
		"  `createdDate` varchar(100) COLLATE utf8mb4_unicode_520_ci DEFAULT NULL," +
		"  PRIMARY KEY (`id`)" +
		") ENGINE=InnoDB AUTO_INCREMENT=76 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci"
	_, err := db.Exec(sql)
	if err != nil {
		panic(err.Error())
	}
}

func Add(c echo.Context) error {
	db := dbConn()
	todo := new(models.ToDo)
	todoText := c.FormValue("todo")
	stmt, err := db.Prepare("INSERT INTO ToDoList(todo, createdDate) VALUES(?,?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	res, err2 := stmt.Exec(todoText, time.Now())
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(res.LastInsertId())
	return c.JSON(http.StatusCreated, todo.Note)
}

func GetAll(c echo.Context) error {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM todoList")
	if err != nil {
		panic(err.Error())
	}
	todo := models.ToDo{}
	var todos []models.ToDo
	for selDB.Next() {
		err = selDB.Scan(&todo.Id, &todo.Note, &todo.Date)
		todos = append(todos, todo)
		if err != nil {
			panic(err.Error())
		}
	}
	defer db.Close()
	return c.JSON(http.StatusOK, todos)
}

func DeleteToDo(e echo.Context) error {
	db := dbConn()
	id := e.QueryParams().Get("id")
	if id == "" {
		return e.JSON(http.StatusBadRequest, "Please, insert the correct paramaters!")
	}
	rmvDB, err := db.Prepare("DELETE FROM todoList WHERE id = ?")
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	res, _ := rmvDB.Exec(id)
	i, err := res.RowsAffected()
	if i == 0 {
		return e.JSON(http.StatusInternalServerError, "Any field was not affected. Please try again with a different parameter.")
	}
	defer db.Close()
	return e.JSON(http.StatusOK, "Todo has been deleted!")
}

func main() {
	createTable()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.POST("/addtodo", Add)
	e.GET("/getalltodo", GetAll)
	e.DELETE("/deletetodo", DeleteToDo)

	e.Logger.Fatal(e.Start(":5959"))
}