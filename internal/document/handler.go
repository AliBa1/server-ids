package document

import (
	"net/http"
	"server-ids/internal/models"
	"server-ids/internal/sessions"
	"server-ids/internal/template"

	"github.com/gorilla/mux"
)

type DocsHandler struct {
	service    *DocsService
	sessionsDB *sessions.SessionsDB
	tmpl       *template.Templates
}

func NewDocsHandler(service *DocsService, sDB *sessions.SessionsDB, template *template.Templates) *DocsHandler {
	return &DocsHandler{service: service, sessionsDB: sDB, tmpl: template}
}

func (h *DocsHandler) GetDocs(w http.ResponseWriter, r *http.Request) {
	var docs []models.Document
	var err error
	docs, err = h.service.GetDocs()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := template.ReturnData{
		Documents: docs,
	}
	h.tmpl.Render(w, "documents", data)
}

func (h *DocsHandler) GetDoc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	data := template.ReturnData{}

	var doc models.Document
	var err error
	doc, err = h.service.GetDoc(title)
	if err != nil {
		data.Error = "Problem occured retreving the document: " + err.Error()
		h.tmpl.Render(w, "document", data)
		return
	}

	if !h.sessionsDB.IsUserLoggedIn(r) {
		data.Error = "Login to access documents"
		h.tmpl.Render(w, "document", data)
		return
	}

	if doc.IsLocked && !h.sessionsDB.IsUserEmployee(r) {
		data.Error = "You don't have access to locked documents"
		h.tmpl.Render(w, "document", data)
		return
	}

	if r.Header.Get("HX-Request") == "true" {
		data.Document = doc
		h.tmpl.Render(w, "document", data)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
