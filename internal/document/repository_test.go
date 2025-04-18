package document

import (
	"server-ids/internal/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDocsDB(t *testing.T) {
	db := database.CreateMockDB()
	defer db.Close()
	repo := NewDocRepository(db)
	docs, err := repo.GetDocs()

	assert.NoError(t, err)
	assert.NotEmpty(t, docs)
	assert.Equal(t, "Onboarding Document", docs[0].Title)
}

func TestGetDocDB(t *testing.T) {
	db := database.CreateMockDB()
	defer db.Close()
	repo := NewDocRepository(db)
	doc, err := repo.GetDoc("Onboarding Document")

	assert.NoError(t, err)
	assert.NotEmpty(t, doc)
	assert.Equal(t, "Onboarding Document", doc.Title)
}
