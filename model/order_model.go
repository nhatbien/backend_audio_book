package model

type Order struct {
	Id         int    `json:"id,omitempty" db:"id, omitempty"`
	UserId     string `json:"user_id,omitempty" db:"user_id, omitempty"`
	BookId     int    `json:"book_id,omitempty" db:"book_id, omitempty"`
	Quantity   int    `json:"quantity,omitempty" db:"quantity, omitempty"`
	TotalPrice int    `json:"total_price,omitempty" db:"total_price, omitempty"`
	CreatedAt  string `json:"created_at,omitempty" db:"created_at, omitempty"`
	UpdatedAt  string `json:"updated_at,omitempty" db:"updated_at, omitempty"`
}
