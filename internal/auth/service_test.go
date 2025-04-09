package auth

import (
	"server-ids/internal/sessions"
	"testing"

	"github.com/stretchr/testify/assert"
)

// integration test: service and db interaction
func TestLogin(t *testing.T) {
	sessionsDB := sessions.NewSessionsDB()
	db := NewAuthDBMemory(sessionsDB)
	service := NewAuthService(db)
	username := "funguy123"
	password := "admin12345"
	token, err := service.Login(username, password)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	assert.Equal(t, username, db.SessionsDB.Sessions[token])
}

// integration test: service and db interaction
func TestLogin_InvalidUser(t *testing.T) {
	sessionsDB := sessions.NewSessionsDB()
	db := NewAuthDBMemory(sessionsDB)
	service := NewAuthService(db)
	username := "notarealuser"
	password := "admin12345"
	token, err := service.Login(username, password)

	assert.Error(t, err)
	assert.Empty(t, token)
	assert.Empty(t, db.SessionsDB.Sessions[token])
}

// integration test: service and db interaction
func TestLogin_WrongPassword(t *testing.T) {
	sessionsDB := sessions.NewSessionsDB()
	db := NewAuthDBMemory(sessionsDB)
	service := NewAuthService(db)
	username := "funguy123"
	password := "wrongpassword"
	token, err := service.Login(username, password)

	assert.Error(t, err)
	assert.Empty(t, token)
	assert.Empty(t, db.SessionsDB.Sessions[token])
	// add checker for failed login attempts updated
}

// integration test: service and db interaction
func TestRegister(t *testing.T) {
	sessionsDB := sessions.NewSessionsDB()
	db := NewAuthDBMemory(sessionsDB)
	service := NewAuthService(db)
	dbUsersLen := len(sessionsDB.Users)
	// test working pass by reference
	serviceUserLen := len(service.db.SessionsDB.Users)
	username := "newuser"
	password := "iamanewuser"
	err := service.Register(username, password)

	assert.NoError(t, err)
	assert.Equal(t, dbUsersLen+1, len(sessionsDB.Users))
	// test working pass by reference
	assert.Equal(t, serviceUserLen+1, len(service.db.SessionsDB.Users))
}

// integration test: service and db interaction
func TestRegister_UsernameTaken(t *testing.T) {
	sessionsDB := sessions.NewSessionsDB()
	db := NewAuthDBMemory(sessionsDB)
	service := NewAuthService(db)
	dbUsersLen := len(sessionsDB.Users)
	// test working pass by reference
	serviceUserLen := len(service.db.SessionsDB.Users)
	username := "funguy123"
	password := "iamanewuser"
	err := service.Register(username, password)

	assert.Error(t, err)
	assert.Equal(t, dbUsersLen, len(sessionsDB.Users))
	// test working pass by reference
	assert.Equal(t, serviceUserLen, len(service.db.SessionsDB.Users))
}

// integration test: service and db interaction
func TestGetUsersService(t *testing.T) {
	sessionsDB := sessions.NewSessionsDB()
	db := NewAuthDBMemory(sessionsDB)
	service := NewAuthService(db)
	users, err := service.GetAllUsers()

	assert.NoError(t, err)
	assert.NotEmpty(t, users)
}
