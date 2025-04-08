package detector

import "net/http"

type DDoSDetection struct {
	// rules []string
}


func (s *DDoSDetection) Run(w http.ResponseWriter, r *http.Request, d *Detector) {
	// if detected
	// detector.AddAlert("high", "SQL Injection", "", "")
}