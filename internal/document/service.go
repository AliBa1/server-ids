package document

import (
	"fmt"
	"server-ids/internal/models"
)

// handles buisness logic and calls database

type DocsService struct {
	db DocsDB
}

func NewDocsService(db DocsDB) *DocsService {
	return &DocsService{db: db}
}

func (d *DocsService) GetAllDocs() ([]models.Document, error) {
	docs, err := d.db.GetAllDocs()
	if err != nil {
		fmt.Printf("Error occured while getting documents: %s\n", err)
		return nil, err
	}
	return docs, err
}
