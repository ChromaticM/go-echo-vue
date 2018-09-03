// todo.go
package main

//
//database "database/sql"
import (
	"database/sql"
	"scotch/handlers"

	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	//Here we check for any db errors then exit
	if err != nil {
		panic(err)
	}
	//If we don't gey any erros but somehow still don't get a db connecction
	//we exit as well.
	if db == nil {
		panic("db nil")
	}
	return db
}

func migrate(db *sql.DB) {
	sql := "CREATE TABLE IF NOT EXISTS tasks(id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, name VARCHAR NOT NULL);"
	_, err := db.Exec(sql)
	//Exit if something goes wrong with our SQL statement above
	if err != nil {
		panic(err)
	}
}

func main() {
	db := initDB("storage.db")
	migrate(db)

	// Create a new instance of Echo
	e := echo.New()
	e.File("/", "public/index.html")
	e.GET("/tasks", handlers.GetTasks(db))
	e.PUT("/tasks", handlers.PutTask(db))
	e.DELETE("/tasks/:id", handlers.DeleteTask(db))
	//e.GET("/tasks", func(c echo.Context) error { return c.JSON(200, handlers.GetTasks(db)) })
	//e.PUT("/tasks", func(c echo.Context) error { return c.JSON(200, handlers.PutTask(db)) })
	//	e.GET("/tasks/:id", func(c echo.Context) error { return c.JSON(200, "GET Task "+c.Param("id")) })
	//e.DELETE("/tasks/:id", func(c echo.Context) error { return c.JSON(200, handlers.DeleteTask(db)) })

	// Start as a web server
	e.Logger.Fatal(e.Start(":8000"))
}
