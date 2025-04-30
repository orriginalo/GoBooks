package bookService

type BookService interface {
	CreateBook(title string, author string, year int) (Book, error)
	GetAllBooks() ([]Book, error)
	DeleteBook(id string) error
}

type bookService struct {
	repo BookRepository
}

func (s *bookService) CreateBook(title string, author string, year int) (Book, error) {
	newBook := Book{
		Title:  title,
		Author: author,
		Year:   year,
	}
	if err := s.repo.CreateBook(newBook); err != nil {
		return Book{}, err
	}
	return newBook, nil
}

func (s *bookService) GetAllBooks() ([]Book, error) {
	return s.repo.GetAllBooks()
}

func (s *bookService) DeleteBook(id string) error {
	return s.repo.DeleteBook(id)
}

func NewBookService(r BookRepository) BookService {
	return &bookService{repo: r}
}
