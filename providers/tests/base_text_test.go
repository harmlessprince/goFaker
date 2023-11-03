package tests

import (
	"github.com/harmlessprince/goFaker/providers"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

var baseTextInstanceTest = providers.NewBaseText()

func TestRealText(t *testing.T) {
	t.Run("test RealText max length", func(t *testing.T) {
		testWith := []int{10, 20, 50, 70, 90, 120, 150, 200, 500}
		for i := 0; i < len(testWith); i++ {
			expected, _ := baseTextInstanceTest.RealText(providers.RealTextParam{
				MaxNbChars: testWith[i],
			})
			assert.Less(t, len(expected), testWith[i])
		}

	})
	t.Run("test RealText min length", func(t *testing.T) {
		testWith := []int{10, 20, 50, 70, 90, 120, 150, 200, 500}
		for i := 0; i < len(testWith); i++ {
			expected, _ := baseTextInstanceTest.RealText(providers.RealTextParam{
				MaxNbChars: testWith[i],
			})
			result := int(math.Round(float64(testWith[i]) * 0.8))
			assert.GreaterOrEqual(t, len(expected), result)
		}

	})

	t.Run("test RealText max index", func(t *testing.T) {
		_, err := baseTextInstanceTest.RealText(providers.RealTextParam{
			MaxNbChars: 200,
			IndexSize:  11,
		})
		assert.Error(t, err)
	})
}
