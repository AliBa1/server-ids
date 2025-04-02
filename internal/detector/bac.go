package detector

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Broken Access Control detection
type BACDetection struct {
}

func (s *BACDetection) Run(w http.ResponseWriter, r *http.Request, detector *Detector) (bool, error) {
	found := false
	// rawIP, _, _ := net.SplitHostPort(r.RemoteAddr)
	// ip := net.ParseIP(rawIP)

	decodedURL, err := url.QueryUnescape(r.URL.String())
	if err != nil {
		return false, fmt.Errorf("problem decoding URL: %w", err)
	}

	// check for role elevation
	if strings.Contains(decodedURL, "/role") {
		// if the user trying to change the role is not an admin or logged in

		// if not an admin
		// msg := username + "tried to change " + urlUsername + "'s role from a " + urlUserRole + " to a " + urlRole

		// if not logged in
		// msg := "unauthenticated person tried to change " + urlUsername + "'s role from a " + urlUserRole + " to a " + urlRole

		// detector.AddAlert("high", "Broken Access Control", msg, ip)
		found = true
	}

	// check for accessing protected document without admin role

	// check other protected routes OR run in authorization middleware function

	// check for accessing API with missing access controls for POST, PUT and DELETE

	return found, nil
}

// Possibly use in authorization middleware function to check if a non user is accessing anything they shouldn't
