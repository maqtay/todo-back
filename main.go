package main

import (
	"ToDo/handler"
	"ToDo/repository"
	"ToDo/services"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	_ "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/labstack/echo/v4/middleware"
)

func dbConn() (db *sql.DB, err error)  {
	dbDriver := "mysql"
	dbUser := "todo"
	dbPass := "todoback"
	dbName := "todoDB"
	db, _ = sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(localhost:3306)/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db, nil
}

func createTable() error{
	db, err := dbConn()
	if err != nil {

	}
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
	res, err := db.Exec(sql)
	if err != nil {
		panic(err.Error())
	}
	_, err1 = res.RowsAffected()
	if err1 != nil {
		panic(err1.Error())
	}
	return nil
}

func main() {
	err := createTable()
	if err != nil {
		return 
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	db, err := dbConn()
	if err != nil {
		panic(err.Error())
	}

	todoRepo := repository.NewToDoStore(db)
	todoService := services.NewTodoService(todoRepo)
	todohandler := handler.NewTodoHandler(todoService)

	e.POST("/addtodo", todohandler.Add)
	e.GET("/getalltodo", todohandler.GetAll)
	e.DELETE("/deletetodo", todohandler.Delete)

	e.Logger.Fatal(e.Start(":5858"))
}