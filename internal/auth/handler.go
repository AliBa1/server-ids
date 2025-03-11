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
	fmt.Fprint(w, "Welcome to the auth service\n")
}

func (h *AuthHandler) PostLogin(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Checking credentials...\n")
	username := "u"
	password := "p"
	h.service.Login(username, password)
}

func (h *AuthHandler) PostRegister(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Checking credentials...\n")
	username := "u"
	password := "p"
	h.service.Register(username, password)
}