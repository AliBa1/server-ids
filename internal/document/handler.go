package document

import (
	"net/http"
	"server-ids/internal/models"

	"github.com/gorilla/mux"
)

type DocsHandler struct {
	service *DocsService
}

func NewDocsHandler(service *DocsService) *DocsHandler {
	return &DocsHandler{service: service}
}

func (h *DocsHandler) GetDocs(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Getting documents from the database...")
	// fmt.Fprintln(w, "")

	var docs []models.Document
	var err error
	docs, err = h.service.GetDocs()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.service.DisplayDocs(docs, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *DocsHandler) GetDoc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	// fmt.Fprintf(w, "Getting document '%s' from the database...\n", title)
	// fmt.Fprintln(w, "")

	var doc models.Document
	var err error
	doc, err = h.service.GetDoc(title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.service.DisplayDoc(doc, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
