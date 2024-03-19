package main

import (
	"encoding/json"
	"net/http"
)



func addTodo(w http.ResponseWriter, r *http.Request) {
	db := connectDb()
	defer db.Close()

	var todo Todo
	json.NewDecoder(r.Body).Decode(&todo)

	defer r.Body.Close()

	if todo.Task == "" {
		response := responseData{
			Success: false,
			Message: "Task cannot be empty",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

		return
	}

	//insert data into database
	result, err := db.Exec("INSERT INTO todo_list (name) VALUES (?)", todo.Task)

	if err != nil {
		response := responseData{
			Success: false,
			Message: "Server failed",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		response := responseData{
			Success: false,
			Message: "unsuccessful submission",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	if rowsAffected > 0 {
		response := responseData{
			Success: true,
			Message: "successfully created",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
