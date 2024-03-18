package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	db, err := sql.Open("mysql", "todoList:todoList@(localhost:3306)/todoList?parseTime=true")

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

	router.HandleFunc("/api/addTodo", addTodo).Methods("POST")

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
	log.Println(r.Body)

	log.Println("Hello World!")
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
