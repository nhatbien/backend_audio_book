package repo_impl

import (
	"backend/biedeptrai"
	"backend/db"
	"backend/model"
	"backend/repository"

	"gorm.io/gorm/clause"
)

type BookRepoImpl struct {
	sql *db.Sql
}

func NewBookRepo(sql *db.Sql) repository.BookRepo {
	return &BookRepoImpl{sql: sql}
}

func (n *BookRepoImpl) SaveBook(book model.Book, category []int) (model.Book, error) {
	err := n.sql.Db.Create(&book).Error
	if err != nil {
		return book, err
	}
	return book, nil
}

func (n *BookRepoImpl) UpdateBook(book model.Book, category []int) (model.Book, error) {
	var categories []*model.BookCategory

	if count := n.sql.Db.Where(&model.Book{Id: book.Id}).First(new(model.Book)).RowsAffected; count <= 0 {
		return book, biedeptrai.ErrorBookNotFound
	}
	if count := n.sql.Db.Where("id IN (?)", category).Find(&categories).RowsAffected; count <= 0 {
		return book, biedeptrai.ErrorCategoryNotFound
	}

	if len(category) > 0 {
		err := n.sql.Db.Model(&book).Association("BookCategory").Replace(categories)
		if err != nil {
			return book, err
		}
	}

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
	err := n.sql.Db.Preload(clause.Associations).Find(&books).Error
	if err != nil {
		return books, err
	}
	return books, nil
}

func (n *BookRepoImpl) SelectBookById(bookId int) (model.Book, error) {
	var book model.Book
	err := n.sql.Db.Preload(clause.Associations).Where("id = ?", bookId).First(&book).Error
	if err != nil {
		return book, err
	}
	return book, nil
}
