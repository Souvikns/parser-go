package test

import (
	"testing"

	parser "github.com/Souvikns/parser-go/pkg"
	"github.com/stretchr/testify/assert"
)

func TestV3Spec(t *testing.T) {
	valid, err := parser.Validate(SPEC_3_0_0)
	assert.Nil(t, err)
	assert.True(t, valid)
}
