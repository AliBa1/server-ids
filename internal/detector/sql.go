package detector

import "net/http"

type SQLDetection struct {
	// rules []string
}


func (s *SQLDetection) Run(w http.ResponseWriter, r *http.Request, detector *Detector) {
	// use a switch case most likely
	// if detected
	// detector.AddAlert("high", "SQL Injection", "", "")
}