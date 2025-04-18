package sessions_test

import (
	"net/http"
	"server-ids/internal/auth"
	"server-ids/internal/database"
	"server-ids/internal/models"
	"server-ids/internal/sessions"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetSessionUser(t *testing.T) {
	db := database.CreateMockDB()
	defer db.Close()
	sessions := sessions.NewSessions(db)
	ar := auth.NewAuthRepository(db)

	key := uuid.New()
	testUser := models.NewUser("funguy123", "admin12345", "admin")
	ar.AddSession(key, *testUser)
	user, err := sessions.GetSessionUser(key)

	assert.NoError(t, err)
	assert.Equal(t, testUser.Username, user.Username)
}

func TestGetSessionUser_NotExist(t *testing.T) {
	db := database.CreateMockDB()
	defer db.Close()
	sessions := sessions.NewSessions(db)
	ar := auth.NewAuthRepository(db)

	key := uuid.New()
	testUser := models.NewUser("fakeuser", "admin12345", "admin")
	ar.AddSession(key, *testUser)
	user, err := sessions.GetSessionUser(key)

	assert.Error(t, err)
	assert.Nil(t, user)
}

func TestGetUserFromRequest(t *testing.T) {
	db := database.CreateMockDB()
	defer db.Close()
	sessions := sessions.NewSessions(db)
	ar := auth.NewAuthRepository(db)

	r, err := http.NewRequest(http.MethodGet, "/test", nil)
	if err != nil {
		t.Error(err)
	}
	key := uuid.New()
	r.AddCookie(&http.Cookie{Name: "session_key", Value: key.String()})
	testUser := *models.NewUser("funguy123", "admin12345", "admin")
	ar.AddSession(key, testUser)

	user, err := sessions.GetUserFromRequest(r)

	assert.NoError(t, err)
	assert.Equal(t, testUser.Username, user.Username)
}

func TestGetUserFromRequest_NotExist(t *testing.T) {
	db := database.CreateMockDB()
	defer db.Close()
	sessions := sessions.NewSessions(db)
	ar := auth.NewAuthRepository(db)

	r, err := http.NewRequest(http.MethodGet, "/test", nil)
	if err != nil {
		t.Error(err)
	}
	key := uuid.New()
	r.AddCookie(&http.Cookie{Name: "session_key", Value: key.String()})
	testUser := *models.NewUser("fakeuser", "admin12345", "admin")
	ar.AddSession(key, testUser)

	user, err := sessions.GetUserFromRequest(r)

	assert.Error(t, err)
	assert.Nil(t, user)
}

func TestIsUserLoggedIn(t *testing.T) {
	db := database.CreateMockDB()
	defer db.Close()
	sessions := sessions.NewSessions(db)
	ar := auth.NewAuthRepository(db)

	r, err := http.NewRequest(http.MethodGet, "/test", nil)
	if err != nil {
		t.Error(err)
	}
	key := uuid.New()
	r.AddCookie(&http.Cookie{Name: "session_key", Value: key.String()})
	testUser := *models.NewUser("funguy123", "admin12345", "admin")
	ar.AddSession(key, testUser)

	isLoggedIn := sessions.IsUserLoggedIn(r)

	assert.True(t, isLoggedIn)
}

func TestIsUserLoggedIn_Not(t *testing.T) {
	db := database.CreateMockDB()
	defer db.Close()
	sessions := sessions.NewSessions(db)

	r, err := http.NewRequest(http.MethodGet, "/test", nil)
	if err != nil {
		t.Error(err)
	}

	isLoggedIn := sessions.IsUserLoggedIn(r)

	assert.False(t, isLoggedIn)
}

func TestIsUserEmployee(t *testing.T) {
	db := database.CreateMockDB()
	defer db.Close()
	sessions := sessions.NewSessions(db)
	ar := auth.NewAuthRepository(db)

	r, err := http.NewRequest(http.MethodGet, "/test", nil)
	if err != nil {
		t.Error(err)
	}
	key := uuid.New()
	r.AddCookie(&http.Cookie{Name: "session_key", Value: key.String()})
	testUser := *models.NewUser("1819twenty", "emp12345", "employee")
	ar.AddSession(key, testUser)

	isEmployee := sessions.IsUserEmployee(r)

	assert.True(t, isEmployee)
}

func TestIsUserEmployee_Not(t *testing.T) {
	db := database.CreateMockDB()
	defer db.Close()
	sessions := sessions.NewSessions(db)
	ar := auth.NewAuthRepository(db)

	r, err := http.NewRequest(http.MethodGet, "/test", nil)
	if err != nil {
		t.Error(err)
	}
	key := uuid.New()
	r.AddCookie(&http.Cookie{Name: "session_key", Value: key.String()})
	testUser := *models.NewUser("secure21", "guest12345", "guest")
	ar.AddSession(key, testUser)

	isEmployee := sessions.IsUserEmployee(r)

	assert.False(t, isEmployee)
}

func TestIsUserAdmin(t *testing.T) {
	db := database.CreateMockDB()
	defer db.Close()
	sessions := sessions.NewSessions(db)
	ar := auth.NewAuthRepository(db)

	r, err := http.NewRequest(http.MethodGet, "/test", nil)
	if err != nil {
		t.Error(err)
	}
	key := uuid.New()
	r.AddCookie(&http.Cookie{Name: "session_key", Value: key.String()})
	testUser := *models.NewUser("funguy123", "admin12345", "admin")
	ar.AddSession(key, testUser)

	isAdmin := sessions.IsUserAdmin(r)

	assert.True(t, isAdmin)
}

func TestIsUserAdmin_Not(t *testing.T) {
	db := database.CreateMockDB()
	defer db.Close()
	sessions := sessions.NewSessions(db)
	ar := auth.NewAuthRepository(db)

	r, err := http.NewRequest(http.MethodGet, "/test", nil)
	if err != nil {
		t.Error(err)
	}
	key := uuid.New()
	r.AddCookie(&http.Cookie{Name: "session_key", Value: key.String()})
	testUser := *models.NewUser("secure21", "guest12345", "guest")
	ar.AddSession(key, testUser)

	isAdmin := sessions.IsUserAdmin(r)

	assert.False(t, isAdmin)
}
