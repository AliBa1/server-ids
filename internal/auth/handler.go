package auth

import (
	"fmt"
	"net/http"
	"strings"
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
	fmt.Fprintln(w, "Checking credentials...")
	username := req.FormValue("username")
	password := req.FormValue("password")
	if username == "" || password == "" {
		fmt.Fprintln(w, "Missing username or password")
		return
	}
	key, err := h.service.Login(username, password)
	// work in progress vvvvv
	fmt.Printf("Current Auth Header: %s\n", w.Header().Get("Authorization"))
	w.Header().Set("Authorization", "Basic "+key.String())
	fmt.Printf("Current Auth Header: %s\n", w.Header().Get("Authorization"))
	
	if err != nil {
		fmt.Fprintf(w, "Error: %s\n", err)
		return
	}
	fmt.Fprintf(w, "Hello %s! You are now logged in.\n", username)
}

func (h *AuthHandler) PostRegister(w http.ResponseWriter, req *http.Request) {
	username := req.FormValue("username")
	password := req.FormValue("password")
	if username == "" || password == "" {
		fmt.Fprintln(w, "Missing username or password")
		return
	}
	err := h.service.Register(username, password)
	if err != nil {
		fmt.Fprintf(w, "Error: %s\n", err)
		return
	}
	fmt.Fprintf(w, "Welcome %s! Your account has been created\n", username)
}

func (h *AuthHandler) GetUsers(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Getting users from the database...")
	fmt.Fprintln(w, "")
	users, err := h.service.GetAllUsers()

	if err != nil {
		fmt.Fprintf(w, "Error occured while getting users: %s\n", err)
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
