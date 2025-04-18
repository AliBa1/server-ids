package document

import "server-ids/internal/models"

// handles buisness logic and calls database

type DocsService struct {
	docRepo *DocsRepository
}

func NewDocsService(dr *DocsRepository) *DocsService {
	return &DocsService{docRepo: dr}
}

func (d *DocsService) GetDocs() ([]models.Document, error) {
	docs, err := d.docRepo.GetDocs()
	return docs, err
}

func (d *DocsService) GetDoc(title string) (*models.Document, error) {
	doc, err := d.docRepo.GetDoc(title)
	if err != nil {
		return nil, err
	}
	return doc, err
}
