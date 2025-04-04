package middleware

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	r, err := http.NewRequest(http.MethodGet, "/testing", nil)
	if err != nil {
		t.Error(err)
	}
	rr := httptest.NewRecorder()

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	middleware := &Middleware{}
	wrappedHandler := middleware.Logger(testHandler)

	var logOutput strings.Builder
	log.SetOutput(&logOutput)

	wrappedHandler.ServeHTTP(rr, r)

	log.SetOutput(nil)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.True(t, strings.Contains(logOutput.String(), "- HTTP/1.1 GET Request @ URL: /testing"))
}

func TestAuthorization(t *testing.T) {
	sessionId := uuid.New()
	sessions := map[uuid.UUID]string{sessionId: "funguy123"}

	r, err := http.NewRequest(http.MethodGet, "/testing", nil)
	if err != nil {
		t.Error(err)
	}
	// r.AddCookie(&http.Cookie{Name: "session_key", Value: "valid_session"})
	r.AddCookie(&http.Cookie{Name: "session_key", Value: sessionId.String()})
	rr := httptest.NewRecorder()

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	middleware := &Middleware{}
	wrappedHandler := middleware.Authorization(testHandler, sessions)
	wrappedHandler.ServeHTTP(rr, r)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestAuthorization_Unauthorized(t *testing.T) {
	// sessionId := uuid.New()
	sessions := map[uuid.UUID]string{}

	r, err := http.NewRequest(http.MethodGet, "/testing", nil)
	if err != nil {
		t.Error(err)
	}
	rr := httptest.NewRecorder()

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	middleware := &Middleware{}
	wrappedHandler := middleware.Authorization(testHandler, sessions)
	wrappedHandler.ServeHTTP(rr, r)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)

	// should've used this elsewhere as well
	// assert.HTTPStatusCode(t, wrappedHandler, http.MethodGet, "/testing", nil, http.StatusUnauthorized)
}
