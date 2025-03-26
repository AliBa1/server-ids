package detector

import "net/http"

type XSSDetection struct {
	// rules []string
}


func (s *XSSDetection) Run(w http.ResponseWriter, r *http.Request, detector *Detector) {
	// if detected
	// detector.AddAlert("high", "SQL Injection", "", "")
}