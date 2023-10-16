package helpers

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

//func TestName(t *testing.T) {
//	t.Run("description", func(t *testing.T) {})
//	t.Run("description", func(t *testing.T) {})
//}

func TestRandomElement(t *testing.T) {
	t.Run("valid length of random keys", func(t *testing.T) {
		h := Helper{}
		ages := []interface{}{1, 2, 4, 8, 4}
		randomElement, _ := h.RandomElementsKeys(ages, 3, false)
		assert.Equal(t, len(randomElement.([]int)), 3)
	})
	t.Run("invalid random key size returns error ", func(t *testing.T) {
		h := Helper{}
		ages := []interface{}{1, 2, 4, 8, 4}
		_, err := h.RandomElementsKeys(ages, 9, true)
		assert.Error(t, err)
	})
}

func TestRandomNumberBetween(t *testing.T) {
	t.Run("number generated is within bound", func(t *testing.T) {
		h := Helper{}
		min := 2
		max := 20
		randomNumber, _ := h.RandomNumberBetween(min, max)
		assert.True(t, randomNumber >= min && randomNumber <= max)
	})
	t.Run("min value must be greater than max value", func(t *testing.T) {
		h := Helper{}
		min := 20
		max := 20
		_, err := h.RandomNumberBetween(min, max)
		assert.Error(t, err)
	})
}
