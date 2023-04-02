package repository

import (
	"backend/model"
)

type CategoryBookRepo interface {
	SaveCategory(category model.BookCategory) (model.BookCategory, error)
	UpdateCategory(category model.BookCategory) (model.BookCategory, error)
	//DeleteCategory(categoryId int) error
	SelectAllCategory() ([]model.BookCategory, error)
	SelectCategoryById(categoryId int) (model.BookCategory, error)
}
