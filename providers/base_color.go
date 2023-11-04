package providers

import (
	"errors"
	"fmt"
	"github.com/harmlessprince/goFaker/constants"
	"strconv"
	"strings"
)

type BaseColor struct {
	BaseProvider
}
type ColorInterface interface {
	HexColor() (string, error)
	SafeHexColor() (string, error)
	RgbColorAsArray() ([]string, error)
	RgbColor() (string, error)
	RgbCssColor() (string, error)
	RgbaCssColor() (string, error)
	SafeColorName() (string, error)
	ColorName() (string, error)
	HslColor() (string, error)
	HslColorAsArray() ([]int, error)
}

// HexColor generates a random hexadecimal color code.
//
// Returns:
// - A string representing a randomly generated hexadecimal color code.
func (b *BaseColor) HexColor() (string, error) {
	numberBtw, errNumberBtw := b.NumberBetween(1, 16777215)
	if errNumberBtw != nil {
		return "", errors.New(errNumberBtw.Error())
	}
	numberBtwToHex := helperInstance.DecimalToHexDecimal(int64(numberBtw))
	paddedNumberBtwToHex := helperInstance.StrPadLeft(numberBtwToHex, 6, "0")
	return "#" + paddedNumberBtwToHex, nil
}

// SafeHexColor generates a random safe hexadecimal color code.
//
// Returns:
// - A string representing a randomly generated safe hexadecimal color code.
func (b *BaseColor) SafeHexColor() (string, error) {
	numberBtw, errNumberBtw := b.NumberBetween(0, 255)
	if errNumberBtw != nil {
		return "", errors.New(errNumberBtw.Error())
	}
	numberBtwToHex := helperInstance.DecimalToHexDecimal(int64(numberBtw))
	color := helperInstance.StrPadLeft(numberBtwToHex, 3, "0")
	split := strings.Split(color, "")
	return "#" + split[0] + split[0] + split[1] + split[1] + split[2] + split[2], nil
}

// RgbColorAsArray generates a random RGB color represented as an array of strings.
//
// Returns:
// - An array of strings representing a randomly generated RGB color.
func (b *BaseColor) RgbColorAsArray() ([]string, error) {
	color, err := b.HexColor()
	if err != nil {
		return nil, err
	}
	return []string{
		strconv.FormatInt(helperInstance.HexDecimalToDecimal(helperInstance.Substr(color, 1, 2)), 10),
		strconv.FormatInt(helperInstance.HexDecimalToDecimal(helperInstance.Substr(color, 3, 2)), 10),
		strconv.FormatInt(helperInstance.HexDecimalToDecimal(helperInstance.Substr(color, 5, 2)), 10),
	}, nil
}

// RgbColor generates a random RGB color code.
//
// Returns:
// - A string representing a randomly generated RGB color code.
func (b *BaseColor) RgbColor() (string, error) {
	rgbColorAsArray, err := b.RgbColorAsArray()
	if err != nil {
		return "", err
	}
	return strings.Join(rgbColorAsArray, ","), nil
}

// RgbCssColor generates a random CSS-compatible RGB color code.
//
// Returns:
// - A string representing a randomly generated CSS-compatible RGB color code.
func (b *BaseColor) RgbCssColor() (string, error) {
	rgbColor, err := b.RgbColor()
	if err != nil {
		return "", err
	}
	return "rgb(" + rgbColor + ")", nil
}

// RgbaCssColor generates a random CSS-compatible RGBA color code.
//
// Returns:
// - A string representing a randomly generated CSS-compatible RGBA color code.
func (b *BaseColor) RgbaCssColor() (string, error) {
	floatOptions := RandomFloatOptions{
		NumberOfMaxDecimals: 1,
		Min:                 0,
		Max:                 1,
	}
	randomFloat, err := b.RandomFloat(floatOptions)
	if err != nil {
		return "", err
	}
	randomFloatToString := strconv.FormatFloat(randomFloat, 'f', -1, 64)
	rgbColor, err := b.RgbColor()
	if err != nil {
		return "", err
	}
	return "rgb(" + rgbColor + "," + randomFloatToString + ")", nil
}

// SafeColorName generates a random safe color name.
//
// Returns:
// - A string representing a randomly generated safe color name.
func (b *BaseColor) SafeColorName() (string, error) {
	color, err := b.RandomElementFromStringSlice(constants.SafeColorNames)
	if err != nil {
		return "", err
	}
	return color, nil
}

// ColorName generates a random color name.
//
// Returns:
// - A string representing a randomly generated color name.
func (b *BaseColor) ColorName() (string, error) {
	color, err := b.RandomElementFromStringSlice(constants.AllColorNames)
	if err != nil {
		return "", err
	}
	return color, nil
}

// HslColor generates a random HSL color code.
//
// Returns:
// - A string representing a randomly generated HSL color code.
func (b *BaseColor) HslColor() (string, error) {
	first, err := b.NumberBetween(0, 360)
	if err != nil {
		return "", err
	}
	second, err := b.NumberBetween(0, 100)
	if err != nil {
		return "", err
	}
	third, err := b.NumberBetween(0, 100)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d,%d,%d", first, second, third), nil
}

// HslColorAsArray generates a random HSL color represented as an array of integers.
//
// Returns:
// - An array of integers representing a randomly generated HSL color.
func (b *BaseColor) HslColorAsArray() ([]int, error) {
	first, err := b.NumberBetween(0, 360)
	if err != nil {
		return nil, err
	}
	second, err := b.NumberBetween(0, 100)
	if err != nil {
		return nil, err
	}
	third, err := b.NumberBetween(0, 100)
	if err != nil {
		return nil, err
	}
	return []int{
		first, second, third,
	}, nil
}
