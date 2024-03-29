package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Initalize variable
var db *sql.DB

// Init/Setup DB Connection
func InitDB() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Connect DB Postgres
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort))

	if err != nil {
		log.Fatal(err)
	}

	// DB Pooling
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connect DB Success")
}

// Get Connection DB
func GetDB() *sql.DB {
	return db
}
