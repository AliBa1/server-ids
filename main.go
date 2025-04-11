package main

import (
	"fmt"
	"log"
	"net/http"
	"server-ids/internal/auth"
	"server-ids/internal/document"
	"server-ids/internal/middleware"
	"server-ids/internal/sessions"
	"server-ids/internal/template"
	"server-ids/internal/user"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	sessionsDB := sessions.NewSessionsDB()

	middleware := middleware.NewMiddleware(sessionsDB)
	
	// add to all other register routes
	tmpl := template.NewTemplate()

	authDB := auth.NewAuthDBMemory(sessionsDB)
	authService := auth.NewAuthService(authDB)
	auth.RegisterAuthRoutes(r, middleware, authService, tmpl)

	userService := user.NewUserService(authDB)
	user.RegisterUserRoutes(r, middleware, userService)

	docsDB := document.NewDocsDBMemory()
	documentService := document.NewDocsService(docsDB)
	document.RegisterDocumentRoutes(r, middleware, documentService, sessionsDB)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := template.ReturnData{Error: ""}
		tmpl.Render(w, "login", data)
	})

	fmt.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("Server error: %s\n", err)
	}
}
