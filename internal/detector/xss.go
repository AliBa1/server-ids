package detector

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
)

type XSSDetection struct {
	// rules []string
}

func checkXSS(s string) bool {
	patterns := [2][2]string{
		{"<", ">"},
		{"%3C", "%3E"},
	}

	for _, rules := range patterns {
		if strings.Contains(s, rules[0]) && strings.Contains(s, rules[1]) {
			return true
		}
	}

	return false
}

func (x *XSSDetection) Run(w http.ResponseWriter, r *http.Request, d *Detector) (bool, error) {
	found := false
	rawIP, _, _ := net.SplitHostPort(r.RemoteAddr)
	ip := net.ParseIP(rawIP)

	decodedURL, err := url.QueryUnescape(r.URL.String())
	if err != nil {
		return false, fmt.Errorf("problem decoding URL: %w", err)
	}

	// check url
	if checkXSS(decodedURL) {
		msg := "detected in url path: " + decodedURL
		d.AddAlert(5, "medium", "XSS Attack", msg, ip)
		found = true
	}

	// check cookies
	for _, cookie := range r.Cookies() {
		if checkXSS(cookie.String()) {
			msg := "detected in cookie: " + cookie.String()
			d.AddAlert(6, "medium", "XSS Attack", msg, ip)
			found = true
		}
	}

	// check all header values
	for name, values := range r.Header {
		for _, value := range values {
			if checkXSS(value) {
				msg := "detected in HTTP header " + name + ": " + value
				d.AddAlert(7, "medium", "XSS Attack", msg, ip)
				found = true
			}
		}
	}

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
			if checkXSS(s) {
				msg := "detected in body: " + s
				d.AddAlert(8, "medium", "XSS Attack", msg, ip)
				found = true
			}
		}
	}

	return found, nil
}
