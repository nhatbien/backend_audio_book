package model

type Book struct {
	Id           int            `json:"id,omitempty" db:"id, omitempty"`
	Title        string         `json:"title,omitempty" db:"title, omitempty"`
	Author       string         `json:"author,omitempty" db:"author, omitempty"`
	Content      string         `json:"content,omitempty" db:"content, omitempty"`
	Img          string         `json:"img,omitempty" db:"img, omitempty"`
	Audio        string         `json:"audio,omitempty" db:"audio, omitempty"`
	Price        int            `json:"price,omitempty" db:"price, omitempty"`
	Discount     int            `json:"discount,omitempty" db:"discount, omitempty"`
	IsHot        bool           `json:"is_hot,omitempty" db:"is_hot, omitempty"`
	IsNew        bool           `json:"is_new,omitempty" db:"is_new, omitempty"`
	IsBestSeller bool           `json:"is_best_seller,omitempty" db:"is_best_seller, omitempty"`
	IsSale       bool           `json:"is_sale,omitempty" db:"is_sale, omitempty"`
	IsFree       bool           `json:"is_free,omitempty" db:"is_free, omitempty"`
	Status       int            `json:"status,omitempty" gorm:"default:0"`
	CreatedAt    string         `json:"created_at,omitempty" db:"created_at, omitempty"`
	UpdatedAt    string         `json:"updated_at,omitempty" db:"updated_at, omitempty"`
	BookCategory []BookCategory `json:"book_category,omitempty" gorm:"many2many:meta_book_category;"`
}

type AudioBookChapter struct {
	Id        int    `json:"id,omitempty" db:"id, omitempty"`
	Name      string `json:"name,omitempty" db:"name, omitempty"`
	BookId    int    `json:"book_id,omitempty" db:"book_id, omitempty"`
	Audio     string `json:"audio,omitempty" db:"audio, omitempty"`
	IsDeleted bool   `json:"is_deleted,omitempty" db:"is_deleted, omitempty"`
	CreatedAt string `json:"created_at,omitempty" db:"created_at, omitempty"`
	UpdatedAt string `json:"updated_at,omitempty" db:"updated_at, omitempty"`
}
