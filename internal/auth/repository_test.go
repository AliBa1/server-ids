package auth

import (
	"server-ids/internal/database"
	"server-ids/internal/models"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAddSession(t *testing.T) {
	db := database.CreateMockDB()
	defer db.Close()
	repo := NewAuthRepository(db)
	key := uuid.New()
	user := models.NewUser("testuser", "password", "guest")
	err := repo.AddSession(key, *user)

	assert.NoError(t, err)
}
