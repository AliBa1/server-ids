package detector

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"regexp"
)

type SQLDetection struct {
	// rules []string
}

func (s *SQLDetection) Run(w http.ResponseWriter, r *http.Request, detector *Detector) (bool, error) {
	// use a switch case most likely
	rules, err := regexp.Compile(`(?i)('|;|--|/\*|\*/|=|\b(and|or|sleep|union|drop|order|select|insert|update|delete|exec|from|where|having)\b)`)

	if err != nil {
		return false, fmt.Errorf("problem compiling regexp in sql detector")
	}

	found := false
	rawIP, _, _ := net.SplitHostPort(r.RemoteAddr)
	ip := net.ParseIP(rawIP)

	decodedURL, err := url.QueryUnescape(r.URL.String())
	if err != nil {
		return false, fmt.Errorf("problem decoding URL: %w", err)
	}

	// check url
	if rules.MatchString(decodedURL) {
		fmt.Println("Checking URL:", decodedURL)
		msg := "detected in url path: " + decodedURL
		detector.AddAlert("warning", "SQL Injection", msg, ip)
		found = true
	}

	// check cookies
	for _, cookie := range r.Cookies() {
		if rules.MatchString(cookie.String()) {
			msg := "detected in cookie: " + cookie.String()
			detector.AddAlert("warning", "SQL Injection", msg, ip)
			found = true
		}
	}

	// check all header values
	for name, values := range r.Header {
		for _, value := range values {
			if rules.MatchString(value) {
				msg := "detected in HTTP header " + name + ": " + value
				detector.AddAlert("warning", "SQL Injection", msg, ip)
				found = true
			}
		}
	}

	// check all body values
	if r.Method == http.MethodPost {
		defer r.Body.Close()
		contents, err := io.ReadAll(r.Body)
		if err != nil {
			return found, err
		}

		if rules.MatchString(string(contents)) {
			msg := "detected in body: " + string(contents)
			detector.AddAlert("warning", "SQL Injection", msg, ip)
			found = true
		}
	}

	return found, nil
}
