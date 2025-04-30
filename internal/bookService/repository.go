package bookService

import "gorm.io/gorm"

type BookRepository interface {
	CreateBook(book Book) error
	GetAllBooks() ([]Book, error)
	DeleteBook(id string) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) CreateBook(book Book) error {
	return r.db.Create(&book).Error
}

func (r *bookRepository) GetAllBooks() ([]Book, error) {
	var books []Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *bookRepository) DeleteBook(id string) error {
	return r.db.Delete(&Book{}, "id = ?", id).Error
}
