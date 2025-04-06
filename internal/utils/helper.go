package utils

import (
	"net/http"

	"github.com/google/uuid"
)

func IsUserLoggedIn(r *http.Request, sessions map[uuid.UUID]string) bool {
	keyCookie, err := r.Cookie("session_key")
	if err != nil {
		return false
	}

	key, err := uuid.Parse(keyCookie.Value)
	if err != nil {
		return false
	}

	if sessions[key] == "" {
		return false
	}

	return true
}
