package scheme

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLoadSchemes(t *testing.T) {
	_, err := LoadSchemes()
	assert.Nil(t, err)
}