package tests

import (
	"fmt"
	"github.com/harmlessprince/goFaker/providers"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var basePaymentInstanceTest = &providers.BasePayment{}

func TestCreditCardType(t *testing.T) {
	t.Run("test CreditCardType returns valid vendor name", func(t *testing.T) {
		expected, _ := basePaymentInstanceTest.CreditCardType()
		l, _ := basePaymentInstanceTest.CreditCardExpirationDate()
		fmt.Println(l)
		assert.Contains(t, []string{"Visa", "Visa Retired", "MasterCard", "American Express", "Discover Card", "JCB"}, expected)
	})
}

func TestCreditCardNumber(t *testing.T) {
	t.Run("test CreditCardNumber returns valid card number for Discover Card", func(t *testing.T) {
		expected, _ := basePaymentInstanceTest.CreditCardNumber(providers.CreditCardNumberParams{CreditCardType: "Discover Card"})
		assert.Regexp(t, "^6011\\d{12}$", expected)
	})
	t.Run("test CreditCardNumber returns valid card number Visa", func(t *testing.T) {
		expected, _ := basePaymentInstanceTest.CreditCardNumber(providers.CreditCardNumberParams{CreditCardType: "Visa"})
		assert.Regexp(t, "^4\\d{15}$", expected)
	})
	t.Run("test CreditCardNumber returns valid card number for  Visa Retired", func(t *testing.T) {
		expected, _ := basePaymentInstanceTest.CreditCardNumber(providers.CreditCardNumberParams{CreditCardType: "Visa Retired"})
		assert.Regexp(t, "^4\\d{12}$", expected)
	})
	t.Run("test CreditCardNumber returns valid card number for  MasterCard", func(t *testing.T) {
		expected, _ := basePaymentInstanceTest.CreditCardNumber(providers.CreditCardNumberParams{CreditCardType: "MasterCard"})
		assert.Regexp(t, "^(5[1-5]|2[2-7])\\d{14}$", expected)
	})
	t.Run("test CreditCardNumber returns valid card number for  JCB", func(t *testing.T) {
		expected, _ := basePaymentInstanceTest.CreditCardNumber(providers.CreditCardNumberParams{CreditCardType: "JCB"})
		assert.Regexp(t, "^35(28|89)\\d{12,15}$", expected)
	})
	t.Run("test CreditCardNumber can format output", func(t *testing.T) {
		expected, _ := basePaymentInstanceTest.CreditCardNumber(providers.CreditCardNumberParams{CreditCardType: "Discover Card", Formatted: true})
		assert.Regexp(t, "^6011-\\d{4}-\\d{4}-\\d{4}$", expected)
	})

}

func TestCreditCardExpirationDate(t *testing.T) {
	t.Run("test CreditCardExpirationDate returns valid date by default", func(t *testing.T) {
		expected, _ := basePaymentInstanceTest.CreditCardExpirationDate()
		assert.Greater(t, expected.Unix(), time.Now().Unix())
	})
}

func TestCreditCardDetails(t *testing.T) {
	t.Run("test CreditCardDetails returns random credit card details", func(t *testing.T) {
		expected, _ := basePaymentInstanceTest.CreditCardDetails()
		expectedKeys := []string{"type", "number", "name", "expirationDate"}
		assert.Equal(t, len(expected), 4)
		assert.ElementsMatch(t, expectedKeys, lo.Keys(expected))
	})
}
