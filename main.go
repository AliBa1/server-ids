package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	
	r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "Welcome to the server\n")
	})

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