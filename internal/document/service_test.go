package document

import (
	"server-ids/internal/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

// integration test: service and db interaction
func TestGetDocsService(t *testing.T) {
	db := database.CreateMockDB()
	dr := NewDocRepository(db)
	service := NewDocsService(dr)
	docs, err := service.GetDocs()

	assert.NoError(t, err)
	assert.NotEmpty(t, docs)
}

// integration test: service and db interaction
func TestGetDocService(t *testing.T) {
	db := database.CreateMockDB()
	dr := NewDocRepository(db)
	service := NewDocsService(dr)
	doc, err := service.GetDoc("Top Secret Case Study")

	assert.NoError(t, err)
	assert.NotEmpty(t, doc)
}

// integration test: service and db interaction
func TestGetDocService_NotFound(t *testing.T) {
	db := database.CreateMockDB()
	dr := NewDocRepository(db)
	service := NewDocsService(dr)
	doc, err := service.GetDoc("Does Not Exist")

	assert.Error(t, err)
	assert.Empty(t, doc)
}
