package sessions

import (
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
