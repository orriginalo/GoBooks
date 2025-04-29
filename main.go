package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

const DB_HOST = "db" // "db" - если запущено в Docker-контейнере
const DB_USERNAME = "postgres"
const DB_PASSWORD = "postgres"
const DB_NAME = "postgres"
const DB_PORT = 5432

func initDB() {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v", DB_HOST, DB_USERNAME, DB_PASSWORD, DB_NAME, DB_PORT)
	var err error

	db, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatalf("не удалось подключится к базе данных: %v", err)
	}

	if err := db.AutoMigrate(&Book{}); err != nil {
		log.Fatalf("не удалось произвести миграцию: %v", err)
	}
}

type Book struct {
	ID     uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

type AddBook struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func deleteBook(c echo.Context) error {
	id := c.Param("id") // Получаем ID из URL

	// Ищем книгу в списке
	if err := db.Delete(&Book{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "не удалось удалить книгу"})
	}
	return c.NoContent(http.StatusNoContent)
}

func getBooks(c echo.Context) error {
	var books []Book
	if err := db.Find(&books).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "не удалось получить книги"})
	}

	return c.JSON(http.StatusOK, books)
}

func addBook(c echo.Context) error {
	var newBook AddBook
	if err := c.Bind(&newBook); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "неправильная форма данных"})
	}

	// Проверка обязательных полей
	if newBook.Title == "" || newBook.Author == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "title и author обязательны"})
	}

	book := Book{
		Title:  newBook.Title,
		Author: newBook.Author,
		Year:   newBook.Year,
	}

	if err := db.Create(&book).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "не удалось создать книгу"})
	}

	return c.JSON(http.StatusCreated, book)
}

func main() {
	initDB()

	e := echo.New()

	// Настройки CORS с явным разрешением всех методов
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"}, // Разрешить все домены (для разработки)
	// 	AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodOptions},
	// }))

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/api/books", getBooks)
	e.POST("/api/books", addBook)
	e.DELETE("/api/books/:id", deleteBook)

	// Запускаем на 0.0.0.0, чтобы работало и с localhost, и с 127.0.0.1
	e.Logger.Fatal(e.Start("0.0.0.0:8080"))
}
