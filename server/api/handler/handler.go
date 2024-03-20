package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gerismumo/golang-todo/server/internal/model"
	"github.com/gerismumo/golang-todo/server/api/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Todo struct {
	Task string `json:"task"`
}


func Routes() {

	//routes

	router := mux.NewRouter()

	//routes
	router.HandleFunc("/api/addTodo", model.AddTodo).Methods("POST")
	router.HandleFunc("/api/getTodo", model.GetTodo).Methods("GET")
	router.HandleFunc("/api/deleteTodo/{id}", model.DeleteTodo).Methods("DELETE")
	router.HandleFunc("/api/editTodo/{id}", model.EditTodo).Methods("PUT")

	//cors handler

	corsHandler := middleware.AllowOnlyFrom("http://localhost:3000", router)
	corsHandler = handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(corsHandler)

	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}


