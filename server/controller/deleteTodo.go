package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gerismumo/golang-todo/server/connect"
)



func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	db := connect.ConnectDb()

	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		response := responseData{
			Success: false,
			Message: "ID cannot be empty",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

		return
	}

	result, err := db.Exec("DELETE FROM todo_list WHERE id =?", id)

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
			Message: "successfully deleted",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

		return
	}
}
