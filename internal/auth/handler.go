package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

// handle HTTP requests can call services

type AuthHandler struct {
	service *AuthService
}

func NewAuthHandler(service *AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) GetAuth(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Welcome to the auth service")
}

func (h *AuthHandler) PostLogin(w http.ResponseWriter, req *http.Request) {
	username := req.FormValue("username")
	password := req.FormValue("password")
	if username == "" || password == "" {
		http.Error(w, "Missing username or password", http.StatusBadRequest)
		return
	}
	token, err := h.service.Login(username, password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
	fmt.Fprintf(w, "Hello %s! You are now logged in.\n", username)
}

func (h *AuthHandler) PostRegister(w http.ResponseWriter, req *http.Request) {
	username := req.FormValue("username")
	password := req.FormValue("password")
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

func (h *AuthHandler) GetUsers(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Getting users from the database...")
	fmt.Fprintln(w, "")
	users, err := h.service.GetAllUsers()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(users) < 1 {
		fmt.Fprintln(w, "There are no users")
	} else {
		fmt.Fprintf(w, "%-15s %-10s %-55s %-20s\n", "Username", "Role", "Last Login Date", "Last Login IP")
		fmt.Fprintf(w, "%s\n", strings.Repeat("-", 100))
		for _, u := range users {
			fmt.Fprintf(w, "%-15s %-10s %-55s %-20s\n", u.Username, u.Role, u.LastLoginDate, u.LastLoginIP.String())
		}
	}
}
