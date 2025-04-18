package database_test

import (
	"server-ids/internal/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDBConection(t *testing.T) {
	db := database.CreateMockDB()
	defer db.Close()

	assert.NotNil(t, db)
}
