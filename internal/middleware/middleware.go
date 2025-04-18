package middleware

import (
	"fmt"
	"log"
	"net/http"
	"server-ids/internal/detector"
	"server-ids/internal/sessions"
)

type Middleware struct {
	chain []func(http.HandlerFunc) http.HandlerFunc
	// sessionsDB *sessions.SessionsDB
	Sessions *sessions.Sessions
}

// runs in reverse order
func NewMiddleware(s *sessions.Sessions) *Middleware {
	// m := &Middleware{sessionsDB: sDB}
	m := &Middleware{Sessions: s}
	m.addToChain(m.Logger)
	// m.addToChain(m.Authorization)
	m.addToChain(m.IDS)
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
		d := detector.NewDetector()
		d.AddService(&detector.BACDetection{
			// SessionsDB: m.sessionsDB,
			Sessions: m.Sessions,
		})
		d.Run(w, r)

		if !m.Sessions.IsUserLoggedIn(r) {
			http.Error(w, "Unauthorized: Login to gain access to this route", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}

// run before authorization
func (m *Middleware) IDS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Composite design pattern
		d := detector.NewDetector()

		err := r.ParseForm()
		if err != nil {
			fmt.Printf("something went wrong parsing form data")
		}

		if len(r.Form) > 0 {
			d.AddService(&detector.SQLDetection{})
			d.AddService(&detector.XSSDetection{})
		}

		d.AddService(&detector.BACDetection{
			Sessions: m.Sessions,
		})

		d.Run(w, r)

		// if possibility of Login attack
		// 		detector.AddService(&LoginDetection{})

		// if possibility of DDoS attack
		// 		detector.AddService(&DDoSDetection{})

		next(w, r)
	}
}
