package model

import "encoding/json"

type Response struct {
	Status  bool        `json:"Status"`
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
} //@name Response
type MyData struct {
	Value string
}

func (d MyData) MarshalJSON() ([]byte, error) {
	if d.Value == "" {
		return []byte("null"), nil
	}
	return json.Marshal(d.Value)
}
func NewResponse(u *Response) *Response {
	return &Response{
		Status:  u.Status,
		Message: u.Message,
		Data:    u.Data,
	}
}
