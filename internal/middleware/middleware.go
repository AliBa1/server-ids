package middleware

import (
	"fmt"
	"log"
	"net/http"
)

type Middleware struct {
	chain []func(http.HandlerFunc) http.HandlerFunc
}

type MiddlewareInterface interface {
	addToChain()
	ApplyMiddleware() http.HandlerFunc
	Logger() http.HandlerFunc
	Authorization() http.HandlerFunc
}

func NewMiddleware() *Middleware {
	m := &Middleware{}
	m.addToChain(m.Logger)
	m.addToChain(m.Authorization)
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

// incomplete
func (middleware *Middleware) Authorization(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key, err := r.Cookie("session_key")
		if err != nil {
			fmt.Printf("Unauthed: %s\n", err)
			// return
		} else {
			fmt.Printf("Current Auth Header: %s\n", key)
		}

		// check if key valid
		next(w, r)
	}
}
