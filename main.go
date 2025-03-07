package main

import (
	"fmt"
	"log"
	"net/http"
	"server-ids/middleware"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	middleware := middleware.NewMiddleware()
	
	r.HandleFunc("/", middleware.ApplyMiddleware(func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "Welcome to the server\n")
	}))
	
	r.HandleFunc("/welcome/{user}", middleware.ApplyMiddleware(func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		user := vars["user"]
		fmt.Fprintf(w, "Hi, %s! Hope your having a good day!\n", user)
	}))

	fmt.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("Server error: %s\n", err)
	}
}