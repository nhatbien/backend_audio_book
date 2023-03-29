package request

type CategoryBookSave struct {
	Name        string `json:"name"  validate:"required"`
	Description string `json:"description" validate:"required"`
	Images      string `json:"images" validate:"required"`
}
