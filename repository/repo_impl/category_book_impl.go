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
