package models

type Document struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	IsLocked bool   `json:"is_locked"`
}

func NewDocument(title string, content string, isLocked bool) *Document {
	return &Document{
		Title: title,
		Content: content,
		IsLocked: isLocked,
		// date created?
	}
}