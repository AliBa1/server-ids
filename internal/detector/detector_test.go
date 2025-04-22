package detector

import (
	"net"
	"net/http"
	"net/http/httptest"
	"server-ids/internal/template"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddService(t *testing.T) {
	tmpl := template.NewTestTemplate()
	detector := NewDetector(tmpl)
	sql := &SQLDetection{}

	detector.AddService(sql)

	assert.Len(t, detector.Services, 1)
	assert.Equal(t, sql, detector.Services[0])
}

func TestRunSQL(t *testing.T) {
	tmpl := template.NewTestTemplate()
	detector := NewDetector(tmpl)
	sql := &SQLDetection{}
	detector.AddService(sql)

	r := httptest.NewRequest("GET", "/", nil)
	r.AddCookie(&http.Cookie{Name: "session", Value: "' OR IF(1=1, SLEEP(5), 0) --"})
	rr := httptest.NewRecorder()

	detector.Run(rr, r)

	assert.Len(t, detector.Services, 1)
	assert.Greater(t, len(detector.Alerts), 0)
}

func TestRunSQL_NoAttack(t *testing.T) {
	tmpl := template.NewTestTemplate()
	detector := NewDetector(tmpl)
	sql := &SQLDetection{}
	detector.AddService(sql)

	r := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	detector.Run(rr, r)

	assert.Len(t, detector.Services, 1)
	assert.Equal(t, len(detector.Alerts), 0)
}

func TestAddAlert(t *testing.T) {
	tmpl := template.NewTestTemplate()
	detector := NewDetector(tmpl)

	alert := Alert{
		SignatureID: 1,
		Revision:    1,
		Severity:    "high",
		AttackType:  "SQL Injection",
		Message:     "detected in cookies: 1=1",
		SourceIP:    net.ParseIP("000.000.0.0"),
	}

	detector.AddAlert(alert.SignatureID, alert.Revision, alert.Severity, alert.AttackType, alert.Message, alert.SourceIP)

	assert.Len(t, detector.Alerts, 1)
}
