package request

type CartSave struct {
	Books []uint `json:"books" db:"books, omitempty" validate:"required"`
}
type CartItemSave struct {
	BookId   uint ` validate:"required"`
	Quantity int  ` `
}
