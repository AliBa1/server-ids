package user

import "time"

type Role int

const (
	GUEST Role = iota
	EMPLOYEE
	ADMIN
)

type FailedLoginInfo struct {
	Attempts int      
	First    time.Time
	Last     time.Time
	Rate     float64
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role Role `json:"role"`
	// Role string `json:"role"`
	LastLoginDate time.Time `json:"last_login_date"`
	LastLoginIP string `json:"last_login_ip"`
	FailedLoginAttempts map[string]FailedLoginInfo `json:"failed_login_attempts"`
}

func NewUser(username string, password string, role Role) *User {
	return &User{
		Username:            username,
		Password:            password,
		Role:                role,
		LastLoginDate:       time.Now(),
		LastLoginIP:         "",
		FailedLoginAttempts: make(map[string]FailedLoginInfo),
	}
}