package document

import (
	"log"
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
	docs, err = h.service.GetAllDocs()

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
	log.Printf("Vars: %+v\n", vars)
	title := vars["title"]
	log.Printf("Title: %s\n", title)

	// fmt.Fprintf(w, "Getting document '%s' from the database...\n", title)
	// fmt.Fprintln(w, "")

	var doc models.Document
	var err error
	doc, err = h.service.GetDoc(title)
	if err != nil {
		log.Println("Error 1")
		log.Printf("%s\n", title)
		// log.Printf("%s\n", doc.Title)
		log.Printf("%s\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.service.DisplayDoc(doc, w)
	if err != nil {
		log.Println("Error 2")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
