package document

import (
	"server-ids/internal/middleware"
	"server-ids/internal/sessions"
	"server-ids/internal/template"

	"github.com/gorilla/mux"
)

func RegisterDocumentRoutes(r *mux.Router, m *middleware.Middleware, service *DocsService, sessions *sessions.Sessions, template *template.Templates) {
	handler := NewDocsHandler(service, sessions, template)

	r.HandleFunc("/docs", m.ApplyMiddleware(handler.GetDocs)).Methods("GET")
	r.HandleFunc("/docs/{title}", m.ApplyMiddleware(handler.GetDoc)).Methods("GET")
}
