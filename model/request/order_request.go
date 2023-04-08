package request

type OrderSave struct {
	CartId uint ` validate:"required"`
}

type OrderStatusChange struct {
	Status int ` validate:"required"`
}
