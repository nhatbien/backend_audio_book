package request

type OrderSave struct {
	CartId uint ` validate:"required"`
}
