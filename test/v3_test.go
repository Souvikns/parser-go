package test

import (
	"encoding/json"
	"testing"

	"github.com/Souvikns/parser-go/pkg/models/v3"
	"github.com/stretchr/testify/assert"
)

func TestAsyncAPI(t *testing.T) {
	specString := `{"asyncapi": "3.0.0"}`
	var asyncapi v3.AsyncAPIDocument
	err := json.Unmarshal([]byte(specString), &asyncapi)

	assert.Nil(t, err)
	assert.True(t, asyncapi.Version() == "3.0.0")
}