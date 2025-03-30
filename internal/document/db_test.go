package document

import (
	"server-ids/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDocs(t *testing.T) {
	db := NewDocsDBMemory()
	docs, err := db.GetAllDocs()

	assert.NoError(t, err)
	assert.NotEmpty(t, docs)
}

func TestGetDocs_Empty(t *testing.T) {
	db := NewDocsDBMemory()
	db.Documents = []models.Document{}
	docs, err := db.GetAllDocs()

	assert.NoError(t, err)
	assert.Empty(t, docs)
}

func TestGetDoc(t *testing.T) {
	db := NewDocsDBMemory()
	docs, err := db.GetDoc("Top Secret Case Study #1")

	assert.NoError(t, err)
	assert.NotEmpty(t, docs)
}

func TestGetDoc_NotFound(t *testing.T) {
	db := NewDocsDBMemory()
	docs, err := db.GetDoc("This Doesn't Exist")

	assert.Error(t, err)
	assert.Empty(t, docs)
}
