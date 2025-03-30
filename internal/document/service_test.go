package document

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// integration test: service and db interaction
func TestGetDocs(t *testing.T) {
	db := NewDocsDBMemory()
	service := NewDocsService(db)
	docs, err := service.GetDocs()

	assert.NoError(t, err)
	assert.NotEmpty(t, docs)
}

// func TestDisplayDocs(t *testing.T) {
// 	db := NewDocsDBMemory()
// 	service := NewDocsService(db)
// 	docs, err := service.GetDocs()
// 	err = service.DisplayDocs(docs)

// 	assert.NoError(t, err)
// }

