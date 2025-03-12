package auth

import (
	"fmt"
	"net/http"
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
	username := "u"
	password := "p"
	h.service.Login(username, password)
}

func (h *AuthHandler) PostRegister(w http.ResponseWriter, req *http.Request) {
	username := req.FormValue("username")
	password := req.FormValue("password")
	if username == "" || password == "" {
		fmt.Fprintln(w, "Missing username or password")
		return
	}
	fmt.Fprint(w, "Creating your account...\n")
	h.service.Register(username, password)
}

func (h *AuthHandler) GetUsers(w http.ResponseWriter, req *http.Request) {
	users, err := h.service.db.GetAllUsers()
	if err != nil {
		fmt.Printf("Error occured while getting users: %s\n", err)
		return
	}

	fmt.Fprintln(w, "Getting users from the database...")

	if len(users) < 1 {
		fmt.Fprintln(w, "There are no users")
		return
	}

	fmt.Fprintf(w, "%-15s %-10s %-55s %-20s\n", "Username", "Role", "Last Login Date", "Last Login IP")
	fmt.Fprintf(w, "%-100s\n", "---------------------------------------------------------------------------------------------------")
	for _, u := range users {
		fmt.Fprintf(w, "%-15s %-10s %-55s %-20s\n", u.Username, u.Role, u.LastLoginDate, u.LastLoginIP.String())
	}
}
