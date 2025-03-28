package user

import (
	"server-ids/internal/middleware"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(r *mux.Router, m *middleware.Middleware, service *UserService) {
	handler := NewUserHandler(service)

	// r.HandleFunc("/user", m.ApplyMiddleware(handler.GetUser)).Methods("GET")

	// change vvv to "/users/{username}/role" and make it PATCH instead of PUT
	r.HandleFunc("/user/update-role", m.ApplyMiddleware(handler.UpdateRole)).Methods("PUT")
}
