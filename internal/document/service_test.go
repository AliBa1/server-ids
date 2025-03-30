package document

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// integration test: service and db interaction
func TestGetDocsService(t *testing.T) {
	db := NewDocsDBMemory()
	service := NewDocsService(db)
	docs, err := service.GetDocs()

	assert.NoError(t, err)
	assert.NotEmpty(t, docs)
}

// integration test: service and db interaction
func TestGetDocService(t *testing.T)  {
	db := NewDocsDBMemory()
	service := NewDocsService(db)
	doc, err := service.GetDoc("Top Secret Case Study #1")

	assert.NoError(t, err)
	assert.NotEmpty(t, doc)
}

// integration test: service and db interaction
func TestGetDocService_NotFound(t *testing.T)  {
	db := NewDocsDBMemory()
	service := NewDocsService(db)
	doc, err := service.GetDoc("Does Not Exist")

	assert.Error(t, err)
	assert.Empty(t, doc)
}