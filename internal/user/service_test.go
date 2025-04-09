package user

import (
	"server-ids/internal/auth"
	"server-ids/internal/sessions"
	"testing"

	"github.com/stretchr/testify/assert"
)

// integration test: service and db interaction
func TestUpdateRole(t *testing.T) {
	sessionsDB := sessions.NewSessionsDB()
	authDB := auth.NewAuthDBMemory(sessionsDB)
	service := NewUserService(authDB)
	err := service.UpdateRole("patrick", "admin")

	assert.NoError(t, err)

	patrick, err := authDB.GetUser("patrick")
	assert.Equal(t, "admin", patrick.Role)
	assert.NoError(t, err)
}

// integration test: service and db interaction
func TestUpdateRole_NotExist(t *testing.T) {
	sessionsDB := sessions.NewSessionsDB()
	authDB := auth.NewAuthDBMemory(sessionsDB)
	service := NewUserService(authDB)
	err := service.UpdateRole("iamnotauser", "admin")

	assert.Error(t, err)
}
