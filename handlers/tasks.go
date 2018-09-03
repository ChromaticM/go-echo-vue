//tasks.go

package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"scotch/models"

	"github.com/labstack/echo"
)

type H map[string]interface{}

//GetTasks endpoint
func GetTasks(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetTasks(db))
	}
}

//PUtTask endpoint
func PutTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		//Instantiate a new task
		var task models.Task

		//Map incoming JSON body to the new task
		c.Bind(&task)

		//Add a task using our new models
		id, err := models.PutTask(db, task.Name)

		//Return a JSON response if successful
		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"created": id,
			})
		} else { //Handle any errors
			return err
		}
	}
}

func DeleteTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		//id, _ := strconv.Atoi(c.Param("id"))
		//return c.JSON(http.StatusOK, H{
		//	"deleted": id,
		//})

		id, _ := strconv.Atoi(c.Param("id"))
		//Use our new model to delete a tasks
		_, err := models.DeleteTask(db, id)
		//Return a JSON response on success
		if err == nil {
			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
		} else {
			//Handle errors
			return err
		}
	}
}
