package detector

import (
	"fmt"
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
	fmt.Println("Possible attack detected!")
	fmt.Printf("Severity: %s\n", a.Severity)
	fmt.Printf("Type: %s\n", a.AttackType)
	fmt.Printf("Date/Time: %s\n", a.Time)
	fmt.Printf("Message: %s\n", a.Message)
	// fmt.Printf("IP: %s\n", a.SourceIP)
	fmt.Println("")
}

func (a *Alert) LogToFile() {
	
}

func (a *Alert) WriteToUser() {

}
