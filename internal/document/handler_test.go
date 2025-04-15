package document

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"server-ids/internal/sessions"
	"server-ids/internal/template"
	"strings"
	"testing"

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

	db := NewDocsDBMemory()
	sessionsDB := sessions.NewSessionsDB()
	service := NewDocsService(db)
	tmpl := template.NewTestTemplate()
	handler := NewDocsHandler(service, sessionsDB, tmpl)
	handler.GetDocs(rr, r)

	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	assert.Equal(t, "documents", tmpl.LastRenderedBlock)
	assert.NotEmpty(t, tmpl.LastRenderedData)
}

// integration test: HTTP, service, and db interaction
func TestGetDocHandler(t *testing.T) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/docs/First%20Doc%20Ever", nil)
	if err != nil {
		t.Error(err)
	}

	db := NewDocsDBMemory()
	sessionsDB := sessions.NewSessionsDB()
	service := NewDocsService(db)
	tmpl := template.NewTestTemplate()
	handler := NewDocsHandler(service, sessionsDB, tmpl)

	// use this if there are route variables
	router := mux.NewRouter()
	router.HandleFunc("/docs/{title}", handler.GetDoc).Methods("GET")
	router.ServeHTTP(rr, r)

	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	assert.NoError(t, err)
	assert.Equal(t, "document", tmpl.LastRenderedBlock)
	assert.NotNil(t, tmpl.LastRenderedData)
}

// integration test: HTTP, service, and db interaction
func TestGetDocHandler_NotLoggedIn(t *testing.T) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/docs/Onboarding%20Document", nil)
	if err != nil {
		t.Error(err)
	}

	db := NewDocsDBMemory()
	sessionsDB := sessions.NewSessionsDB()
	service := NewDocsService(db)
	tmpl := template.NewTestTemplate()
	handler := NewDocsHandler(service, sessionsDB, tmpl)

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
func TestGetDocHandler_NotFound(t *testing.T) {
	rr := httptest.NewRecorder()
	formData := url.Values{}
	formData.Set("title", "This Doesn't Exist")

	r, err := http.NewRequest(http.MethodGet, "/docs/NonExistant", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Error(err)
	}

	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	db := NewDocsDBMemory()
	sessionsDB := sessions.NewSessionsDB()
	service := NewDocsService(db)
	tmpl := template.NewTestTemplate()
	handler := NewDocsHandler(service, sessionsDB, tmpl)

	// use this if there are route variables
	router := mux.NewRouter()
	router.HandleFunc("/docs/{title}", handler.GetDoc).Methods("GET")
	router.ServeHTTP(rr, r)

	expected := template.ReturnData{
		Error: "Problem occured retreving the document: document titled 'NonExistant' not found",
	}
	assert.Equal(t, "document", tmpl.LastRenderedBlock)
	assert.Equal(t, expected, tmpl.LastRenderedData)
}
