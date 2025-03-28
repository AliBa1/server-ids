package auth

import (
	"net"
	"server-ids/internal/models"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetUsersDB(t *testing.T) {
	db := NewAuthDBMemory()
	users, err := db.GetUsers()

	assert.NoError(t, err)
	assert.NotEmpty(t, users)
}

func TestGetUsersDB_Empty(t *testing.T) {
	db := NewAuthDBMemory()
	db.Users = []models.User{}
	users, err := db.GetUsers()

	assert.NoError(t, err)
	assert.Empty(t, users)
}

func TestGetUser(t *testing.T) {
	db := NewAuthDBMemory()
	user, err := db.GetUser("funguy123")

	assert.NoError(t, err)
	assert.NotEmpty(t, user)
}

func TestGetUser_NotFound(t *testing.T) {
	db := NewAuthDBMemory()
	user, err := db.GetUser("idonotexist")

	assert.Error(t, err)
	assert.Empty(t, user)
}

func TestCreateUser(t *testing.T) {
	db := NewAuthDBMemory()
	prevUsersLen := len(db.Users)
	newUser := models.User{
		Username:            "newuser",
		Password:            "thisismypassword",
		Role:                "employee",
		LastLoginDate:       time.Now(),
		LastLoginIP:         net.ParseIP("202.28.138.47"),
		FailedLoginAttempts: make(map[string]models.FailedLoginInfo),
	}
	err := db.CreateUser(newUser)

	assert.NoError(t, err)
	assert.Len(t, db.Users, prevUsersLen+1)
	assert.Contains(t, db.Users, newUser)
}

func TestUpdateUser(t *testing.T) {
	db := NewAuthDBMemory()
	user := db.Users[0]
	updatedUser := models.User{
		Username:            "funguy123",
		Password:            "updatedpassword",
		Role:                "employee",
		LastLoginDate:       time.Now(),
		LastLoginIP:         net.ParseIP("202.28.138.47"),
		FailedLoginAttempts: make(map[string]models.FailedLoginInfo),
	}
	err := db.UpdateUser(updatedUser)

	assert.NoError(t, err)
	// replace with id if using id for users
	assert.Equal(t, updatedUser.Username, user.Username)
	assert.NotEqual(t, updatedUser, user)
	assert.Contains(t, db.Users, updatedUser)
}

func TestUpdateUser_NotFound(t *testing.T) {
	db := NewAuthDBMemory()
	user := models.User{
		Username:            "idonotexist",
		Password:            "updatedpassword",
		Role:                "employee",
		LastLoginDate:       time.Now(),
		LastLoginIP:         net.ParseIP("202.28.138.47"),
		FailedLoginAttempts: make(map[string]models.FailedLoginInfo),
	}
	err := db.UpdateUser(user)

	assert.Error(t, err)
}

func TestAddLoginKey(t *testing.T) {
	db := NewAuthDBMemory()
	token := uuid.New()
	username := "funguy123"
	db.AddSessionToken(token, username)

	assert.Contains(t, db.Sessions, token)
	assert.Equal(t, db.Sessions[token], username)
}
