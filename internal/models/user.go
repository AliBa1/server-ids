package models

import (
	"time"
)

type FailedLoginInfo struct {
	Attempts int
	First    time.Time
	Last     time.Time
	Rate     float64
}

// Role can be 'guest', 'employee', 'admin'
type User struct {
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	Role          string    `json:"role"`
	LastLoginDate time.Time `json:"last_login_date"`
	// LastLoginIP         net.IP                     `json:"last_login_ip"`
	// FailedLoginAttempts map[string]FailedLoginInfo `json:"failed_login_attempts"`
}

func NewUser(username string, password string, role string) *User {
	return &User{
		Username:      username,
		Password:      password,
		Role:          role,
		LastLoginDate: time.Now(),
		// LastLoginIP:         "",
		// FailedLoginAttempts: make(map[string]FailedLoginInfo),
	}
}
