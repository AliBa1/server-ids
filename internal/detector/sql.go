package detector

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"regexp"
)

type SQLDetection struct {
	// rules []string
}

func (s *SQLDetection) Run(w http.ResponseWriter, r *http.Request, d *Detector) (bool, error) {
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
		msg := "detected in url path: " + decodedURL
		d.AddAlert(2, 1, "medium", "SQL Injection", msg, ip)
		found = true
	}

	// check cookies
	for _, cookie := range r.Cookies() {
		if rules.MatchString(cookie.Value) {
			msg := "detected in cookie: " + cookie.Value
			d.AddAlert(2, 2, "medium", "SQL Injection", msg, ip)
			found = true
		}
	}

	// check all header values
	// for name, values := range r.Header {
	// 	for _, value := range values {
	// 		if rules.MatchString(value) && name != "Content-Type" && name != "Accept" && name != "Cookie" {
	// 			msg := "detected in HTTP header " + name + ": " + value
	// 			d.AddAlert(3, "medium", "SQL Injection", msg, ip)
	// 			found = true
	// 		}
	// 	}
	// }

	// check all body values
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			return found, fmt.Errorf("something went wrong parsing form data")
		}

		contents := []string{}
		for _, vals := range r.Form {
			contents = append(contents, vals...)
		}

		for _, s := range contents {
			if rules.MatchString(string(s)) {
				msg := "detected in body: " + string(s)
				d.AddAlert(2, 3, "medium", "SQL Injection", msg, ip)
				found = true
			}
		}
	}

	return found, nil
}
