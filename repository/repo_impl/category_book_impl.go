package repo_impl

import (
	"backend/biedeptrai"
	"backend/db"
	"backend/model"
	"backend/repository"
)

type CategoryBookRepoImpl struct {
	sql *db.Sql
}

func NewCategoryBookRepo(sql *db.Sql) repository.CategoryBookRepo {
	return &CategoryBookRepoImpl{sql: sql}
}

func (n *CategoryBookRepoImpl) SaveCategory(category model.BookCategory) (model.BookCategory, error) {

	if count := n.sql.Db.Where(&model.BookCategory{Name: category.Name}).First(new(model.BookCategory)).RowsAffected; count > 0 {
		return category, biedeptrai.ErrorCategoryConflict
	}

	err := n.sql.Db.Create(&category).Error

	if err != nil {
		return category, err
	}
	return category, nil

}

func (n *CategoryBookRepoImpl) UpdateCategory(category model.BookCategory) (model.BookCategory, error) {

	if count := n.sql.Db.Where(&model.BookCategory{Id: category.Id}).First(new(model.BookCategory)).RowsAffected; count <= 0 {
		return category, biedeptrai.ErrorCategoryNotFound
	}

	err := n.sql.Db.Updates(&category).Error

	if err != nil {
		return category, err
	}
	return category, nil

}

func (n *CategoryBookRepoImpl) SelectAllCategory() ([]model.BookCategory, error) {
	var categories []model.BookCategory
	err := n.sql.Db.Find(&categories).Error
	if err != nil {
		return categories, err
	}
	return categories, nil
}
