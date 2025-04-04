package detector

import (
	"net"
	"time"
)

type Alert struct {
	SignatureID int
	Severity    string // "low", "medium", "high"
	AttackType  string
	Time        time.Time
	Message     string
	SourceIP    net.IP
}

func (alert *Alert) SendEmail() {

}

func (alert *Alert) SendText() {

}

func (alert *Alert) LogToConsole() {
	// log.Printf("jj")
}

func (alert *Alert) LogToFile() {
	
}

func (alert *Alert) WriteToUser() {

}
