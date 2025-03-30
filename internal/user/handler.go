package user

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// handle HTTP requests can call services

type UserHandler struct {
	service *UserService
}

func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) UpdateRole(w http.ResponseWriter, r *http.Request) {
	// add middleware to authorize before running
	// track user who changed role?
	vars := mux.Vars(r)
	username := vars["username"]
	newRole := r.FormValue("newRole")

	if username == "" || newRole == "" {
		http.Error(w, "Missing username or new role", http.StatusBadRequest)
		return
	}

	err := h.service.UpdateRole(username, newRole)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s now has the %s role\n", username, newRole)
}
