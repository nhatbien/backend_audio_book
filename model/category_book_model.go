package model

type BookCategory struct {
	Id          int    `json:"id," gorm:"not null;primaryKey"`
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description" `
	Images      string `json:"images" `
	Book        []Book `json:"book" gorm:"many2many:user_emails;"`
}
