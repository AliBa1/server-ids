package document

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

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
	service := NewDocsService(db)
	handler := NewDocsHandler(service)
	handler.GetDocs(rr, r)

	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	defer rr.Result().Body.Close()

	responseMsg, err := io.ReadAll(rr.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, responseMsg)
}

// // integration test: HTTP, service, and db interaction
// func TestGetDocHandler(t *testing.T) {
// 	rr := httptest.NewRecorder()

// 	r, err := http.NewRequest(http.MethodGet, "/docs/Onboarding%20Document", nil)
// 	log.Printf("Path: %s", r.URL.Path)

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	db := NewDocsDBMemory()
// 	service := NewDocsService(db)
// 	handler := NewDocsHandler(service)
// 	handler.GetDoc(rr, r)

// 	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
// 	defer rr.Result().Body.Close()

// 	responseMsg, err := io.ReadAll(rr.Body)
// 	assert.NoError(t, err)
// 	assert.NotEmpty(t, responseMsg)
// }

// // integration test: HTTP, service, and db interaction
// func TestGetDocHandler_NotFound(t *testing.T) {
// 	rr := httptest.NewRecorder()
// 	formData := url.Values{}
// 	formData.Set("title", "This Doesn't Exist")

// 	r, err := http.NewRequest(http.MethodGet, "/docs/HTTHTHT/", strings.NewReader(formData.Encode()))
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	db := NewDocsDBMemory()
// 	service := NewDocsService(db)
// 	handler := NewDocsHandler(service)
// 	handler.GetDoc(rr, r)

// 	assert.Equal(t, http.StatusInternalServerError, rr.Result().StatusCode)
// 	defer rr.Result().Body.Close()

// 	responseMsg, err := io.ReadAll(rr.Body)
// 	assert.NoError(t, err)
// 	assert.Equal(t, "document titled 'This Doesn't Exist' not found\n", string(responseMsg))
// }
