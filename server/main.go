package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gerismumo/golang-todo/server/controller"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Todo struct {
	Task string `json:"task"`
}

type responseData struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func main() {

	//routes

	router := mux.NewRouter()

	//routes
	router.HandleFunc("/api/addTodo", controller.AddTodo).Methods("POST")
	router.HandleFunc("/api/getTodo", controller.GetTodo).Methods("GET")
	router.HandleFunc("/api/deleteTodo/{id}", controller.DeleteTodo).Methods("DELETE")
	router.HandleFunc("/api/editTodo/{id}", controller.EditTodo).Methods("PUT")

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
