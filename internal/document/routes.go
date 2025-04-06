package document

import (
	"server-ids/internal/auth"
	"server-ids/internal/middleware"

	"github.com/gorilla/mux"
)

func RegisterDocumentRoutes(r *mux.Router, m *middleware.Middleware, service *DocsService, authDB *auth.AuthDBMemory) {
	handler := NewDocsHandler(service, authDB)

	r.HandleFunc("/docs", m.ApplyMiddleware(handler.GetDocs)).Methods("GET")
	r.HandleFunc("/docs/{title}", m.ApplyMiddleware(handler.GetDoc)).Methods("GET")
}
