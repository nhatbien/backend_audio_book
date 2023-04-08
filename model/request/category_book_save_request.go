package request

type CategoryBookSave struct {
	Name        string `  validate:"required"`
	Description string ` validate:"required"`
	Images      string ` validate:"required"`
}
type CategoryBookUpdate struct {
	Name        string `  `
	Description string ` `
	Images      string ``
}
