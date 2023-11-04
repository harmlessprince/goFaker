package providers

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/binary"
	"github.com/harmlessprince/goFaker/constants"
)

type MiscellaneousInterface interface {
	Boolean(chanceOfGettingTrue ...int) (bool, error)
	Md5() (string, error)
	Sha1() (string, error)
	Sha256() (string, error)
	Locale() (string, error)
	CountryCode() (string, error)
	CountryISOAlpha3() (string, error)
	LanguageCode() (string, error)
	CurrencyCode() (string, error)
	Emoji() (string, error)
}

type BaseMiscellaneous struct {
	BaseProvider
}

func (m *BaseMiscellaneous) GetEmojis() []string {
	return constants.Emojis
}
func (m *BaseMiscellaneous) GetLanguageCodes() []string {
	return constants.LanguageCodes
}

func (m *BaseMiscellaneous) CountryCodes() []string {
	return constants.CountryCodes
}
func (m *BaseMiscellaneous) GetCountryISOAlpha3() []string {
	return constants.CountryISOAlpha3
}

func (m *BaseMiscellaneous) GetLocaleData() []string {
	return constants.LocaleData
}

func (m *BaseMiscellaneous) GetCurrencyCodes() []string {
	return constants.CurrencyCode
}

// Boolean generates a random boolean value.
//
// Parameters:
// - chanceOfGettingTrue (optional): Specify the probability of getting 'true' as a percentage (default is 50%).
//
// Returns:
// - A boolean value, either true or false.
func (m *BaseMiscellaneous) Boolean(chanceOfGettingTrue ...int) (bool, error) {
	defaultChanceOfGettingTrue := 50
	if len(chanceOfGettingTrue) > 0 {
		defaultChanceOfGettingTrue = chanceOfGettingTrue[0]
	}
	number, err := m.NumberBetween(1, 100)
	if err != nil {
		return false, err
	}
	return number <= defaultChanceOfGettingTrue, nil
}

// Md5 generates a random MD5 hash.
//
// Returns:
// - A string representing a randomly generated MD5 hash.
func (m *BaseMiscellaneous) Md5() (string, error) {
	number, err := m.NumberBetween()
	if err != nil {
		return "", err
	}
	h := md5.New()
	byteData := m.intToBytes(number)
	hash := h.Sum(byteData)
	return string(hash), nil
}

// Sha1 generates a random SHA-1 hash.
//
// Returns:
// - A string representing a randomly generated SHA-1 hash.
func (m *BaseMiscellaneous) Sha1() (string, error) {
	number, err := m.NumberBetween()
	if err != nil {
		return "", err
	}
	h := sha1.New()
	byteData := m.intToBytes(number)
	hash := h.Sum(byteData)
	return string(hash), nil
}

// Sha256 generates a random SHA-256 hash.
//
// Returns:
// - A string representing a randomly generated SHA-256 hash.
func (m *BaseMiscellaneous) Sha256() (string, error) {
	number, err := m.NumberBetween()
	if err != nil {
		return "", err
	}
	h := sha256.New()
	byteData := m.intToBytes(number)
	hash := h.Sum(byteData)
	return string(hash), nil
}

// Locale generates a random locale code.
//
// Returns:
// - A string representing a randomly generated locale code.
func (m *BaseMiscellaneous) Locale() (string, error) {
	localeData, err := m.RandomElementFromStringSlice(m.GetLocaleData())
	if err != nil {
		return "", err
	}
	return localeData, nil
}

// CountryCode generates a random country code.
//
// Returns:
// - A string representing a randomly generated country code.
func (m *BaseMiscellaneous) CountryCode() (string, error) {
	countryCode, err := m.RandomElementFromStringSlice(m.CountryCodes())
	if err != nil {
		return "", err
	}
	return countryCode, nil
}

// CountryISOAlpha3 generates a random ISO 3166-1 alpha-3 country code.
//
// Returns:
// - A string representing a randomly generated ISO 3166-1 alpha-3 country code.
func (m *BaseMiscellaneous) CountryISOAlpha3() (string, error) {
	countryCode, err := m.RandomElementFromStringSlice(m.GetCountryISOAlpha3())
	if err != nil {
		return "", err
	}
	return countryCode, nil
}

// LanguageCode generates a random language code.
//
// Returns:
// - A string representing a randomly generated language code.
func (m *BaseMiscellaneous) LanguageCode() (string, error) {
	langCode, err := m.RandomElementFromStringSlice(m.GetLanguageCodes())
	if err != nil {
		return "", err
	}
	return langCode, nil
}

// CurrencyCode generates a random currency code.
//
// Returns:
// - A string representing a randomly generated currency code.
func (m *BaseMiscellaneous) CurrencyCode() (string, error) {
	currencyCode, err := m.RandomElementFromStringSlice(m.GetCurrencyCodes())
	if err != nil {
		return "", err
	}
	return currencyCode, nil
}

// Emoji generates a random emoji.
//
// Returns:
// - A string representing a randomly generated emoji.
func (m *BaseMiscellaneous) Emoji() (string, error) {
	emoji, err := m.RandomElementFromStringSlice(m.GetEmojis())
	if err != nil {
		return "", err
	}
	return emoji, nil
}

func (m *BaseMiscellaneous) intToBytes(n int) []byte {
	// Convert the integer to a 4-byte representation (assuming int is 32 bits).
	byteSlice := make([]byte, 4)
	binary.LittleEndian.PutUint32(byteSlice, uint32(n))
	return byteSlice
}
