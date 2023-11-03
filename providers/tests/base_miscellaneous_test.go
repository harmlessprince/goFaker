package tests

import (
	"github.com/harmlessprince/goFaker/providers"
	"github.com/stretchr/testify/assert"
	"testing"
)

var baseMiscInstanceTest = &providers.BaseMiscellaneous{}

func TestBoolean(t *testing.T) {
	t.Run("test Boolean returns bool without parameters", func(t *testing.T) {
		expected, _ := baseMiscInstanceTest.Boolean()
		assert.True(t, IsBoolean(expected))
	})
	t.Run("test Boolean returns bool with parameters", func(t *testing.T) {
		expected, _ := baseMiscInstanceTest.Boolean(40)
		assert.True(t, IsBoolean(expected))
	})
}

func IsBoolean(value interface{}) bool {
	_, ok := value.(bool)
	return ok
}
