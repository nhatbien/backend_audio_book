package request

import "backend/model"

type BookSaveRequest struct {
	BookName     string                ` validate:"required"`
	Author       string                ` validate:"required"`
	Content      string                ` validate:"required"`
	Img          string                ` validate:"required"`
	Audio        string                ``
	Price        float64               ``
	IsHot        bool                  ``
	IsNew        bool                  ``
	IsBestSeller bool                  ``
	IsSale       bool                  ``
	IsFree       bool                  ``
	Status       int                   ``
	BookCategory []*model.BookCategory ``
}

type BookUpdateRequest struct {
	BookName     string
	Author       string
	Content      string
	Img          string
	Audio        string
	Price        float64
	IsHot        bool
	IsNew        bool
	IsBestSeller bool
	IsSale       bool
	IsFree       bool
	Status       int
	BookCategory []int
}
