package user_test

import (
	"server-ids/internal/database"
	"server-ids/internal/models"
	"server-ids/internal/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

// integration test: service and db interaction
func TestGetUsersService(t *testing.T) {
	db := database.CreateMockDB()
	defer db.Close()
	ur := user.NewUserRepository(db)
	service := user.NewUserService(ur)
	users, err := service.GetAllUsers()

	assert.NoError(t, err)
	assert.NotEmpty(t, users)
}

// integration test: service and db interaction
func TestUpdateRole(t *testing.T) {
	db := database.CreateMockDB()
	defer db.Close()
	ur := user.NewUserRepository(db)
	service := user.NewUserService(ur)
	err := service.UpdateRole("patrick", "admin")

	assert.NoError(t, err)

	patrick, err := ur.GetUser("patrick")
	assert.Equal(t, "admin", patrick.Role)
	assert.NoError(t, err)
}

// integration test: service and db interaction
func TestUpdateRole_NotExist(t *testing.T) {
	db := database.CreateMockDB()
	defer db.Close()
	ur := user.NewUserRepository(db)
	service := user.NewUserService(ur)
	err := service.UpdateRole("iamnotauser", "admin")

	assert.Error(t, err)
}

func TestCanEditRole(t *testing.T) {
	db := database.CreateMockDB()
	defer db.Close()
	ur := user.NewUserRepository(db)
	service := user.NewUserService(ur)
	user := models.NewUser("funguy123", "admin12345", "admin")
	canEdit := service.CanEditRole(*user)

	assert.True(t, canEdit)
}

func TestCanEditRole_CanNot(t *testing.T) {
	db := database.CreateMockDB()
	defer db.Close()
	ur := user.NewUserRepository(db)
	service := user.NewUserService(ur)
	user := *models.NewUser("1819twenty", "emp12345", "employee")
	canEdit := service.CanEditRole(user)

	assert.False(t, canEdit)
}
