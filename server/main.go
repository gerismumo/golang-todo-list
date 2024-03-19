package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Todo struct {
	Task string `json:"task"`
}



var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "todoList:todoList@(localhost:3306)/todoList?parseTime=true")

	if err != nil {
		log.Println("Error connecting to the database:", err)
		// log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Println("Error pinging the database:", err)
	} else {
		log.Println("connected to the database")
	}

	//routes

	router := mux.NewRouter()

	//routes
	router.HandleFunc("/api/addTodo", addTodo).Methods("POST")
	router.HandleFunc("/api/getTodo", getTodo).Methods("GET")
	router.HandleFunc("/api/deleteTodo/{id}", deleteTodo).Methods("DELETE")
	router.HandleFunc("/api/editTodo/{id}", editTodo).Methods("PUT")

	//cors handler

	corsHandler := allowOnlyFrom("http://localhost:3000", router)
	corsHandler = handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(corsHandler)

	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}

func addTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	json.NewDecoder(r.Body).Decode(&todo)

	// if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	defer r.Body.Close()

	if todo.Task == "" {
		response := struct {
			Success bool   `json:"success"`
			Message string `json:"message"`
		}{
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
			Message: "successfully created",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	type QueryTodo  struct {
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


func deleteTodo(w http.ResponseWriter, r *http.Request) {
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

	if err!= nil {
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

	if err!= nil {
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

	log.Println("deleteTodo", rowsAffected);

	if(rowsAffected > 0) {
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


func editTodo(w http.ResponseWriter, r *http.Request) {

}

func allowOnlyFrom(allowedDomain string, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		if origin != allowedDomain {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		handler.ServeHTTP(w, r)
	})
}
