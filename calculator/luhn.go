package calculator

import (
	"errors"
	"fmt"
	"github.com/harmlessprince/goFaker/helpers"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Luhn struct {
}

func (l *Luhn) checkSum(number string) float64 {
	numberInDigits := []rune(number)
	length := len(number)
	sum := 0
	fmt.Println(number)
	for i := length - 1; i >= 0; i -= 2 {
		digit := numberInDigits[i]
		digitAsString := string(digit)
		digitValue, err := strconv.Atoi(digitAsString)
		if err != nil {
			message := fmt.Sprintf("Error converting digit %s to integer: %v\n", digitAsString, err)
			log.Fatal(message)
		}
		sum += digitValue
	}

	for i := length - 2; i >= 0; i -= 2 {
		digit := numberInDigits[i]
		digitAsString := string(digit)
		digitValue, err := strconv.Atoi(digitAsString)
		if err != nil {
			message := fmt.Sprintf("Error converting digit %s to integer: %v\n", digitAsString, err)
			log.Fatal(message)
		}
		doubledDigit := digitValue * 2
		digitAsStringSplit := strings.Split(strconv.Itoa(doubledDigit), "")
		digitsAsInt := helpers.Helper{}.SliceOfStringsToInt(digitAsStringSplit)
		sum += helpers.Helper{}.ArraySum(digitsAsInt)
	}
	return float64(sum % 10)
}

func (l *Luhn) ComputeCheckDigit(partialNumber string) string {
	checkDigit := l.checkSum(partialNumber)
	if checkDigit == 0 {
		return "0"
	}
	return strconv.FormatFloat(10-checkDigit, 'f', -1, 64)
}

func (l *Luhn) IsValid(numberString string) bool {
	return l.checkSum(numberString) == 0
}

// GenerateLuhnNumber generates a Luhn compliant number.
func (l *Luhn) GenerateLuhnNumber(numberString string) (string, error) {
	pattern := "^[0-9]+$"
	matched, err := regexp.MatchString(pattern, numberString)
	if err != nil {
		panic(err.Error())
	}
	if !matched {
		return "", errors.New("the argument passed is not a valid digit")
	}
	return numberString + l.ComputeCheckDigit(numberString), nil
}
