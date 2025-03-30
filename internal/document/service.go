package document

import (
	"server-ids/internal/models"
)

// handles buisness logic and calls database

type DocsService struct {
	db *DocsDBMemory
}

func NewDocsService(db *DocsDBMemory) *DocsService {
	return &DocsService{db: db}
}

func (d *DocsService) GetDocs() ([]models.Document, error) {
	docs, err := d.db.GetAllDocs()
	return docs, err
}

func (d *DocsService) GetDoc(title string) (models.Document, error) {
	doc, err := d.db.GetDoc(title)
	return doc, err
}
