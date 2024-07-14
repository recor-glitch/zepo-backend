package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/recor-glitch/zepo-backend/internal/domain/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("err loading: %v", err)
    }
    dsn := os.Getenv("DATABASE_URL")
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
    })
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    err = db.AutoMigrate(&user.User{})
    if err != nil {
        log.Fatalf("Failed to migrate database: %v", err)
    }

    return db
}