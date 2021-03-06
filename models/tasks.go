package models

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// Task is a struct containing Task data
type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// TaskCollection is collection of Tasks
type TaskCollection struct {
	Tasks []Task `json:"items"`
}

func GetTasks(db *sql.DB) TaskCollection {
	sql := "SELECT * FROM tasks"
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	// make sure to cleanup when the program exits
	defer rows.Close()

	result := TaskCollection{}
	for rows.Next() {
		// create task for each row
		task := Task{}
		err2 := rows.Scan(&task.ID, &task.Name)
		if err2 != nil {
			panic(err2)
		}

		result.Tasks = append(result.Tasks, task)
	}

	return result
}

func PutTask(db *sql.DB, name string) (int64, error) {
	sql := "INSERT INTO tasks(name) VALUES(?)"

	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	// Make sure to cleanup after the program exits
	defer stmt.Close()

	// Replace the '?' in our prepared statement with 'name'
	result, err2 := stmt.Exec(name)
	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}

func DeleteTask(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM tasks WHERE id = ?"

	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}

	// Replace the '?' in our prepared statement with 'id'
	result, err2 := stmt.Exec(id)
	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}
