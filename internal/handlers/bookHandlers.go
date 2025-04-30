package handlers

import (
	"gobooks/internal/bookService"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	service bookService.BookService
}

func (h *BookHandler) DeleteBook(c echo.Context) error {
	id := c.Param("id")

	if err := h.service.DeleteBook(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "не удалось удалить книгу"})
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *BookHandler) GetBooks(c echo.Context) error {
	books, err := h.service.GetAllBooks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "не удалось получить книги"})
	}
	return c.JSON(http.StatusOK, books)
}

func (h *BookHandler) AddBook(c echo.Context) error {
	var newBook bookService.AddBook
	if err := c.Bind(&newBook); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "неверный запрос"})
	}

	if newBook.Title == "" || newBook.Author == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "title и author обязательны"})
	}

	book, err := h.service.CreateBook(newBook.Title, newBook.Author, newBook.Year)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "не удалось создать книгу"})
	}

	return c.JSON(http.StatusCreated, book)
}

func NewBookHandler(s bookService.BookService) *BookHandler {
	return &BookHandler{service: s}
}
