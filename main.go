package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func loggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := time.Now().Local()
		url := r.URL.Path
		protocol := r.Proto
		log.Printf("%s: GET Request on %s using protocol: %s", t, url, protocol)
		next(w, r)
	}
}

var middleware = []func(http.HandlerFunc)http.HandlerFunc{
	// middleware that will run from bottom to top
	loggerMiddleware,
}

func welcomeHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Welcome to the server\n")
}

func main() {
	h := welcomeHandler
	for _, m := range middleware{
		h = m(h)
	}

	r := mux.NewRouter()
	
	// r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	// 	fmt.Fprint(w, "Welcome to the server\n")
	// })
	
	r.HandleFunc("/", h)

	r.HandleFunc("/welcome/{user}", func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		user := vars["user"]
		fmt.Fprintf(w, "Hi, %s! Hope your having a good day!\n", user)
	})

	fmt.Println("Listening on port 80")
	err := http.ListenAndServe(":80", r)
	if err != nil {
		log.Fatalf("Server error: %s\n", err)
	}

	// use templating to log in html? OR just print onto localhost
}