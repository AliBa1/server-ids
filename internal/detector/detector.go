package detector

import (
	"net"
	"net/http"
	"time"
)

type DetectionService interface {
	Run(w http.ResponseWriter, r *http.Request, detector *Detector)
}

type Detector struct {
	Services []DetectionService
	Alerts   []Alert
}

func NewDetector() *Detector {
	return &Detector{}
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
	}
}

func (d *Detector) AddAlert(sID int, severity string, attackType string, message string, sourceIP net.IP) {
	d.Alerts = append(d.Alerts, Alert{
		SignatureID: sID,
		Severity:   severity,
		AttackType: attackType,
		Time:       time.Now(),
		Message:    message,
		SourceIP:   sourceIP,
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
