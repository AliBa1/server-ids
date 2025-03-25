package auth

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"server-ids/internal/models"
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

	req, err := http.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Error(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	db := NewAuthDBMemory()
	service := NewAuthService(db)
	handler := NewAuthHandler(service)
	handler.PostLogin(rr, req)

	assert.Equal(t, rr.Result().StatusCode, http.StatusOK)
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

	req, err := http.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Error(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	db := NewAuthDBMemory()
	service := NewAuthService(db)
	handler := NewAuthHandler(service)
	handler.PostLogin(rr, req)

	assert.Equal(t, rr.Result().StatusCode, http.StatusBadRequest)
	defer rr.Result().Body.Close()
}

// integration test: HTTP, service, and db interaction
func TestPostLogin_UserNotExist(t *testing.T) {
	rr := httptest.NewRecorder()
	formData := url.Values{}
	formData.Set("username", "notarealuser")
	formData.Set("password", "admin12345")

	req, err := http.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Error(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	db := NewAuthDBMemory()
	service := NewAuthService(db)
	handler := NewAuthHandler(service)
	handler.PostLogin(rr, req)

	assert.Equal(t, rr.Result().StatusCode, http.StatusBadRequest)
	defer rr.Result().Body.Close()
}

// integration test: HTTP, service, and db interaction
func TestPostRegister(t *testing.T) {
	rr := httptest.NewRecorder()
	formData := url.Values{}
	formData.Set("username", "newuser")
	formData.Set("password", "hiimanewuser")

	req, err := http.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Error(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	db := NewAuthDBMemory()
	service := NewAuthService(db)
	handler := NewAuthHandler(service)
	handler.PostRegister(rr, req)

	assert.Equal(t, rr.Result().StatusCode, http.StatusOK)
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

	req, err := http.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Error(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	db := NewAuthDBMemory()
	service := NewAuthService(db)
	handler := NewAuthHandler(service)
	handler.PostRegister(rr, req)

	assert.Equal(t, rr.Result().StatusCode, http.StatusBadRequest)
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

	req, err := http.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Error(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	db := NewAuthDBMemory()
	service := NewAuthService(db)
	handler := NewAuthHandler(service)
	handler.PostRegister(rr, req)

	assert.Equal(t, rr.Result().StatusCode, http.StatusBadRequest)
	defer rr.Result().Body.Close()

	responseMsg, err := io.ReadAll(rr.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, responseMsg)
}

func TestGetUsersHandler(t *testing.T) {
	rr := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodPost, "/auth/users", nil)
	if err != nil {
		t.Error(err)
	}

	db := NewAuthDBMemory()
	service := NewAuthService(db)
	handler := NewAuthHandler(service)
	handler.GetUsers(rr, req)

	assert.Equal(t, rr.Result().StatusCode, http.StatusOK)
	defer rr.Result().Body.Close()

	responseMsg, err := io.ReadAll(rr.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, responseMsg)
}

func TestGetUsersHandler_NoUsers(t *testing.T) {
	rr := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodPost, "/auth/users", nil)
	if err != nil {
		t.Error(err)
	}

	db := NewAuthDBMemory()
	db.Users = []models.User{}
	service := NewAuthService(db)
	handler := NewAuthHandler(service)
	handler.GetUsers(rr, req)

	assert.Equal(t, rr.Result().StatusCode, http.StatusOK)
	defer rr.Result().Body.Close()

	responseMsg, err := io.ReadAll(rr.Body)
	msgString := string(responseMsg)
	assert.NoError(t, err)
	assert.Equal(t, msgString, "Getting users from the database...\n\nThere are no users\n")
}
