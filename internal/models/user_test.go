package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	username := "test"
	password := "password"
	role := "employee"

	user := NewUser(username, password, role)

	assert.NotNil(t, user)
	assert.Equal(t, username, user.Username)
	assert.Equal(t, password, user.Password)
	assert.Equal(t, role, user.Role)
}
