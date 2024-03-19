package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func connectDb() *sql.DB {

	if err := godotenv.Load(); err != nil {
		log.Println("Error loading.env file")
	}

	var err error

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		log.Println("error occured", err)
	} 

	if err := db.Ping(); err != nil {
		log.Println("Error connecting to the database:", err)
	} else {
		log.Println("connected to the database")
	}

	return db
}
