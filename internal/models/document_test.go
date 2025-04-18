package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDocument(t *testing.T) {
	title := "test"
	content := "content"
	isLocked := true

	doc := NewDocument(title, content, isLocked)

	assert.NotNil(t, doc)
	assert.Equal(t, title, doc.Title)
	assert.Equal(t, content, doc.Content)
	assert.Equal(t, isLocked, doc.IsLocked)
}
