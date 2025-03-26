package detector

import (
	"net"
	"time"
)

type Alert struct {
	Severity   string
	AttackType string
	Time       time.Time
	Message    string
	SourceIP   net.IP
}

func (alert *Alert) SendEmail() {

}

func (alert *Alert) SendText() {

}

func (alert *Alert) LogToConsole() {
	// log.Printf("jj")
}

func (alert *Alert) WriteToUser() {

}
