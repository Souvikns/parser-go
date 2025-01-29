package parser

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLoadDocument(t *testing.T) {
	doc := []byte(`{
		"asyncapi": "3.0.0"
	}`)
	document, err := loadDocument([]byte(doc))
	assert.Nil(t, err)
	assert.NotNil(t, document)
	assert.Equal(t, document.Doc["asyncapi"], "3.0.0")
}

func TestExtractVersion(t *testing.T) {
	doc := map[string]any{
		"asyncapi": "3.0.0",
	}
	version, found := extractVersion(doc)
	assert.True(t, found)
	assert.Equal(t, version, "3.0.0")
}
