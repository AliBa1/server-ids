package main

import (
	"fmt"
	"log"
	"net/http"
	"server-ids/internal/auth"
	"server-ids/internal/database"
	"server-ids/internal/document"
	"server-ids/internal/middleware"
	"server-ids/internal/sessions"
	"server-ids/internal/template"
	"server-ids/internal/user"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/web/").Handler(http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))
	r.PathPrefix("/htmx/").Handler(http.StripPrefix("/htmx/", http.FileServer(http.Dir("htmx"))))
	
	db := database.NewDBConnection()
	defer db.Close()

	userRepo := user.NewUserRepository(db)
	authRepo := auth.NewAuthRepository(db)
	docRepo := document.NewDocRepository(db)

	sessions := sessions.NewSessions(db)

	middleware := middleware.NewMiddleware(sessions)

	tmpl := template.NewTemplate()

	authService := auth.NewAuthService(authRepo, userRepo)
	auth.RegisterAuthRoutes(r, middleware, authService, tmpl)

	userService := user.NewUserService(userRepo)
	user.RegisterUserRoutes(r, middleware, userService, tmpl, sessions)

	documentService := document.NewDocsService(docRepo)
	document.RegisterDocumentRoutes(r, middleware, documentService, sessions, tmpl)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/login", http.StatusFound)
	})

	fmt.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("Server error: %s\n", err)
	}
}
