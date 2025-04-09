package detector

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"server-ids/internal/sessions"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Broken Access Control detection
type BACDetection struct {
	SessionsDB *sessions.SessionsDB
}

func (s *BACDetection) Run(w http.ResponseWriter, r *http.Request, d *Detector) (bool, error) {
	found := false
	rawIP, _, _ := net.SplitHostPort(r.RemoteAddr)
	ip := net.ParseIP(rawIP)

	decodedURL, err := url.QueryUnescape(r.URL.String())
	if err != nil {
		return false, fmt.Errorf("problem decoding URL: %w", err)
	}

	// check for role elevation
	if strings.Contains(decodedURL, "/role") {
		vars := mux.Vars(r)
		username := vars["username"]
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		r.ParseForm()
		newRole := r.FormValue("newRole")

		// if not logged in
		if !s.SessionsDB.IsUserLoggedIn(r) {
			msg := "unauthenticated person tried to change " + username + "'s role to a " + newRole
			d.AddAlert(9, "high", "BAC Attack", msg, ip)
			found = true
		}

		// if not an admin
		sessionKey, err := r.Cookie("session_key")

		if err == nil {
			keyUUID, err := uuid.Parse(sessionKey.Value)
			if err != nil {
				return found, err
			}

			attemptedUsername, err := s.SessionsDB.GetUsername(keyUUID)
			if err != nil {
				return found, err
			}

			user, err := s.SessionsDB.GetUser(attemptedUsername)
			if err == nil && user.Role != "admin" {
				msg := user.Username + " tried to change " + username + "'s role to a " + newRole
				d.AddAlert(10, "medium", "BAC Attack", msg, ip)
				found = true
			}
		}
	}

	// check for accessing protected document without admin role

	// check other protected routes OR run in authorization middleware function

	// check for accessing API with missing access controls for POST, PUT and DELETE
	return found, nil
}

// Possibly use in authorization middleware function to check if a non user is accessing anything they shouldn't
