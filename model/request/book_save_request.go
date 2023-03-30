package request

type BookSaveRequest struct {
	BookName string `json:"book_name" validate:"required"`
	Author   string `json:"author" validate:"required"`
	Content  string `json:"content" validate:"required"`
	Img      string `json:"img" validate:"required"`
	Audio    string `json:"audio" `
	Price    int    `json:"price"`
}
