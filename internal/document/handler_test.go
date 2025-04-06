package document

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"server-ids/internal/auth"
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
	authDB := auth.NewAuthDBMemory()
	service := NewDocsService(db)
	handler := NewDocsHandler(service, authDB)
	handler.GetDocs(rr, r)

	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	defer rr.Result().Body.Close()

	responseMsg, err := io.ReadAll(rr.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, responseMsg)
}

// integration test: HTTP, service, and db interaction
func TestGetDocHandler(t *testing.T) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/docs/First%20Doc%20Ever", nil)
	if err != nil {
		t.Error(err)
	}

	db := NewDocsDBMemory()
	authDB := auth.NewAuthDBMemory()
	service := NewDocsService(db)
	handler := NewDocsHandler(service, authDB)

	// use this if there are route variables
	router := mux.NewRouter()
	router.HandleFunc("/docs/{title}", handler.GetDoc).Methods("GET")
	router.ServeHTTP(rr, r)

	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	defer rr.Result().Body.Close()

	responseMsg, err := io.ReadAll(rr.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, responseMsg)
}

// integration test: HTTP, service, and db interaction
func TestGetDocHandler_NotLoggedIn(t *testing.T) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/docs/Onboarding%20Document", nil)
	if err != nil {
		t.Error(err)
	}

	db := NewDocsDBMemory()
	authDB := auth.NewAuthDBMemory()
	service := NewDocsService(db)
	handler := NewDocsHandler(service, authDB)

	// use this if there are route variables
	router := mux.NewRouter()
	router.HandleFunc("/docs/{title}", handler.GetDoc).Methods("GET")
	router.ServeHTTP(rr, r)

	assert.Equal(t, http.StatusUnauthorized, rr.Result().StatusCode)
	defer rr.Result().Body.Close()

	responseMsg, err := io.ReadAll(rr.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, responseMsg)
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
	authDB := auth.NewAuthDBMemory()
	service := NewDocsService(db)
	handler := NewDocsHandler(service, authDB)

	// use this if there are route variables
	router := mux.NewRouter()
	router.HandleFunc("/docs/{title}", handler.GetDoc).Methods("GET")
	router.ServeHTTP(rr, r)

	assert.Equal(t, http.StatusInternalServerError, rr.Result().StatusCode)
	defer rr.Result().Body.Close()

	responseMsg, err := io.ReadAll(rr.Body)
	assert.NoError(t, err)
	assert.Equal(t, "document titled 'NonExistant' not found\n", string(responseMsg))
}
