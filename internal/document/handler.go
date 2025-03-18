package document

import (
	"fmt"
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

func (h *DocsHandler) GetDocs(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Getting documents from the database...")
	fmt.Fprintln(w, "")

	var docs []models.Document
	var err error
	docs, err = h.service.GetAllDocs()

	if err != nil {
		fmt.Fprintf(w, "Error occured while getting documents: %s\n", err)
		return
	}

	err = h.service.DisplayDocs(docs, w)
	if err != nil {
		fmt.Fprintf(w, "Error occured while displaying documents: %s\n", err)
		return
	}
}

func (h *DocsHandler) GetDoc(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	title := vars["title"]

	fmt.Fprintf(w, "Getting document '%s' from the database...\n", title)
	fmt.Fprintln(w, "")

	var doc models.Document
	var err error
	doc, err = h.service.GetDoc(title)

	if err != nil {
		fmt.Fprintf(w, "Error occured while getting document: %s\n", err)
		return
	}

	err = h.service.DisplayDoc(doc, w)
	if err != nil {
		fmt.Fprintf(w, "Error occured while displaying document: %s\n", err)
		return
	}
}
