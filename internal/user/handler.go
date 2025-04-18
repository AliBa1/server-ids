package user

import (
	"net/http"
	"server-ids/internal/sessions"
	"server-ids/internal/template"

	"github.com/gorilla/mux"
)

// handle HTTP requests can call services

type UserHandler struct {
	service  *UserService
	tmpl     *template.Templates
	sessions *sessions.Sessions
}

func NewUserHandler(service *UserService, template *template.Templates, sessions *sessions.Sessions) *UserHandler {
	return &UserHandler{service: service, tmpl: template, sessions: sessions}
}

func (h *UserHandler) UpdateRole(w http.ResponseWriter, r *http.Request) {
	// track user who changed role?
	vars := mux.Vars(r)
	username := vars["username"]
	newRole := r.FormValue("newRole")
	users, err := h.service.userRepo.GetUsers()
	data := template.ReturnData{Users: users}
	if err != nil {
		data.Error = "Problem occured on the server. Try again."
		h.tmpl.Render(w, "users", data)
		return
	}

	if username == "" || newRole == "" {
		data.Error = "Missing username or new role"
		h.tmpl.Render(w, "users", data)
		return
	}

	if !h.sessions.IsUserAdmin(r) {
		data.Error = "You don't have the ability to change roles. Contact an admin for help."
		h.tmpl.Render(w, "users", data)
		return
	}

	err = h.service.UpdateRole(username, newRole)
	if err != nil {
		data.Error = err.Error()
		h.tmpl.Render(w, "users", data)
		return
	}

	users, err = h.service.userRepo.GetUsers()
	if err != nil {
		data.Error = "Problem occured on the server. Try again."
		h.tmpl.Render(w, "users", data)
		return
	}

	data.Users = users
	h.tmpl.Render(w, "users", data)
}
