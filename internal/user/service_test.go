package user

import (
	"server-ids/internal/auth"
	"testing"

	"github.com/stretchr/testify/assert"
)

// integration test: service and db interaction
func TestUpdateRole(t *testing.T) {
	authdb := auth.NewAuthDBMemory()
	service := NewUserService(authdb)
	err := service.UpdateRole("patrick", "admin")

	assert.NoError(t, err)
	
	patrick, err := authdb.GetUser("patrick")
	assert.Equal(t, "admin", patrick.Role)
	assert.NoError(t, err)
}
