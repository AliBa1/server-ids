package user_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"server-ids/internal/database"
	"server-ids/internal/sessions"
	"server-ids/internal/template"
	"server-ids/internal/user"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

// integration test: HTTP, service, and db interaction
func TestUpdateRoleHandler(t *testing.T) {
	rr := httptest.NewRecorder()
	formData := url.Values{}
	formData.Set("newRole", "admin")

	r, err := http.NewRequest(http.MethodPatch, "/users/patrick/role", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Error(err)
	}

	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	db := database.CreateMockDB()
	defer db.Close()
	sessions := sessions.NewSessions(db)
	ur := user.NewUserRepository(db)
	service := user.NewUserService(ur)
	tmpl := template.NewTestTemplate()
	handler := user.NewUserHandler(service, tmpl, sessions)

	// use this if there are route variables
	router := mux.NewRouter()
	router.HandleFunc("/users/{username}/role", handler.UpdateRole).Methods("PATCH")
	router.ServeHTTP(rr, r)

	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	assert.NoError(t, err)
	assert.Equal(t, "users", tmpl.LastRenderedBlock)
	assert.NotNil(t, tmpl.LastRenderedData)
}

func TestUpdateRoleHandler_NoRole(t *testing.T) {
	rr := httptest.NewRecorder()
	formData := url.Values{}
	// formData.Set("newRole", "admin")

	r, err := http.NewRequest(http.MethodPatch, "/users/patrick/role", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Error(err)
	}

	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	db := database.CreateMockDB()
	defer db.Close()
	sessions := sessions.NewSessions(db)
	ur := user.NewUserRepository(db)
	service := user.NewUserService(ur)
	tmpl := template.NewTestTemplate()
	handler := user.NewUserHandler(service, tmpl, sessions)

	// use this if there are route variables
	router := mux.NewRouter()
	router.HandleFunc("/users/{username}/role", handler.UpdateRole).Methods("PATCH")
	router.ServeHTTP(rr, r)

	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	assert.NoError(t, err)
	assert.Equal(t, "users", tmpl.LastRenderedBlock)
}
