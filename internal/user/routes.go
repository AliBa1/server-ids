package user

import (
	"server-ids/internal/middleware"
	"server-ids/internal/template"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(r *mux.Router, m *middleware.Middleware, service *UserService, template *template.Templates) {
	handler := NewUserHandler(service, template)

	// r.HandleFunc("/user", m.ApplyMiddleware(handler.GetUser)).Methods("GET")
	r.HandleFunc("/users/{username}/role", m.Authorization(m.ApplyMiddleware(handler.UpdateRole))).Methods("PATCH")
}
