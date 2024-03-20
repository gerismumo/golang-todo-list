package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gerismumo/golang-todo/server/connect"
)

func EditTodo(w http.ResponseWriter, r *http.Request) {
	db := connect.ConnectDb()

	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	var todo Todo
	json.NewDecoder(r.Body).Decode(&todo)

	if id == "" || todo.Task == "" {
		response := responseData{
			Success: false,
			Message: "Fill all the fields",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

		return
	}

	result, err := db.Exec("UPDATE todo_list SET name =? WHERE id = ?", todo.Task, id)

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
			Message: " edited successfully",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

		return
	}
}
