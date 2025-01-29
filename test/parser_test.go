package test

import (
	"testing"

	"github.com/Souvikns/parser-go/pkg/parser"
	"github.com/stretchr/testify/assert"
)

func TestV3Spec(t *testing.T) {
	var asyncapi parser.Asyncapi_3_0_0
	err := parser.Parse(string(SPEC_3_0_0), &asyncapi)
	assert.Nil(t, err)
	assert.Equal(t, asyncapi.Asyncapi, "3.0.0")
}