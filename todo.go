package main

import (
	"./handlers"
	"database/sql"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("Activating TODO app (vue-echo-spa)")

	// set DB
	db := initDB("storage.db")
	migrate(db)

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// static files
	e.File("/", "public/index.html")
	e.Static("/public/js", "public/js")

	// Route => handler
	e.GET("/tasks", handlers.GetTasks(db))
	e.PUT("/tasks", handlers.PutTask(db))
	e.DELETE("/tasks/:id", handlers.DeleteTask(db))

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	// Here we check for any db errors then exit
	if err != nil {
		panic(err)
	}

	// If we don't get any errors but somehow still don't get a db connection
	// we exit as well
	if db == nil {
		panic("db nil")
	}
	return db
}

func migrate(db *sql.DB) {
	sql := `
    CREATE TABLE IF NOT EXISTS tasks(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        name VARCHAR NOT NULL
    );
    `

	_, err := db.Exec(sql)
	// Exit if something goes wrong with our SQL statement above
	if err != nil {
		panic(err)
	}
}
