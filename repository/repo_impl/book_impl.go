package repo_impl

import (
	"backend/db"
	"backend/model"
	"backend/repository"
)

type BookRepoImpl struct {
	sql *db.Sql
}

func NewBookRepo(sql *db.Sql) repository.BookRepo {
	return &BookRepoImpl{sql: sql}
}

func (n *BookRepoImpl) SaveBook(book model.Book) (model.Book, error) {
	err := n.sql.Db.Create(&book).Error
	if err != nil {
		return book, err
	}
	return book, nil
}

func (n *BookRepoImpl) UpdateBook(book model.Book) (model.Book, error) {
	err := n.sql.Db.Updates(&book).Error
	if err != nil {
		return book, err
	}
	return book, nil
}

func (n *BookRepoImpl) DeleteBook(bookId int) error {
	err := n.sql.Db.Where("id = ?", bookId).Delete(&model.Book{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (n *BookRepoImpl) SelectAllBook() ([]model.Book, error) {
	var books []model.Book
	err := n.sql.Db.Find(&books).Error
	if err != nil {
		return books, err
	}
	return books, nil
}
