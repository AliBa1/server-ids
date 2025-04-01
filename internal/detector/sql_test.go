package detector

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// type MockDetector struct {
// 	mock.Mock
// }

// func (m *MockDetector) AddAlert(level, alertType, message string, ip net.IP) {
// 	m.Called(level, alertType, message, ip)
// }

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
			fmt.Println(detector.Alerts)
			assert.Equal(t, test.wasDetected, found)
			if test.wasDetected && len(detector.Alerts) > 0 {
				assert.Equal(t, test.message, detector.Alerts[0].Message)
			}
		})
	}
}
