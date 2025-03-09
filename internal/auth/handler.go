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