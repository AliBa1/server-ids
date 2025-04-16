package detector

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXSSDetection(t *testing.T) {
	tests := []struct {
		name        string
		request     *http.Request
		wasDetected bool
		message     string
	}{
		{
			name:        "XSS attack in URL",
			request:     httptest.NewRequest("GET", "/%3Cscript%3Ealert(%27Hacked!%27)%3C/script%3E", nil), // url = "<script>alert('Hacked!')</script>"
			wasDetected: true,
			message:     "detected in url path: /<script>alert('Hacked!')</script>",
		},
		{
			name: "XSS attack in headers",
			request: func() *http.Request {
				r := httptest.NewRequest("GET", "/", nil)
				r.Header.Set("X-Test-Header", "<body onload=alert('test1')>")
				return r
			}(),
			wasDetected: true,
			message:     "detected in HTTP header X-Test-Header: <body onload=alert('test1')>",
		},
		{
			name: "XSS attack in cookies",
			request: func() *http.Request {
				r := httptest.NewRequest("GET", "/", nil)
				r.AddCookie(&http.Cookie{Name: "session", Value: "<IMG SRC=j&#X41vascript:alert('test2')>"})
				return r
			}(),
			wasDetected: true,
			message:     "detected in cookie: session=\"<IMG SRC=j&#X41vascript:alert('test2')>\"",
		},
		{
			name: "XSS attack in body",
			request: func() *http.Request {
				formData := url.Values{}
				// formData.Set("username", "%3Cimg src='http://bad.url' onerror=alert(document.cookie);%3E")
				formData.Set("username", "<img src='http://bad.url' onerror=alert(document.cookie);>")
				formData.Set("password", "password")

				r := httptest.NewRequest("POST", "/login", strings.NewReader(formData.Encode()))
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				return r
			}(),
			wasDetected: true,
			message:     "detected in body: <img src='http://bad.url' onerror=alert(document.cookie);>",
		},
		{
			name: "False positive single quote in body",
			request: func() *http.Request {
				formData := url.Values{}
				formData.Set("username", "user")
				formData.Set("password", "1>pass<3")

				r := httptest.NewRequest("POST", "/login", strings.NewReader(formData.Encode()))
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				return r
			}(),
			wasDetected: true,
			message:     "detected in body: 1>pass<3",
		},
		{
			name: "Only one angle bracket",
			request: func() *http.Request {
				formData := url.Values{}
				formData.Set("username", "%3c3")
				formData.Set("password", "password")

				r := httptest.NewRequest("POST", "/login", strings.NewReader(formData.Encode()))
				return r
			}(),
			wasDetected: false,
			message:     "detected in body: %3c3",
		},
		{
			name:        "No XSS attack",
			request:     httptest.NewRequest("GET", "/docs", nil),
			wasDetected: false,
			message:     "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			detector := NewDetector()
			xssDetection := &XSSDetection{}

			found, err := xssDetection.Run(httptest.NewRecorder(), test.request, detector)

			assert.NoError(t, err)
			assert.Equal(t, test.wasDetected, found)
			if test.wasDetected && len(detector.Alerts) > 0 {
				assert.Equal(t, test.message, detector.Alerts[0].Message)
			}
		})
	}
}
