package document

import (
	"server-ids/internal/middleware"
	"server-ids/internal/sessions"

	"github.com/gorilla/mux"
)

func RegisterDocumentRoutes(r *mux.Router, m *middleware.Middleware, service *DocsService, sDB *sessions.SessionsDB) {
	handler := NewDocsHandler(service, sDB)

	r.HandleFunc("/docs", m.ApplyMiddleware(handler.GetDocs)).Methods("GET")
	r.HandleFunc("/docs/{title}", m.ApplyMiddleware(handler.GetDoc)).Methods("GET")
}
