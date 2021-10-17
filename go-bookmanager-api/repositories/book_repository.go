package repositories

import (
	"gorm.io/gorm"
	"matheusmv.com/go-bookmanager-api/models"
)

type BookRepository interface {
	SelectById(id int) (models.Book, error)
	SelectAll() ([]models.Book, error)
	Insert(book models.Book) (models.Book, error)
	Update(book models.Book) (models.Book, error)
	Delete(book models.Book) (bool, error)
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *bookRepository {
	return &bookRepository{db}
}

func (r *bookRepository) SelectById(id int) (models.Book, error) {
	var book models.Book

	err := r.db.Model(&book).Preload("Authors").First(&book, "id = ?", id).Error

	if err != nil {
		return book, err
	}

	return book, nil
}

func (r *bookRepository) SelectAll() ([]models.Book, error) {
	var books []models.Book

	err := r.db.Model(&books).Preload("Authors").Find(&books).Error

	if err != nil {
		return books, err
	}

	return books, nil
}

func (r *bookRepository) Insert(book models.Book) (models.Book, error) {
	err := r.db.Create(&book).Error

	if err != nil {
		return book, err
	}

	return book, nil
}

func (r *bookRepository) Update(book models.Book) (models.Book, error) {
	err := r.db.Save(&book).Error

	if err != nil {
		return book, err
	}

	return book, nil
}

func (r *bookRepository) Delete(book models.Book) (bool, error) {
	err := r.db.Delete(&book).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
