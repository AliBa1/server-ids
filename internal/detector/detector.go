package detector

import (
	"net"
	"net/http"
	"server-ids/internal/template"
	"time"
)

type DetectionService interface {
	Run(w http.ResponseWriter, r *http.Request, d *Detector) (bool, error)
}

type Detector struct {
	Services []DetectionService
	Alerts   []Alert
	tmpl     *template.Templates
}

func NewDetector(t *template.Templates) *Detector {
	return &Detector{tmpl: t}
}

func (d *Detector) AddService(service DetectionService) {
	d.Services = append(d.Services, service)
}

func (d *Detector) Run(w http.ResponseWriter, r *http.Request) {
	for _, service := range d.Services {
		service.Run(w, r, d)
	}

	if len(d.Alerts) > 0 {
		d.AlertAdmin()
		d.tmpl.Render(w, "alert", nil)
	}
}

func (d *Detector) AddAlert(sID int, rev int, severity string, attackType string, message string, sourceIP net.IP) {
	d.Alerts = append(d.Alerts, Alert{
		SignatureID: sID,
		Revision:    rev,
		Severity:    severity,
		AttackType:  attackType,
		Time:        time.Now(),
		Message:     message,
		SourceIP:    sourceIP,
	})
}

func (d *Detector) AlertAdmin() {
	for _, alert := range d.Alerts {
		alert.LogToConsole()

		if alert.Severity == "medium" {
			alert.SendEmail()
		}

		if alert.Severity == "high" {
			alert.SendEmail()
			alert.SendText()
		}
	}
}
