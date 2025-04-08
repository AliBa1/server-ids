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

func (a *Alert) SendEmail() {

}

func (a *Alert) SendText() {

}

func (a *Alert) LogToConsole() {
	// log.Printf("jj")
}

func (a *Alert) LogToFile() {
	
}

func (a *Alert) WriteToUser() {

}
