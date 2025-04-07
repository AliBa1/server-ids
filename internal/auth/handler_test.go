package auth

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"server-ids/internal/models"
	"server-ids/internal/sessions"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// integration test: HTTP, service, and db interaction
func TestPostLogin(t *testing.T) {
	rr := httptest.NewRecorder()
	formData := url.Values{}
	formData.Set("username", "funguy123")
	formData.Set("password", "admin12345")

	r, err := http.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Error(err)
	}

	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	sessionsDB := sessions.NewSessionsDB()
	db := NewAuthDBMemory(sessionsDB)
	service := NewAuthService(db)
	handler := NewAuthHandler(service)
	handler.PostLogin(rr, r)

	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	defer rr.Result().Body.Close()

	responseMsg, err := io.ReadAll(rr.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, responseMsg)
}

// integration test: HTTP, service, and db interaction
func TestPostLogin_MissingPassword(t *testing.T) {
	rr := httptest.NewRecorder()
	formData := url.Values{}
	formData.Set("username", "notarealuser")
	formData.Set("password", "")

	r, err := http.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Error(err)
	}

	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	sessionsDB := sessions.NewSessionsDB()
	db := NewAuthDBMemory(sessionsDB)
	service := NewAuthService(db)
	handler := NewAuthHandler(service)
	handler.PostLogin(rr, r)

	assert.Equal(t, http.StatusBadRequest, rr.Result().StatusCode)
	defer rr.Result().Body.Close()
}

// integration test: HTTP, service, and db interaction
func TestPostLogin_UserNotExist(t *testing.T) {
	rr := httptest.NewRecorder()
	formData := url.Values{}
	formData.Set("username", "notarealuser")
	formData.Set("password", "admin12345")

	r, err := http.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Error(err)
	}

	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	sessionsDB := sessions.NewSessionsDB()
	db := NewAuthDBMemory(sessionsDB)
	service := NewAuthService(db)
	handler := NewAuthHandler(service)
	handler.PostLogin(rr, r)

	assert.Equal(t, http.StatusBadRequest, rr.Result().StatusCode)
	defer rr.Result().Body.Close()
}

// integration test: HTTP, service, and db interaction
func TestPostRegister(t *testing.T) {
	rr := httptest.NewRecorder()
	formData := url.Values{}
	formData.Set("username", "newuser")
	formData.Set("password", "hiimanewuser")

	r, err := http.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Error(err)
	}

	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	sessionsDB := sessions.NewSessionsDB()
	db := NewAuthDBMemory(sessionsDB)
	service := NewAuthService(db)
	handler := NewAuthHandler(service)
	handler.PostRegister(rr, r)

	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	defer rr.Result().Body.Close()

	responseMsg, err := io.ReadAll(rr.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, responseMsg)
}

// integration test: HTTP, service, and db interaction
func TestPostRegister_MissingPassword(t *testing.T) {
	rr := httptest.NewRecorder()
	formData := url.Values{}
	formData.Set("username", "newuser")
	formData.Set("password", "")

	r, err := http.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Error(err)
	}

	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	sessionsDB := sessions.NewSessionsDB()
	db := NewAuthDBMemory(sessionsDB)
	service := NewAuthService(db)
	handler := NewAuthHandler(service)
	handler.PostRegister(rr, r)

	assert.Equal(t, http.StatusBadRequest, rr.Result().StatusCode)
	defer rr.Result().Body.Close()

	responseMsg, err := io.ReadAll(rr.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, responseMsg)
}

// integration test: HTTP, service, and db interaction
func TestPostRegister_UserExists(t *testing.T) {
	rr := httptest.NewRecorder()
	formData := url.Values{}
	formData.Set("username", "funguy123")
	formData.Set("password", "admin12345")

	r, err := http.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Error(err)
	}

	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	sessionsDB := sessions.NewSessionsDB()
	db := NewAuthDBMemory(sessionsDB)
	service := NewAuthService(db)
	handler := NewAuthHandler(service)
	handler.PostRegister(rr, r)

	assert.Equal(t, http.StatusBadRequest, rr.Result().StatusCode)
	defer rr.Result().Body.Close()

	responseMsg, err := io.ReadAll(rr.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, responseMsg)
}

func TestGetUsersHandler(t *testing.T) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodPost, "/auth/users", nil)
	if err != nil {
		t.Error(err)
	}

	sessionsDB := sessions.NewSessionsDB()
	db := NewAuthDBMemory(sessionsDB)
	service := NewAuthService(db)
	handler := NewAuthHandler(service)
	handler.GetUsers(rr, r)

	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	defer rr.Result().Body.Close()

	responseMsg, err := io.ReadAll(rr.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, responseMsg)
}

func TestGetUsersHandler_NoUsers(t *testing.T) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodPost, "/auth/users", nil)
	if err != nil {
		t.Error(err)
	}

	sessionsDB := sessions.NewSessionsDB()
	db := NewAuthDBMemory(sessionsDB)
	db.Users = []models.User{}
	service := NewAuthService(db)
	handler := NewAuthHandler(service)
	handler.GetUsers(rr, r)

	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	defer rr.Result().Body.Close()

	responseMsg, err := io.ReadAll(rr.Body)
	msgString := string(responseMsg)
	assert.NoError(t, err)
	// assert.Equal(t, "Getting users from the database...\n\nThere are no users\n", msgString)
	assert.Equal(t, "There are no users\n", msgString)
}
