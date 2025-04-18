package user_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"server-ids/internal/auth"
	"server-ids/internal/database"
	"server-ids/internal/models"
	"server-ids/internal/sessions"
	"server-ids/internal/template"
	"server-ids/internal/user"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

// integration test: HTTP, service, and db interaction
func TestGetUsersHandler(t *testing.T) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodPost, "/users", nil)
	if err != nil {
		t.Error(err)
	}

	db := database.CreateMockDB()
	defer db.Close()
	ur := user.NewUserRepository(db)
	service := user.NewUserService(ur)
	tmpl := template.NewTestTemplate()
	sessions := sessions.NewSessions(db)
	handler := user.NewUserHandler(service, tmpl, sessions)
	handler.GetUsers(rr, r)

	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	assert.NoError(t, err)
	assert.Equal(t, "users", tmpl.LastRenderedBlock)
	assert.NotNil(t, tmpl.LastRenderedData)
}

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
	ar := auth.NewAuthRepository(db)
	ur := user.NewUserRepository(db)
	service := user.NewUserService(ur)
	tmpl := template.NewTestTemplate()
	handler := user.NewUserHandler(service, tmpl, sessions)

	key := uuid.New()
	r.AddCookie(&http.Cookie{Name: "session_key", Value: key.String()})
	testUser := *models.NewUser("funguy123", "admin12345", "admin")
	ar.AddSession(key, testUser)

	// use this if there are route variables
	router := mux.NewRouter()
	router.HandleFunc("/users/{username}/role", handler.UpdateRole).Methods("PATCH")
	router.ServeHTTP(rr, r)

	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	assert.NoError(t, err)
	assert.Equal(t, "users", tmpl.LastRenderedBlock)
	assert.NotNil(t, tmpl.LastRenderedData)
}

func TestUpdateRoleHandler_NotAdmin(t *testing.T) {
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
	ar := auth.NewAuthRepository(db)
	ur := user.NewUserRepository(db)
	service := user.NewUserService(ur)
	tmpl := template.NewTestTemplate()
	handler := user.NewUserHandler(service, tmpl, sessions)

	key := uuid.New()
	r.AddCookie(&http.Cookie{Name: "session_key", Value: key.String()})
	testUser := *models.NewUser("1819twenty", "emp12345", "employee")
	ar.AddSession(key, testUser)

	// use this if there are route variables
	router := mux.NewRouter()
	router.HandleFunc("/users/{username}/role", handler.UpdateRole).Methods("PATCH")
	router.ServeHTTP(rr, r)

	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	assert.NoError(t, err)
	assert.Equal(t, "users", tmpl.LastRenderedBlock)
	assert.NotNil(t, tmpl.LastRenderedData)
}

func TestUpdateRoleHandler_NotExistingUser(t *testing.T) {
	rr := httptest.NewRecorder()
	formData := url.Values{}
	formData.Set("newRole", "admin")

	r, err := http.NewRequest(http.MethodPatch, "/users/iamnotreal/role", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Error(err)
	}
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	db := database.CreateMockDB()
	defer db.Close()
	sessions := sessions.NewSessions(db)
	ar := auth.NewAuthRepository(db)
	ur := user.NewUserRepository(db)
	service := user.NewUserService(ur)
	tmpl := template.NewTestTemplate()
	handler := user.NewUserHandler(service, tmpl, sessions)

	key := uuid.New()
	r.AddCookie(&http.Cookie{Name: "session_key", Value: key.String()})
	testUser := *models.NewUser("funguy123", "admin12345", "admin")
	ar.AddSession(key, testUser)

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
