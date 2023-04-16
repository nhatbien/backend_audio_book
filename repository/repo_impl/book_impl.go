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

func (n *BookRepoImpl) SaveBook(book model.Book) (model.Book, error) {
	if err := n.sql.Db.Create(&book).Error; err != nil {
		return book, err
	}

	if err := n.sql.Db.Preload(clause.Associations).Find(&book).Error; err != nil {
		return book, err
	}
	/* err := n.sql.Db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&book).Error; err != nil {
			return err
		}

		if len(category) > 0 {
			var categories []*model.BookCategory
			if count := n.sql.Db.Where("id IN (?)", category).Find(&categories).RowsAffected; count <= 0 {
				return biedeptrai.ErrorCategoryNotFound
			}

			if err := n.sql.Db.Preload(clause.Associations).Model(&book).Association("BookCategory").Replace(categories); err != nil {
				return err
			}

		}
		if err := n.sql.Db.Updates(&book).Error; err != nil {
			return err
		}
		tx.Commit()

		return nil
	})
	if err != nil {
		return book, err
	}
	*/
	return book, nil
}

func (n *BookRepoImpl) UpdateBook(book model.Book, category []int) (model.Book, error) {
	var categories []*model.BookCategory

	if count := n.sql.Db.First(new(model.Book), book.ID).RowsAffected; count <= 0 {
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
	if count := n.sql.Db.Find(&model.Book{ID: uint(bookId)}).RowsAffected; count <= 0 {
		return biedeptrai.ErrorBookNotFound
	}

	n.sql.Db.Find(&model.Book{ID: uint(bookId)}).Association("BookCategory").Clear()
	n.sql.Db.Where("book_id = ?", bookId).Delete(&model.CartItem{})

	/* err := n.sql.Db.Model(&model.CartItem{}).Where("book_id = ?", bookId).Delete(&model.CartItem{}).Error
	if err != nil {
		return err

	} */

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

func (n *BookRepoImpl) SearchBookByName(bookName string) ([]model.Book, error) {
	var books []model.Book
	err := n.sql.Db.Preload(clause.Associations).Where("book_name LIKE ?", "%"+bookName+"%").Find(&books).Error
	if err != nil {
		return books, err
	}
	return books, nil
}
