package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)


func deleteTodo(w http.ResponseWriter, r *http.Request) {
	db := connectDb()
	
	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		response := struct {
			Success bool   `json:"success"`
			Message string `json:"message"`
		}{
			Success: false,
			Message: "ID cannot be empty",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

		return
	}

	result, err := db.Exec("DELETE FROM todo_list WHERE id =?", id)

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
			Message: "successfully deleted",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

		return
	}
}