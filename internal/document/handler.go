package document

import (
	"fmt"
	"net/http"
	"server-ids/internal/models"
	"strings"
)

type DocsHandler struct {
	service *DocsService
}

func NewDocsHandler(service *DocsService) *DocsHandler {
	return &DocsHandler{service: service}
}

func (h *DocsHandler) GetDocs(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Getting documents from the database...")
	fmt.Fprintln(w, "")

	var docs []models.Document
	var err error
	docs, err = h.service.GetAllDocs()

	if err != nil {
		fmt.Fprintf(w, "Error occured while getting documents: %s\n", err)
	}

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
}

// func (h *DocsHandler) GetDoc(w http.ResponseWriter, req *http.Request) {

// }