package sessions

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAddLoginKey(t *testing.T) {
	sessionsDB := NewSessionsDB()
	token := uuid.New()
	username := "funguy123"
	sessionsDB.AddSession(token, username)

	assert.Contains(t, sessionsDB.Sessions, token)
	assert.Equal(t, sessionsDB.Sessions[token], username)
}

func TestGetUsername(t *testing.T) {
	sessionsDB := NewSessionsDB()
	token := uuid.New()
	username := "funguy123"
	sessionsDB.AddSession(token, username)
	stored_username, err := sessionsDB.GetUsername(token)

	assert.NoError(t, err)
	assert.Contains(t, sessionsDB.Sessions, token)
	assert.Equal(t, username, stored_username)
}

func TestGetUsername_NotExist(t *testing.T) {
	sessionsDB := NewSessionsDB()
	token := uuid.New()
	_, err := sessionsDB.GetUsername(token)

	assert.Error(t, err)
}

func TestGetUser(t *testing.T) {
	sessionsDB := NewSessionsDB()

	username := "funguy123"
	user, err := sessionsDB.GetUser(username)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, username, user.Username)
}

func TestGetUser_NotExist(t *testing.T) {
	sessionsDB := NewSessionsDB()

	username := "usernotreal"
	user, err := sessionsDB.GetUser(username)

	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, fmt.Sprintf("user '%s' not found", username), err.Error())
}

func TestIsUserLoggedIn(t *testing.T) {
	sessionsDB := NewSessionsDB()

	r := httptest.NewRequest("GET", "/", nil)
	sessionKey := uuid.New()
	username := "funguy123"
	sessionsDB.AddSession(sessionKey, username)
	r.AddCookie(&http.Cookie{
		Name:  "session_key",
		Value: sessionKey.String(),
	})

	assert.True(t, sessionsDB.IsUserLoggedIn(r))
}

func TestIsUserLoggedIn_NotLoggedIn(t *testing.T) {
	sessionsDB := NewSessionsDB()

	invalidReq := httptest.NewRequest("GET", "/", nil)
	sessionKey := uuid.New()
	invalidReq.AddCookie(&http.Cookie{
		Name:  "session_key",
		Value: sessionKey.String(),
	})

	assert.False(t, sessionsDB.IsUserLoggedIn(invalidReq))
}
