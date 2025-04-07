package middleware

import (
	"fmt"
	"log"
	"net/http"
	"server-ids/internal/sessions"
)

type Middleware struct {
	chain      []func(http.HandlerFunc) http.HandlerFunc
	sessionsDB *sessions.SessionsDB
}

// runs in reverse order
func NewMiddleware(sDB *sessions.SessionsDB) *Middleware {
	m := &Middleware{sessionsDB: sDB}
	m.addToChain(m.Logger)
	// m.addToChain(m.Authorization)
	// m.addToChain(m.IDS)
	return m
}

func (m *Middleware) addToChain(newMiddleware func(http.HandlerFunc) http.HandlerFunc) {
	m.chain = append([]func(http.HandlerFunc) http.HandlerFunc{newMiddleware}, m.chain...)
}

func (middleware *Middleware) ApplyMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	for _, m := range middleware.chain {
		handler = m(handler)
	}
	return handler
}

func (m *Middleware) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		httpMethod := r.Method
		protocol := r.Proto
		log.Printf("- %s %s Request @ URL: %s", protocol, httpMethod, url)
		next(w, r)
	}
}

func (m *Middleware) Authorization(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !m.sessionsDB.IsUserLoggedIn(r) {
			fmt.Println("2222222")
			http.Error(w, "Unauthorized: Login to gain access to this route", http.StatusUnauthorized)
			fmt.Println("e333333333")
			return
		}

		next(w, r)
	}
}

// run before authorization
func (m *Middleware) IDS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Composite design pattern

		// detector := &detector.Detector{}

		// if possibility of SQL attack
		// 		detector.AddService(&SQLDetection{})

		// if possibility of Login attack
		// 		detector.AddService(&LoginDetection{})

		// if possibility of XSS attack
		// 		detector.AddService(&XSSDetection{})

		// if possibility of DDoS attack
		// 		detector.AddService(&DDoSDetection{})

		// if possibility of BAL (broken access control) attack
		// 		detector.AddService(&BALDetection{})
	}
}
