package main

import (
	"encoding/json"
	"net/http"
)

func getTodo(w http.ResponseWriter, r *http.Request) {
	db := connectDb()
	defer db.Close()

	type QueryTodo struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		CreatedAt string `json:"createdAt"`
	}

	rows, err := db.Query("SELECT * FROM todo_list")

	if err != nil {
		response := struct {
			Success bool        `json:"success"`
			Message string      `json:"message"`
			Data    []QueryTodo `json:"data"`
		}{
			Success: false,
			Message: "Server failed",
			Data:    nil,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}

	defer rows.Close()

	var todos []QueryTodo
	for rows.Next() {
		var queryTodo QueryTodo

		err := rows.Scan(&queryTodo.ID, &queryTodo.Name, &queryTodo.CreatedAt)

		if err != nil {
			response := struct {
				Success bool        `json:"success"`
				Message string      `json:"message"`
				Data    []QueryTodo `json:"data"`
			}{
				Success: false,
				Message: "Server failed",
				Data:    nil,
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}
		todos = append(todos, queryTodo)
	}

	response := struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Data    []QueryTodo `json:"data"`
	}{
		Success: true,
		Message: "successfully executed",
		Data:    todos,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}