package user

import (
	"fmt"
	"net/http"
)

// handle HTTP requests can call services

type UserHandler struct {
	service *UserService
}

func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) UpdateRole(w http.ResponseWriter, req *http.Request) {
	// add middleware to authorize before running
	// track user who changed role?
	username := req.FormValue("username")
	newRole := req.FormValue("newRole")
	if username == "" || newRole == "" {
		fmt.Fprintln(w, "Missing username or new role")
		return
	}
	err := h.service.UpdateRole(username, newRole)
	if err != nil {
		fmt.Fprintf(w, "Error occured while updating %s's role: %s\n", username, err)
	}
	fmt.Fprintf(w, "%s now has the %s role\n", username, newRole)
}
