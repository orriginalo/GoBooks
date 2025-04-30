package main

import (
	bookS "gobooks/internal/bookService"
	"gobooks/internal/db"
	"gobooks/internal/handlers"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db, err := db.InitDB()
	if err != nil {
		log.Fatalf("Не удалось подключится к БД: %v", err)
	}

	bookRepo := bookS.NewBookRepository(db)
	bookService := bookS.NewBookService(bookRepo)
	bookHandlers := handlers.NewBookHandler(bookService)

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/api/books", bookHandlers.GetBooks)
	e.POST("/api/books", bookHandlers.AddBook)
	e.DELETE("/api/books/:id", bookHandlers.DeleteBook)

	e.Logger.Fatal(e.Start("0.0.0.0:8080"))
}
