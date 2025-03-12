package models

import (
	"math/rand"
	"net"
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
	Username            string                     `json:"username"`
	Password            string                     `json:"password"`
	Role                string                     `json:"role"`
	LastLoginDate       time.Time                  `json:"last_login_date"`
	LastLoginIP         net.IP                     `json:"last_login_ip"`
	FailedLoginAttempts map[string]FailedLoginInfo `json:"failed_login_attempts"`
}

func NewUser(username string, password string, role string) *User {
	return &User{
		Username:            username,
		Password:            password,
		Role:                role,
		LastLoginDate:       time.Now(),
		LastLoginIP:         getRandomIP(),
		FailedLoginAttempts: make(map[string]FailedLoginInfo),
	}
}

func getRandomIP() net.IP {
	// from random IP generator
	ipAddrs := []string{
		"202.28.138.47",
		"95.201.59.42",
		"220.249.20.246",
		"103.83.16.19",
		"128.42.12.72",
		"142.14.120.253",
		"244.141.202.125",
		"240.147.243.177",
		"20.82.38.65",
		"5.44.249.82",
		"164.202.109.115",
		"100.60.191.6",
		"121.112.7.81",
		"40.97.181.139",
		"232.6.136.166",
		"30.222.58.247",
		"253.94.106.94",
		"223.244.42.121",
		"120.41.162.92",
		"174.212.113.107",
		"32.208.61.231",
		"148.50.165.204",
		"168.15.125.138",
		"131.148.212.175",
		"107.61.171.114",
	}
	return net.ParseIP(ipAddrs[rand.Intn(len(ipAddrs))])
}
