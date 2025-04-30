package db

import (
	"fmt"
	"gobooks/internal/bookService"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

const DB_HOST = "db" // "db" - если запущено в Docker-контейнере
const DB_USERNAME = "postgres"
const DB_PASSWORD = "postgres"
const DB_NAME = "postgres"
const DB_PORT = 5432

func InitDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v", DB_HOST, DB_USERNAME, DB_PASSWORD, DB_NAME, DB_PORT)
	var err error

	db, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatalf("не удалось подключится к базе данных: %v", err)
	}

	if err := db.AutoMigrate(&bookService.Book{}); err != nil {
		log.Fatalf("не удалось произвести миграцию: %v", err)
	}

	return db, nil
}
