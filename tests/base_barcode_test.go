package tests

import (
	"github.com/harmlessprince/goFaker/calculator"
	"github.com/harmlessprince/goFaker/providers"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

var eanCalc = &calculator.Ean{}
var isbnCalc = &calculator.Isbn{}

//	func TestName(t *testing.T) {
//		t.Run("description", func(t *testing.T) {})
//		t.Run("description", func(t *testing.T) {})
//	}
var barcodeInstanceTest = &providers.BaseBarcode{}

func TestEan8(t *testing.T) {
	t.Run("validate Ean8 pattern", func(t *testing.T) {
		//matches a string that consists of exactly 8 digits
		re := regexp.MustCompile(`^\d{8}$`)
		expected, _ := barcodeInstanceTest.Ean8()
		assert.Regexp(t, re, expected)
	})
	t.Run("validate Ean8 is valid", func(t *testing.T) {
		expected, _ := barcodeInstanceTest.Ean8()
		isValid, _ := eanCalc.IsValid(expected)
		assert.True(t, isValid)
	})
}

func TestEan13(t *testing.T) {
	t.Run("validate Ean13 pattern", func(t *testing.T) {
		//matches a string that consists of exactly 8 digits
		re := regexp.MustCompile(`^\d{13}$`)
		expected, _ := barcodeInstanceTest.Ean13()
		assert.Regexp(t, re, expected)
	})
	t.Run("validate Ean13 is valid", func(t *testing.T) {
		expected, _ := barcodeInstanceTest.Ean13()
		isValid, _ := eanCalc.IsValid(expected)
		assert.True(t, isValid)
	})
}

func TestIsbn10(t *testing.T) {
	t.Run("validate Isbn10 pattern", func(t *testing.T) {
		//matches a string that consists of exactly 8 digits
		re := regexp.MustCompile(`^\d{9}[0-9X]$`)
		expected, _ := barcodeInstanceTest.Isbn10()
		assert.Regexp(t, re, expected)
	})
	t.Run("validate Isbn10 is valid", func(t *testing.T) {
		expected, _ := barcodeInstanceTest.Isbn10()
		isValid, _ := isbnCalc.IsValid(expected)
		assert.True(t, isValid)
	})
}

func TestIsbn13(t *testing.T) {
	t.Run("validate Isbn13 pattern", func(t *testing.T) {
		//matches a string that consists of exactly 8 digits
		re := regexp.MustCompile(`^\d{13}$`)
		expected, _ := barcodeInstanceTest.Isbn13()
		assert.Regexp(t, re, expected)
	})
	t.Run("validate Isbn13 is valid", func(t *testing.T) {
		expected, _ := barcodeInstanceTest.Isbn13()
		isValid, _ := eanCalc.IsValid(expected)
		assert.True(t, isValid)
	})
}
