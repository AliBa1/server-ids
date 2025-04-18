package detector

import "net/http"

type LoginAttackDetection struct {
	// honeypot vvvv
	// failedAttempts []string
}

func (s *LoginAttackDetection) Run(w http.ResponseWriter, r *http.Request, d *Detector) {
	// if detected
	// detector.AddAlert("high", "SQL Injection", "", "")
}
