package repo_impl

import (
	"backend/biedeptrai"
	"backend/db"
	"backend/model"
	"backend/repository"
	"time"

	"gorm.io/gorm/clause"
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

func (n *CategoryBookRepoImpl) UpdateCategory(category model.BookCategory, categoryId uint) (model.BookCategory, error) {

	if count := n.sql.Db.Where(&model.BookCategory{ID: categoryId}).First(new(model.BookCategory)).RowsAffected; count <= 0 {
		return category, biedeptrai.ErrorCategoryNotFound
	}
	if count := n.sql.Db.Where(&model.BookCategory{Name: category.Name}).First(new(model.BookCategory)).RowsAffected; count > 0 {
		return category, biedeptrai.ErrorCategoryConflict
	}
	category.UpdatedAt = time.Now()
	err := n.sql.Db.Where(&model.BookCategory{ID: categoryId}).Updates(category).Error

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

func (n *CategoryBookRepoImpl) SelectCategoryById(categoryId int) (model.BookCategory, error) {
	var category model.BookCategory
	err := n.sql.Db.Where("id = ?", categoryId).Preload(clause.Associations).First(&category).Error
	if err != nil {
		return category, err
	}
	return category, nil
}
