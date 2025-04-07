package user

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"server-ids/internal/auth"
	"server-ids/internal/sessions"
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

	sessionsDB := sessions.NewSessionsDB()
	authDB := auth.NewAuthDBMemory(sessionsDB)
	service := NewUserService(authDB)
	handler := NewUserHandler(service)

	// use this if there are route variables
	router := mux.NewRouter()
	router.HandleFunc("/users/{username}/role", handler.UpdateRole).Methods("PATCH")
	router.ServeHTTP(rr, r)

	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	defer rr.Result().Body.Close()

	responseMsg, err := io.ReadAll(rr.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, responseMsg)
	assert.Equal(t, "patrick now has the admin role\n", string(responseMsg))
}

func TestUpdateRoleHandler_NoUser(t *testing.T) {
	rr := httptest.NewRecorder()
	formData := url.Values{}
	// formData.Set("newRole", "admin")

	r, err := http.NewRequest(http.MethodPatch, "/users/patrick/role", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Error(err)
	}

	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	sessionsDB := sessions.NewSessionsDB()
	authDB := auth.NewAuthDBMemory(sessionsDB)
	service := NewUserService(authDB)
	handler := NewUserHandler(service)

	// use this if there are route variables
	router := mux.NewRouter()
	router.HandleFunc("/users/{username}/role", handler.UpdateRole).Methods("PATCH")
	router.ServeHTTP(rr, r)

	assert.Equal(t, http.StatusBadRequest, rr.Result().StatusCode)
	defer rr.Result().Body.Close()

	responseMsg, err := io.ReadAll(rr.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, responseMsg)
	assert.Equal(t, "Missing username or new role\n", string(responseMsg))
}
