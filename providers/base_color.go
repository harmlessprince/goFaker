package providers

import (
	"fmt"
	"github.com/harmlessprince/goFaker/constants"
	"log"
	"strconv"
	"strings"
)

type BaseColor struct {
	BaseProvider
}

func (b *BaseColor) HexColor() string {
	numberBtw, errNumberBtw := b.NumberBetween(1, 16777215)
	if errNumberBtw != nil {
		log.Fatal(errNumberBtw.Error())
		return ""
	}
	numberBtwToHex := helperInstance.DecimalToHexDecimal(int64(numberBtw))
	paddedNumberBtwToHex := helperInstance.StrPadLeft(numberBtwToHex, 6, "0")
	return "#" + paddedNumberBtwToHex
}

func (b *BaseColor) SafeHexColor() string {
	numberBtw, errNumberBtw := b.NumberBetween(0, 255)
	if errNumberBtw != nil {
		log.Fatal(errNumberBtw.Error())
		return ""
	}
	numberBtwToHex := helperInstance.DecimalToHexDecimal(int64(numberBtw))
	color := helperInstance.StrPadLeft(numberBtwToHex, 3, "0")
	split := strings.Split(color, "")
	return "#" + split[0] + split[0] + split[1] + split[1] + split[2] + split[2]
}

func (b *BaseColor) RgbColorAsArray() []string {
	color := b.HexColor()
	return []string{
		strconv.FormatInt(helperInstance.HexDecimalToDecimal(helperInstance.Substr(color, 1, 2)), 10),
		strconv.FormatInt(helperInstance.HexDecimalToDecimal(helperInstance.Substr(color, 3, 2)), 10),
		strconv.FormatInt(helperInstance.HexDecimalToDecimal(helperInstance.Substr(color, 5, 2)), 10),
	}
}

func (b *BaseColor) RgbColor() string {
	return strings.Join(b.RgbColorAsArray(), ",")
}

func (b *BaseColor) RgbCssColor() string {
	return "rgb(" + b.RgbColor() + ")"
}

func (b *BaseColor) RgbaCssColor() string {
	floatOptions := RandomFloatOptions{
		NumberOfMaxDecimals: 1,
		Min:                 0,
		Max:                 1,
	}
	randomFloat, err := b.RandomFloat(floatOptions)
	if err != nil {
		log.Fatal(err.Error())
		return ""
	}
	randomFloatToString := strconv.FormatFloat(randomFloat, 'f', -1, 64)

	return "rgb(" + b.RgbColor() + "," + randomFloatToString + ")"
}

func (b *BaseColor) SafeColorName() string {
	color, err := b.RandomElementFromStringSlice(constants.SafeColorNames)
	if err != nil {
		log.Fatal(err.Error())
		return ""
	}
	return color
}

func (b *BaseColor) ColorName() string {
	color, err := b.RandomElementFromStringSlice(constants.AllColorNames)
	if err != nil {
		log.Fatal(err.Error())
		return ""
	}
	return color
}

func (b *BaseColor) HslColor() string {
	first, errFirst := b.NumberBetween(0, 360)
	second, errSecond := b.NumberBetween(0, 100)
	third, errThird := b.NumberBetween(0, 100)
	if errFirst != nil {
		log.Fatal(errFirst.Error())
		return ""
	}
	if errSecond != nil {
		log.Fatal(errSecond.Error())
		return ""
	}
	if errThird != nil {
		log.Fatal(errThird.Error())
		return ""
	}
	return fmt.Sprintf("%d,%d,%d", first, second, third)
}

func (b *BaseColor) HslColorAsArray() []int {
	first, errFirst := b.NumberBetween(0, 360)
	second, errSecond := b.NumberBetween(0, 100)
	third, errThird := b.NumberBetween(0, 100)
	if errFirst != nil {
		log.Fatal(errFirst.Error())
		return []int{}
	}
	if errSecond != nil {
		log.Fatal(errSecond.Error())
		return []int{}
	}
	if errThird != nil {
		log.Fatal(errThird.Error())
		return []int{}
	}
	return []int{
		first, second, third,
	}
}
