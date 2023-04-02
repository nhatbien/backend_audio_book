package model

type Order struct {
	Id         int    `json:"id" db:"id, omitempty"`
	UserId     string `json:"user_id" db:"user_id, omitempty"`
	BookId     int    `json:"book_id" db:"book_id, omitempty"`
	Quantity   int    `json:"quantity" db:"quantity, omitempty"`
	TotalPrice int    `json:"total_price" db:"total_price, omitempty"`
	CreatedAt  string `json:"created_at" db:"created_at, omitempty"`
	UpdatedAt  string `json:"updated_at" db:"updated_at, omitempty"`
	Status     int    `json:"status" db:"status, omitempty"`
}
