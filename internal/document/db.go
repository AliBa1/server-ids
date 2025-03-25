package document

import (
	"fmt"
	"server-ids/internal/mock"
	"server-ids/internal/models"
)

// CRUD database

type DocsDBMemory struct {
	Documents []models.Document
}

func NewDocsDBMemory() *DocsDBMemory {
	return &DocsDBMemory{
		Documents: mock.GetMockDocuments(),
	}
}

func (db *DocsDBMemory) GetAllDocs() ([]models.Document, error) {
	return db.Documents, nil
}

func (db *DocsDBMemory) GetDoc(title string) (models.Document, error) {
	// consider making documents a map for faster retrival
	for _, d := range db.Documents {
		if d.Title == title {
			return d, nil
		}
	}
	return models.Document{}, fmt.Errorf("document titled '%s' not found", title)
}

// func (db *DocsDBMemory) CreateDoc(document models.Document) {
// 	db.Documents = append(db.Documents, document)
// }

// func (db *DocsDBMemory) UpdateDoc(document models.Document) error {
// 	for i, d := range db.Documents {
// 		if d.Title == document.Title {
// 			db.Documents[i] = document
// 			return nil
// 		}
// 	}
// 	return fmt.Errorf("document titled '%s' not found", document.Title)
// }
