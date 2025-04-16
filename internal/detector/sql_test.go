package detector

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSQLDetection(t *testing.T) {
	tests := []struct {
		name        string
		request     *http.Request
		wasDetected bool
		message     string
	}{
		{
			name:        "SQL injection in URL",
			request:     httptest.NewRequest("GET", "/id=1%27%20OR%20%271%27=%271", nil), // url = "/id=1' OR '1'='1"
			wasDetected: true,
			message:     "detected in url path: /id=1' OR '1'='1",
		},
		// {
		// 	name: "SQL injection in headers",
		// 	request: func() *http.Request {
		// 		req := httptest.NewRequest("GET", "/", nil)
		// 		req.Header.Set("X-Test-Header", "' UNION SELECT null, username, password FROM users --")
		// 		return req
		// 	}(),
		// 	wasDetected: true,
		// 	message:     "detected in HTTP header X-Test-Header: ' UNION SELECT null, username, password FROM users --",
		// },
		{
			name: "SQL injection in cookies",
			request: func() *http.Request {
				r := httptest.NewRequest("GET", "/", nil)
				r.AddCookie(&http.Cookie{Name: "session", Value: "' OR IF(1=1, SLEEP(5), 0) --"})
				return r
			}(),
			wasDetected: true,
			message:     "detected in cookie: ' OR IF(1=1, SLEEP(5), 0) --",
		},
		{
			name: "SQL injection in body",
			request: func() *http.Request {
				formData := url.Values{}
				formData.Set("username", "user' #")
				formData.Set("password", "password")

				r := httptest.NewRequest("POST", "/login", strings.NewReader(formData.Encode()))
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				return r
			}(),
			wasDetected: true,
			message:     "detected in body: user' #",
		},
		{
			name: "False positive single quote in body",
			request: func() *http.Request {
				formData := url.Values{}
				formData.Set("username", "D'Angelo")
				formData.Set("password", "password")

				r := httptest.NewRequest("POST", "/login", strings.NewReader(formData.Encode()))
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				return r
			}(),
			wasDetected: true,
			message:     "detected in body: D'Angelo",
		},
		{
			name: "False positive dashes in body",
			request: func() *http.Request {
				formData := url.Values{}
				formData.Set("username", "spider--man")
				formData.Set("password", "password")

				r := httptest.NewRequest("POST", "/login", strings.NewReader(formData.Encode()))
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				return r
			}(),
			wasDetected: true,
			message:     "detected in body: spider--man",
		},
		{
			name:        "No SQL injection",
			request:     httptest.NewRequest("GET", "/docs", nil),
			wasDetected: false,
			message:     "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			detector := NewDetector()
			sqlDetection := &SQLDetection{}

			found, err := sqlDetection.Run(httptest.NewRecorder(), test.request, detector)

			assert.NoError(t, err)
			assert.Equal(t, test.wasDetected, found)
			if test.wasDetected && len(detector.Alerts) > 0 {
				assert.Equal(t, test.message, detector.Alerts[0].Message)
			}
		})
	}
}
