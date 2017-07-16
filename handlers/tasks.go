package handlers

import (
	"../models"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type H map[string]interface{}

// GetTasks endpoint
func GetTasks(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetTasks(db))
	}
}

// PutTask endpoint
func PutTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Instantiate a new task
		var task models.Task

		// Map imcoming JSON body to the new Task
		c.Bind(&task)

		// Add a task using the model
		id, err := models.PutTask(db, task.Name)

		if err == nil {
			// Return a JSON response if successful
			return c.JSON(http.StatusCreated, H{
				"created": id,
			})
		} else {
			// Handle any errors
			return err
		}
	}
}

// DeleteTask endpoint
func DeleteTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		// Use the model to delete a task
		_, err := models.DeleteTask(db, id)

		if err == nil {
			// Return a JSON response on success
			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
		} else {
			// Handle errors
			return err
		}
	}
}
