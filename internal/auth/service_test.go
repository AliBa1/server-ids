package auth

import (
	"server-ids/internal/database"
	"server-ids/internal/sessions"
	"server-ids/internal/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

// integration test: service and db interaction
func TestLogin(t *testing.T) {
	db := database.CreateMockDB()
	defer db.Close()
	ur := user.NewUserRepository(db)
	ar := NewAuthRepository(db)
	service := NewAuthService(ar, ur)
	sessions := sessions.NewSessions(db)

	username := "funguy123"
	password := "admin12345"
	key, err := service.Login(username, password)

	assert.NoError(t, err)
	assert.NotEmpty(t, key)
	loggedInUser, err := sessions.GetSessionUser(key)
	assert.NoError(t, err)
	assert.Equal(t, username, loggedInUser.Username)
}

// integration test: service and db interaction
func TestLogin_InvalidUser(t *testing.T) {
	db := database.CreateMockDB()
	defer db.Close()
	ur := user.NewUserRepository(db)
	ar := NewAuthRepository(db)
	service := NewAuthService(ar, ur)

	username := "notarealuser"
	password := "admin12345"
	key, err := service.Login(username, password)

	assert.Error(t, err)
	assert.Empty(t, key)
}

// integration test: service and db interaction
func TestLogin_WrongPassword(t *testing.T) {
	db := database.CreateMockDB()
	defer db.Close()
	ur := user.NewUserRepository(db)
	ar := NewAuthRepository(db)
	service := NewAuthService(ar, ur)

	username := "funguy123"
	password := "wrongpassword"
	key, err := service.Login(username, password)

	assert.Error(t, err)
	assert.Empty(t, key)
	// add checker for failed login attempts updated
}

// integration test: service and db interaction
func TestRegister(t *testing.T) {
	db := database.CreateMockDB()
	defer db.Close()
	ur := user.NewUserRepository(db)
	ar := NewAuthRepository(db)
	service := NewAuthService(ar, ur)

	ogUsers, err := ur.GetUsers()
	assert.NoError(t, err)

	username := "newuser"
	password := "iamanewuser"
	err = service.Register(username, password)
	assert.NoError(t, err)

	curUsers, err := ur.GetUsers()
	assert.NoError(t, err)

	assert.Equal(t, len(ogUsers)+1, len(curUsers))
}

// integration test: service and db interaction
func TestRegister_UsernameTaken(t *testing.T) {
	db := database.CreateMockDB()
	defer db.Close()
	ur := user.NewUserRepository(db)
	ar := NewAuthRepository(db)
	service := NewAuthService(ar, ur)

	ogUsers, err := ur.GetUsers()
	assert.NoError(t, err)

	username := "funguy123"
	password := "iamanewuser"
	err = service.Register(username, password)
	assert.Error(t, err)

	curUsers, err := ur.GetUsers()
	assert.NoError(t, err)
	assert.Equal(t, len(ogUsers), len(curUsers))
}
