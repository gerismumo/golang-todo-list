package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func editTodo(w http.ResponseWriter, r *http.Request) {
	db := connectDb()

	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	var todo Todo
	json.NewDecoder(r.Body).Decode(&todo)

	if id == "" || todo.Task == "" {
		response := struct {
			Success bool   `json:"success"`
			Message string `json:"message"`
		}{
			Success: false,
			Message: "Fill all the fields",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

		return
	}

	result, err := db.Exec("UPDATE todo_list SET name =? WHERE id = ?", todo.Task, id)

	if err != nil {
		response := struct {
			Success bool   `json:"success"`
			Message string `json:"message"`
		}{
			Success: false,
			Message: "Server failed",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		response := struct {
			Success bool   `json:"success"`
			Message string `json:"message"`
		}{
			Success: false,
			Message: "unsuccessful submission",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	if rowsAffected > 0 {
		response := struct {
			Success bool   `json:"success"`
			Message string `json:"message"`
		}{
			Success: true,
			Message: " edited successfully",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

		return
	}
}
