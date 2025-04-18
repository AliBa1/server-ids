package middleware_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"server-ids/internal/auth"
	"server-ids/internal/database"
	"server-ids/internal/middleware"
	"server-ids/internal/sessions"
	"server-ids/internal/user"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewMiddleware(t *testing.T) {
	sessions := &sessions.Sessions{}
	mw := middleware.NewMiddleware(sessions)

	assert.NotNil(t, mw)
	assert.Equal(t, sessions, mw.Sessions)
}

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

	db := database.CreateMockDB()
	defer db.Close()
	sessions := sessions.NewSessions(db)
	mw := &middleware.Middleware{Sessions: sessions}
	wrappedHandler := mw.Logger(testHandler)

	var logOutput strings.Builder
	log.SetOutput(&logOutput)

	wrappedHandler.ServeHTTP(rr, r)

	log.SetOutput(nil)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.True(t, strings.Contains(logOutput.String(), "- HTTP/1.1 GET Request @ URL: /testing"))
}

func TestAuthorization(t *testing.T) {
	db := database.CreateMockDB()
	defer db.Close()
	sessions := sessions.NewSessions(db)
	ur := user.NewUserRepository(db)
	ar := auth.NewAuthRepository(db)
	key := uuid.New()
	user, err := ur.GetUser("funguy123")
	assert.NoError(t, err)
	ar.AddSession(key, *user)

	r, err := http.NewRequest(http.MethodGet, "/testing", nil)
	if err != nil {
		t.Error(err)
	}
	r.AddCookie(&http.Cookie{Name: "session_key", Value: key.String()})
	rr := httptest.NewRecorder()

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	mw := &middleware.Middleware{Sessions: sessions}
	wrappedHandler := mw.Authorization(testHandler)
	wrappedHandler.ServeHTTP(rr, r)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestAuthorization_Unauthorized(t *testing.T) {
	db := database.CreateMockDB()
	defer db.Close()
	sessions := sessions.NewSessions(db)
	ur := user.NewUserRepository(db)
	ar := auth.NewAuthRepository(db)
	key := uuid.New()
	user, err := ur.GetUser("funguy123")
	assert.NoError(t, err)
	ar.AddSession(key, *user)

	r, err := http.NewRequest(http.MethodGet, "/testing", nil)
	if err != nil {
		t.Error(err)
	}
	rr := httptest.NewRecorder()

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	mw := &middleware.Middleware{Sessions: sessions}
	wrappedHandler := mw.Authorization(testHandler)
	wrappedHandler.ServeHTTP(rr, r)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)

	// should've used this elsewhere as well
	// assert.HTTPStatusCode(t, wrappedHandler, http.MethodGet, "/testing", nil, http.StatusUnauthorized)
}

func TestIDS(t *testing.T) {
	db := database.CreateMockDB()
	defer db.Close()
	sessions := sessions.NewSessions(db)

	formData := url.Values{}
	formData.Set("username", "testuser")

	r, err := http.NewRequest(http.MethodPost, "/testing", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Error(err)
	}

	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	mw := &middleware.Middleware{Sessions: sessions}
	wrappedHandler := mw.IDS(testHandler)
	wrappedHandler.ServeHTTP(rr, r)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestIDS_NoData(t *testing.T) {
	db := database.CreateMockDB()
	defer db.Close()
	sessions := sessions.NewSessions(db)

	r, err := http.NewRequest(http.MethodGet, "/testing", nil)
	if err != nil {
		t.Error(err)
	}
	rr := httptest.NewRecorder()

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})

	mw := &middleware.Middleware{Sessions: sessions}
	wrappedHandler := mw.IDS(testHandler)
	wrappedHandler.ServeHTTP(rr, r)

	assert.Equal(t, http.StatusOK, rr.Code)
}
