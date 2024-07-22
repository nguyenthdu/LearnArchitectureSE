package models

type Book struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	PageCount int     `json:"pageCount"`
	Author    *Author `json:"author"`
}
