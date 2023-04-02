package repository

import (
	"backend/model"
)

type BookRepo interface {
	SaveBook(book model.Book, category []int) (model.Book, error)
	UpdateBook(book model.Book, category []int) (model.Book, error)
	DeleteBook(bookId int) error
	SelectAllBook() ([]model.Book, error)
	SelectBookById(bookId int) (model.Book, error)
}
