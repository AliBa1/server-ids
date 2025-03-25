package document

import (
	"fmt"
	"net/http"
	"server-ids/internal/models"
	"strings"
)

// handles buisness logic and calls database

type DocsService struct {
	db *DocsDBMemory
}

func NewDocsService(db *DocsDBMemory) *DocsService {
	return &DocsService{db: db}
}

func (d *DocsService) GetAllDocs() ([]models.Document, error) {
	docs, err := d.db.GetAllDocs()
	return docs, err
}

func (d *DocsService) DisplayDocs(docs []models.Document, w http.ResponseWriter) error {
	if len(docs) < 1 {
		fmt.Fprintln(w, "There are no documents")
	} else {
		lockEmoji := "🔒"
		unlockEmoji := "🔓"
		fmt.Fprintf(w, "%-30s %-10s\n", "Title", "Locked")
		fmt.Fprintf(w, "%s\n", strings.Repeat("-", 40))
		for _, d := range docs {
			lockStatus := unlockEmoji
			if d.IsLocked {
				lockStatus = lockEmoji
			}
			fmt.Fprintf(w, "%-30s %-10s\n", d.Title, lockStatus)
		}
		fmt.Fprintf(w, "\nTo view any of these documents go to '/docs/{title}'\n")
	}
	return nil
}

func (d *DocsService) GetDoc(title string) (models.Document, error) {
	doc, err := d.db.GetDoc(title)
	return doc, err
}

func (d *DocsService) DisplayDoc(doc models.Document, w http.ResponseWriter) error {
	fmt.Fprintf(w, "Title: %s\n", doc.Title)
	fmt.Fprintf(w, "%s\n", doc.Content)
	return nil
}
