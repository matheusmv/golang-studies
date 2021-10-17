package services

import (
	"matheusmv.com/go-bookmanager-api/models"
	"matheusmv.com/go-bookmanager-api/repositories"
)

type BookService interface {
	FindABookById(id int) (models.Book, error)
	GetAllBooks() ([]models.Book, error)
	CreateABook(book models.Book) (models.Book, error)
	UpdateABook(id int, bookDetails models.Book) (models.Book, error)
	DeleteABook(id int) (bool, error)
}

type bookService struct {
	repository repositories.BookRepository
}

func NewService(repository repositories.BookRepository) *bookService {
	return &bookService{repository}
}

func (s *bookService) FindABookById(id int) (models.Book, error) {
	book, err := s.repository.SelectById(id)

	if err != nil {
		return book, err
	}

	return book, nil
}

func (s *bookService) GetAllBooks() ([]models.Book, error) {
	books, err := s.repository.SelectAll()

	if err != nil {
		return books, err
	}

	return books, nil
}

func (s *bookService) CreateABook(book models.Book) (models.Book, error) {
	newBook := models.Book{}

	newBook.Isbn = book.Isbn
	newBook.Title = book.Title
	newBook.Authors = book.Authors

	entity, err := s.repository.Insert(newBook)

	if err != nil {
		return entity, err
	}

	return entity, nil
}

func (s *bookService) UpdateABook(id int, bookDetails models.Book) (models.Book, error) {
	book, err := s.repository.SelectById(id)

	if err != nil {
		return book, err
	}

	if bookDetails.Isbn != "" {
		book.Isbn = bookDetails.Isbn
	}

	if bookDetails.Title != "" {
		book.Title = bookDetails.Title
	}

	updatedBook, err := s.repository.Update(book)

	if err != nil {
		return updatedBook, err
	}

	return updatedBook, nil
}

func (s *bookService) DeleteABook(id int) (bool, error) {
	book, err := s.repository.SelectById(id)

	if err != nil || book.ID == 0 {
		return false, err
	}

	status, err := s.repository.Delete(book)

	if err != nil {
		return status, err
	}

	return status, nil
}
