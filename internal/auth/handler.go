package auth

import (
	"fmt"
	"net/http"
	"server-ids/internal/template"
	"time"
)

// handle HTTP requests can call services

type AuthHandler struct {
	service *AuthService
	tmpl    *template.Templates
}

func NewAuthHandler(service *AuthService, template *template.Templates) *AuthHandler {
	// return &AuthHandler{service: service}
	return &AuthHandler{service: service, tmpl: template}
}

func (h *AuthHandler) GetLogin(w http.ResponseWriter, r *http.Request) {
	data := template.ReturnData{Error: ""}
	h.tmpl.Render(w, "login", data)
}

func (h *AuthHandler) PostLogin(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	if username == "" || password == "" {
		http.Error(w, "Missing username or password", http.StatusBadRequest)
		return
	}
	token, err := h.service.Login(username, password)

	// tmpl := template.NewTemplate()
	if err != nil {
		data := template.ReturnData{Error: err.Error()}
		h.tmpl.Render(w, "login", data)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_key",
		Value:    token.String(),
		Expires:  time.Now().Add(60 * 24 * time.Hour),
		HttpOnly: true, // protection from man-in-the-middle attacks
		Path:     "/",
		// Secure:   true, // protection from XSS attacks w/ HTTPS: (https://developer.mozilla.org/en-US/docs/Web/HTTP/Guides/Cookies#security)
	})
	// fmt.Fprintf(w, "Hello %s! You are now logged in.\n", username)

	http.Redirect(w, r, "/docs", http.StatusFound)
}

func (h *AuthHandler) PostRegister(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	if username == "" || password == "" {
		http.Error(w, "Missing username or password", http.StatusBadRequest)
		return
	}
	err := h.service.Register(username, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Welcome %s! Your account has been created\n", username)
}

func (h *AuthHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Getting users from the database...")
	// fmt.Fprintln(w, "")
	users, err := h.service.GetAllUsers()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := template.ReturnData{Users: users}
	h.tmpl.Render(w, "users", data)
}
