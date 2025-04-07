package document

import (
	"fmt"
	"net/http"
	"server-ids/internal/models"
	"server-ids/internal/sessions"
	"strings"

	"github.com/gorilla/mux"
)

type DocsHandler struct {
	service    *DocsService
	sessionsDB *sessions.SessionsDB
}

func NewDocsHandler(service *DocsService, sDB *sessions.SessionsDB) *DocsHandler {
	return &DocsHandler{service: service, sessionsDB: sDB}
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

func (h *DocsHandler) GetDoc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	var doc models.Document
	var err error
	doc, err = h.service.GetDoc(title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if doc.IsLocked && !h.sessionsDB.IsUserLoggedIn(r) {
		http.Error(w, "Unauthorized: Login to gain access to this route", http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "Title: %s\n", doc.Title)
	fmt.Fprintf(w, "%s\n", doc.Content)
}
