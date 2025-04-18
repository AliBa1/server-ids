package user_test

import (
	"server-ids/internal/database"
	"server-ids/internal/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
