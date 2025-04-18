package document_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"server-ids/internal/auth"
	"server-ids/internal/database"
	"server-ids/internal/document"
	"server-ids/internal/models"
	"server-ids/internal/sessions"
	"server-ids/internal/template"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

// integration test: HTTP, service, and db interaction
func TestGetDocsHandler(t *testing.T) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/docs", nil)
	if err != nil {
		t.Error(err)
	}

	db := database.CreateMockDB()
	defer db.Close()
	dr := document.NewDocRepository(db)
	service := document.NewDocsService(dr)
	sessions := sessions.NewSessions(db)
	tmpl := template.NewTestTemplate()
	handler := document.NewDocsHandler(service, sessions, tmpl)
	handler.GetDocs(rr, r)

	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	assert.Equal(t, "documents", tmpl.LastRenderedBlock)
	assert.NotEmpty(t, tmpl.LastRenderedData)
}

// integration test: HTTP, service, and db interaction
func TestGetDocHandler(t *testing.T) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/docs/Onboarding%20Document", nil)
	if err != nil {
		t.Error(err)
	}

	db := database.CreateMockDB()
	defer db.Close()
	ar := auth.NewAuthRepository(db)
	dr := document.NewDocRepository(db)
	service := document.NewDocsService(dr)
	sessions := sessions.NewSessions(db)
	tmpl := template.NewTestTemplate()
	handler := document.NewDocsHandler(service, sessions, tmpl)

	key := uuid.New()
	r.AddCookie(&http.Cookie{Name: "session_key", Value: key.String()})
	user := *models.NewUser("funguy123", "admin12345", "admin")
	ar.AddSession(key, user)

	// use this if there are route variables
	router := mux.NewRouter()
	router.HandleFunc("/docs/{title}", handler.GetDoc).Methods("GET")
	router.ServeHTTP(rr, r)

	assert.Equal(t, http.StatusFound, rr.Result().StatusCode)
	assert.NoError(t, err)
	assert.Equal(t, "", tmpl.LastRenderedBlock)
	assert.Nil(t, tmpl.LastRenderedData)
}

// integration test: HTTP, service, and db interaction
func TestGetDocHandler_NotLoggedIn(t *testing.T) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/docs/Onboarding%20Document", nil)
	if err != nil {
		t.Error(err)
	}

	db := database.CreateMockDB()
	defer db.Close()
	dr := document.NewDocRepository(db)
	service := document.NewDocsService(dr)
	sessions := sessions.NewSessions(db)
	tmpl := template.NewTestTemplate()
	handler := document.NewDocsHandler(service, sessions, tmpl)

	// use this if there are route variables
	router := mux.NewRouter()
	router.HandleFunc("/docs/{title}", handler.GetDoc).Methods("GET")
	router.ServeHTTP(rr, r)

	expected := template.ReturnData{
		Error: "Login to access documents",
	}
	assert.Equal(t, "document", tmpl.LastRenderedBlock)
	assert.Equal(t, expected, tmpl.LastRenderedData)
}

// integration test: HTTP, service, and db interaction
func TestGetDocHandler_NotEmployee(t *testing.T) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/docs/Onboarding%20Document", nil)
	if err != nil {
		t.Error(err)
	}

	db := database.CreateMockDB()
	defer db.Close()
	ar := auth.NewAuthRepository(db)
	dr := document.NewDocRepository(db)
	service := document.NewDocsService(dr)
	sessions := sessions.NewSessions(db)
	tmpl := template.NewTestTemplate()
	handler := document.NewDocsHandler(service, sessions, tmpl)

	key := uuid.New()
	r.AddCookie(&http.Cookie{Name: "session_key", Value: key.String()})
	user := *models.NewUser("secure21", "guest12345", "guest")
	ar.AddSession(key, user)

	// use this if there are route variables
	router := mux.NewRouter()
	router.HandleFunc("/docs/{title}", handler.GetDoc).Methods("GET")
	router.ServeHTTP(rr, r)

	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	assert.NoError(t, err)
	assert.Equal(t, "document", tmpl.LastRenderedBlock)
	expected := template.ReturnData{
		Error: "You don't have access to locked documents",
	}
	assert.Equal(t, expected, tmpl.LastRenderedData)
}

// integration test: HTTP, service, and db interaction
func TestGetDocHandler_NotFound(t *testing.T) {
	rr := httptest.NewRecorder()
	formData := url.Values{}
	formData.Set("title", "This Doesn't Exist")

	r, err := http.NewRequest(http.MethodGet, "/docs/NonExistant", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Error(err)
	}

	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	db := database.CreateMockDB()
	defer db.Close()
	dr := document.NewDocRepository(db)
	service := document.NewDocsService(dr)
	sessions := sessions.NewSessions(db)
	tmpl := template.NewTestTemplate()
	handler := document.NewDocsHandler(service, sessions, tmpl)

	// use this if there are route variables
	router := mux.NewRouter()
	router.HandleFunc("/docs/{title}", handler.GetDoc).Methods("GET")
	router.ServeHTTP(rr, r)

	expected := template.ReturnData{
		Error: "Problem occured retreving the document: document 'NonExistant' not found",
	}
	assert.Equal(t, "document", tmpl.LastRenderedBlock)
	assert.Equal(t, expected, tmpl.LastRenderedData)
}
