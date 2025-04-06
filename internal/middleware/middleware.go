package middleware

import (
	"log"
	"net/http"
	"server-ids/internal/utils"

	"github.com/google/uuid"
)

type Middleware struct {
	chain []func(http.HandlerFunc) http.HandlerFunc
}

// runs in reverse order
func NewMiddleware() *Middleware {
	m := &Middleware{}
	m.addToChain(m.Logger)
	// m.addToChain(m.Authorization)
	// m.addToChain(m.IDS)
	return m
}

func (middleware *Middleware) addToChain(newMiddleware func(http.HandlerFunc) http.HandlerFunc) {
	middleware.chain = append([]func(http.HandlerFunc) http.HandlerFunc{newMiddleware}, middleware.chain...)
}

func (middleware *Middleware) ApplyMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	for _, m := range middleware.chain {
		handler = m(handler)
	}
	return handler
}

func (middleware *Middleware) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		httpMethod := r.Method
		protocol := r.Proto
		log.Printf("- %s %s Request @ URL: %s", protocol, httpMethod, url)
		next(w, r)
	}
}

func (middleware *Middleware) Authorization(next http.HandlerFunc, sessions map[uuid.UUID]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !utils.IsUserLoggedIn(r, sessions) {
			http.Error(w, "Unauthorized: Login to gain access to this route", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}

// run before authorization
func (middleware *Middleware) IDS(next http.HandlerFunc) http.HandlerFunc {
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
