package models

type Document struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	IsLocked bool   `json:"is_locked"`
}
