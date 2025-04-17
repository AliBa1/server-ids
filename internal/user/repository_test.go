package user

import (
	"server-ids/internal/database"
	"server-ids/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUsersDB(t *testing.T) {
	db, err := database.CreateMockDB()
	assert.NoError(t, err)
	repo := NewUserRepository(db)
	users, err := repo.GetUsers()

	assert.NoError(t, err)
	assert.NotEmpty(t, users)
	assert.Equal(t, "funguy123", users[0].Username)
}

func TestGetUserDB(t *testing.T) {
	db, err := database.CreateMockDB()
	assert.NoError(t, err)
	repo := NewUserRepository(db)
	user, err := repo.GetUser("funguy123")

	assert.NoError(t, err)
	assert.NotEmpty(t, user)
	assert.Equal(t, "funguy123", user.Username)
}

func TestCreateUserDB(t *testing.T) {
	db, err := database.CreateMockDB()
	assert.NoError(t, err)
	repo := NewUserRepository(db)
	user := models.NewUser("testuser", "password", "guest")
	err = repo.CreateUser(*user)

	assert.NoError(t, err)
}

func TestUpdateUserDB(t *testing.T) {
	db, err := database.CreateMockDB()
	assert.NoError(t, err)
	repo := NewUserRepository(db)
	user := models.NewUser("funguy123", "admin12345", "guest")
	err = repo.UpdateUser(*user)

	assert.NoError(t, err)

	user, err = repo.GetUser("funguy123")
	assert.NoError(t, err)
	assert.Equal(t, "guest", user.Role)
}
