package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("err loading: %v", err)
    }
    dsn := os.Getenv("DATABASE_URL")
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    return db
}