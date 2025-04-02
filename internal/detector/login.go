package detector

import "net/http"

type LoginAttackDetection struct {
	// honeypot vvvv
	// failedAttempts []string
}


func (s *LoginAttackDetection) Run(w http.ResponseWriter, r *http.Request, detector *Detector) {
	// if detected
	// detector.AddAlert("high", "SQL Injection", "", "")
}