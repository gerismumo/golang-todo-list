package main

import (
	"log"
	

	"github.com/gerismumo/golang-todo/server/api/handler"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load("../.env"); err != nil {
		log.Printf("Error loading.env file: %v", err)
	}

	handler.Routes()
}
