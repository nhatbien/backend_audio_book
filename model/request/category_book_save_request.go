package request

type CategoryBookSave struct {
	Name        string `json:"name,omitempty"  validate:"required"`
	Description string `json:"description,omitempty" validate:"required"`
	Images      string `json:"images,omitempty" validate:"required"`
}
