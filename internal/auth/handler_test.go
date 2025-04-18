package auth_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"server-ids/internal/auth"
	"server-ids/internal/database"
	"server-ids/internal/template"
	"server-ids/internal/user"
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

	r, err := http.NewRequest(http.MethodPost, "/login", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Error(err)
	}

	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	db := database.CreateMockDB()
	defer db.Close()
	ar := auth.NewAuthRepository(db)
	ur := user.NewUserRepository(db)
	service := auth.NewAuthService(ar, ur)
	tmpl := template.NewTestTemplate()
	handler := auth.NewAuthHandler(service, tmpl)
	handler.PostLogin(rr, r)

	assert.Equal(t, http.StatusFound, rr.Result().StatusCode)
	assert.NoError(t, err)
	assert.Nil(t, tmpl.LastRenderedData)
}

// integration test: HTTP, service, and db interaction
func TestPostLogin_MissingPassword(t *testing.T) {
	rr := httptest.NewRecorder()
	formData := url.Values{}
	formData.Set("username", "notarealuser")
	formData.Set("password", "")

	r, err := http.NewRequest(http.MethodPost, "/login", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Error(err)
	}

	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	db := database.CreateMockDB()
	defer db.Close()
	ar := auth.NewAuthRepository(db)
	ur := user.NewUserRepository(db)
	service := auth.NewAuthService(ar, ur)
	tmpl := template.NewTestTemplate()
	handler := auth.NewAuthHandler(service, tmpl)
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

	r, err := http.NewRequest(http.MethodPost, "/login", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Error(err)
	}

	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	db := database.CreateMockDB()
	defer db.Close()
	ar := auth.NewAuthRepository(db)
	ur := user.NewUserRepository(db)
	service := auth.NewAuthService(ar, ur)
	tmpl := template.NewTestTemplate()
	handler := auth.NewAuthHandler(service, tmpl)
	handler.PostLogin(rr, r)

	assert.Equal(t, "login", tmpl.LastRenderedBlock)
	assert.NotEmpty(t, tmpl.LastRenderedData)
}

// integration test: HTTP, service, and db interaction
func TestPostRegister(t *testing.T) {
	rr := httptest.NewRecorder()
	formData := url.Values{}
	formData.Set("username", "newuser")
	formData.Set("password", "hiimanewuser")

	r, err := http.NewRequest(http.MethodPost, "/login", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Error(err)
	}

	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	db := database.CreateMockDB()
	defer db.Close()
	ar := auth.NewAuthRepository(db)
	ur := user.NewUserRepository(db)
	service := auth.NewAuthService(ar, ur)
	tmpl := template.NewTestTemplate()
	handler := auth.NewAuthHandler(service, tmpl)
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

	r, err := http.NewRequest(http.MethodPost, "/login", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Error(err)
	}

	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	db := database.CreateMockDB()
	defer db.Close()
	ar := auth.NewAuthRepository(db)
	ur := user.NewUserRepository(db)
	service := auth.NewAuthService(ar, ur)
	tmpl := template.NewTestTemplate()
	handler := auth.NewAuthHandler(service, tmpl)
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

	r, err := http.NewRequest(http.MethodPost, "/login", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Error(err)
	}

	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	db := database.CreateMockDB()
	defer db.Close()
	ar := auth.NewAuthRepository(db)
	ur := user.NewUserRepository(db)
	service := auth.NewAuthService(ar, ur)
	tmpl := template.NewTestTemplate()
	handler := auth.NewAuthHandler(service, tmpl)
	handler.PostRegister(rr, r)

	assert.Equal(t, http.StatusBadRequest, rr.Result().StatusCode)
	defer rr.Result().Body.Close()

	responseMsg, err := io.ReadAll(rr.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, responseMsg)
}

func TestGetUsersHandler(t *testing.T) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodPost, "/users", nil)
	if err != nil {
		t.Error(err)
	}

	db := database.CreateMockDB()
	defer db.Close()
	ar := auth.NewAuthRepository(db)
	ur := user.NewUserRepository(db)
	service := auth.NewAuthService(ar, ur)
	tmpl := template.NewTestTemplate()
	handler := auth.NewAuthHandler(service, tmpl)
	handler.GetUsers(rr, r)

	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	assert.NoError(t, err)
	assert.Equal(t, "users", tmpl.LastRenderedBlock)
	assert.NotNil(t, tmpl.LastRenderedData)
}

// func TestGetUsersHandler_NoUsers(t *testing.T) {
// 	rr := httptest.NewRecorder()

// 	r, err := http.NewRequest(http.MethodPost, "/users", nil)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	db, err := database.CreateMockDB()
// 	assert.NoError(t, err)
// 	ar := NewAuthRepository(db)
// 	ur := user.NewUserRepository(db)
// 	service := NewAuthService(ar, ur)
// 	tmpl := template.NewTestTemplate()
// 	handler := NewAuthHandler(service, tmpl)
// 	handler.GetUsers(rr, r)

// 	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
// 	assert.NoError(t, err)
// 	assert.Equal(t, "users", tmpl.LastRenderedBlock)
// }
