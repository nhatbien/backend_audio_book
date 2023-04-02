package request

type BookSaveRequest struct {
	BookName     string `json:"book_name" validate:"required"`
	Author       string `json:"author" validate:"required"`
	Content      string `json:"content" validate:"required"`
	Img          string `json:"img" validate:"required"`
	Audio        string `json:"audio" `
	Price        int    `json:"price"`
	IsHot        bool   `json:"is_hot" db:"is_hot, omitempty"`
	IsNew        bool   `json:"is_new" db:"is_new, omitempty"`
	IsBestSeller bool   `json:"is_best_seller" db:"is_best_seller, omitempty"`
	IsSale       bool   `json:"is_sale" db:"is_sale, omitempty"`
	IsFree       bool   `json:"is_free" db:"is_free, omitempty"`
	Status       int    `json:"status"`
	BookCategory []int  `json:"book_category"`
}

type BookUpdateRequest struct {
	BookName     string `json:"book_name" `
	Author       string `json:"author" `
	Content      string `json:"content" `
	Img          string `json:"img" `
	Audio        string `json:"audio" `
	Price        int    `json:"price"`
	IsHot        bool   `json:"is_hot" db:"is_hot, omitempty"`
	IsNew        bool   `json:"is_new" db:"is_new, omitempty"`
	IsBestSeller bool   `json:"is_best_seller" db:"is_best_seller, omitempty"`
	IsSale       bool   `json:"is_sale" db:"is_sale, omitempty"`
	IsFree       bool   `json:"is_free" db:"is_free, omitempty"`
	Status       int    `json:"status"`
	BookCategory []int  `json:"book_category"`
}
