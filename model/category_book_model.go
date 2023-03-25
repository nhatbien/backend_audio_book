package model

type BookCategory struct {
	Id          int    `json:"id,omitempty" gorm:"not null;primaryKey"`
	Name        string `json:"name,omitempty" gorm:"not null"`
	Description string `json:"description,omitempty" `
	Images      string `json:"images,omitempty" `
	Book        []Book `json:"book,omitempty" gorm:"many2many:user_emails;"`
}
