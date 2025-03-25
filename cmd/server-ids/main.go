package main

import (
	"fmt"
	"log"
	"net/http"
	"server-ids/internal/auth"
	"server-ids/internal/document"
	"server-ids/internal/middleware"
	"server-ids/internal/user"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	middleware := middleware.NewMiddleware()

	authDB := auth.NewAuthDBMemory()
	authService := auth.NewAuthService(authDB)
	auth.RegisterAuthRoutes(r, middleware, authService)

	userService := user.NewUserService(*authDB)
	user.RegisterUserRoutes(r, middleware, userService)

	docsDB := document.NewDocsDBMemory()
	documentService := document.NewDocsService(docsDB)
	document.RegisterDocumentRoutes(r, middleware, documentService)

	// delete
	r.HandleFunc("/", middleware.ApplyMiddleware(func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "Welcome to the server\n")
	}))

	// delete
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
