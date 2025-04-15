package user

import (
	"fmt"
	"net/http"
	"server-ids/internal/template"

	"github.com/gorilla/mux"
)

// handle HTTP requests can call services

type UserHandler struct {
	service *UserService
	tmpl    *template.Templates
}

func NewUserHandler(service *UserService, template *template.Templates) *UserHandler {
	return &UserHandler{service: service, tmpl: template}
}

func (h *UserHandler) UpdateRole(w http.ResponseWriter, r *http.Request) {
	// add middleware to authorize before running
	// track user who changed role?
	vars := mux.Vars(r)
	username := vars["username"]
	newRole := r.FormValue("newRole")
	data := template.ReturnData{}

	users, err := h.service.authDB.GetUsers()
	if err != nil {
		data.Error = err.Error()
		h.tmpl.Render(w, "users", data)
		return
	}
	data.Users = users
	if username == "" || newRole == "" {
		fmt.Println("New role: ", newRole)
		// http.Error(w, "Missing username or new role", http.StatusBadRequest)
		data.Error = "Missing username or new role"
		h.tmpl.Render(w, "users", data)
		return
	}

	if !h.service.authDB.SessionsDB.IsUserAdmin(r) {
		data.Error = "You don't have the ability to change roles. Contact your admin for help."
		h.tmpl.Render(w, "users", data)
		return
	}

	err = h.service.UpdateRole(username, newRole)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		data.Error = err.Error()
		h.tmpl.Render(w, "users", data)
		return
	}

	// fmt.Fprintf(w, "%s now has the %s role\n", username, newRole)
	users, err = h.service.authDB.GetUsers()
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		data.Error = err.Error()
		h.tmpl.Render(w, "users", data)
		return
	}

	data.Users = users
	h.tmpl.Render(w, "users", data)
}
