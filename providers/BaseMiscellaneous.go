package providers

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/binary"
	"github.com/harmlessprince/goFaker/constants"
)

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

func (m *BaseMiscellaneous) Boolean(chanceOfGettingTrue ...int) bool {
	defaultChanceOfGettingTrue := 50
	if len(chanceOfGettingTrue) > 0 {
		defaultChanceOfGettingTrue = chanceOfGettingTrue[0]
	}
	number, err := m.NumberBetween(1, 100)
	if err != nil {
		panic(err.Error())
		return false
	}
	return number <= defaultChanceOfGettingTrue
}

func (m *BaseMiscellaneous) Md5() string {
	number, err := m.NumberBetween()
	if err != nil {
		panic(err.Error())
		return ""
	}
	h := md5.New()
	byteData := m.intToBytes(number)
	hash := h.Sum(byteData)
	return string(hash)
}

func (m *BaseMiscellaneous) Sha1() string {
	number, err := m.NumberBetween()
	if err != nil {
		panic(err.Error())
		return ""
	}
	h := sha1.New()
	byteData := m.intToBytes(number)
	hash := h.Sum(byteData)
	return string(hash)
}

func (m *BaseMiscellaneous) Sha256() string {
	number, err := m.NumberBetween()
	if err != nil {
		panic(err.Error())
		return ""
	}
	h := sha256.New()
	byteData := m.intToBytes(number)
	hash := h.Sum(byteData)
	return string(hash)
}

func (m *BaseMiscellaneous) Locale() string {
	localeData, err := m.RandomElementFromStringSlice(m.GetLocaleData())
	if err != nil {
		panic(err.Error())
		return ""
	}
	return localeData
}

func (m *BaseMiscellaneous) CountryCode() string {
	countryCode, err := m.RandomElementFromStringSlice(m.CountryCodes())
	if err != nil {
		panic(err.Error())
		return ""
	}
	return countryCode
}

func (m *BaseMiscellaneous) CountryISOAlpha3() string {
	countryCode, err := m.RandomElementFromStringSlice(m.GetCountryISOAlpha3())
	if err != nil {
		panic(err.Error())
		return ""
	}
	return countryCode
}

func (m *BaseMiscellaneous) LanguageCode() string {
	langCode, err := m.RandomElementFromStringSlice(m.GetLanguageCodes())
	if err != nil {
		panic(err.Error())
		return ""
	}
	return langCode
}

func (m *BaseMiscellaneous) CurrencyCode() string {
	currencyCode, err := m.RandomElementFromStringSlice(m.GetCurrencyCodes())
	if err != nil {
		panic(err.Error())
		return ""
	}
	return currencyCode
}

func (m *BaseMiscellaneous) Emoji() string {
	emoji, err := m.RandomElementFromStringSlice(m.GetEmojis())
	if err != nil {
		panic(err.Error())
		return ""
	}
	return emoji
}

func (m *BaseMiscellaneous) intToBytes(n int) []byte {
	// Convert the integer to a 4-byte representation (assuming int is 32 bits).
	byteSlice := make([]byte, 4)
	binary.LittleEndian.PutUint32(byteSlice, uint32(n))
	return byteSlice
}
