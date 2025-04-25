package detector

import "net/http"

type LoginAttackDetection struct {
	// honeypot vvvv
	// failedAttempts []string
}

func (s *LoginAttackDetection) Run(w http.ResponseWriter, r *http.Request, d *Detector) {
	// store all failed login attempts (will count as honeypot)

	// if X failed attempts in X seconds/minutes
	// detector.AddAlert("high", "SQL Injection", "", "")
	
	// if X failed attempts in X seconds/minutes
	// detector.AddAlert("high", "SQL Injection", "", "")
	
	// if failed attempt has common credentials admin123, pass123, etc.
	// detector.AddAlert("high", "SQL Injection", "", "")
	
	// if X amount of username tried from same IP (can encrypt ip or store session in cookies?)
	// detector.AddAlert("high", "SQL Injection", "", "")
	
	// if login attempt made from unusual location
	// detector.AddAlert("high", "SQL Injection", "", "")
}
