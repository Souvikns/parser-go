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

func TestInfoObject(t *testing.T) {
	specString := `{"title": "simple server", "version": "0.0.1"}`
	var info v3.Info
	err := json.Unmarshal([]byte(specString), &info)

	assert.Nil(t, err)
	assert.True(t, info.Title() == "simple server")
}

func TestContact(t *testing.T) {
	specString := `{"asyncapi": "3.0.0", "info": {"title": "", "version": "0.0.1", "contact": {"name": "Souvik"}}}`
	var asyncapi v3.AsyncAPIDocument
	err := json.Unmarshal([]byte(specString), &asyncapi)
	assert.Nil(t, err)
	assert.True(t,  asyncapi.Info().Contact().Name() == "Souvik")
}