package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"server-ids/internal/auth"
	"server-ids/internal/document"
	"server-ids/internal/middleware"
	"server-ids/internal/sessions"
	"server-ids/internal/user"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	sessionsDB := sessions.NewSessionsDB()

	middleware := middleware.NewMiddleware(sessionsDB)

	authDB := auth.NewAuthDBMemory(sessionsDB)
	authService := auth.NewAuthService(authDB)
	auth.RegisterAuthRoutes(r, middleware, authService)

	userService := user.NewUserService(authDB)
	user.RegisterUserRoutes(r, middleware, userService)

	docsDB := document.NewDocsDBMemory()
	documentService := document.NewDocsService(docsDB)
	document.RegisterDocumentRoutes(r, middleware, documentService, sessionsDB)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("internal/views/login.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	})

	fmt.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("Server error: %s\n", err)
	}
}
