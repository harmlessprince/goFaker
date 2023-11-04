package calculator

import (
	"fmt"
	"github.com/samber/lo"
	"regexp"
	"strconv"
	"strings"
)

const isbnPattern = `^\d{9}[0-9X]$`

type Isbn struct {
}

func (i *Isbn) CheckSum(digits string) (string, error) {
	if _, err := strconv.ParseInt(digits, 10, 32); err != nil {
		return "", err
	}
	length := 9
	if len(digits) != length {
		return "", fmt.Errorf("input length should be equal to %d", length)
	}
	splitDigits := strings.Split(digits, "")
	mappedDigits := lo.Map(splitDigits, func(element string, index int) int {
		parsedValue, _ := strconv.ParseInt(element, 10, 32)
		return int(int64(int(parsedValue) * (10 - index)))
	})
	sumOfMappedDigits := lo.Sum(mappedDigits)
	result := (11 - (sumOfMappedDigits % 11)) % 11
	// 10 is replaced by X
	if result < 10 {
		return strconv.FormatInt(int64(result), 10), nil
	}
	return "X", nil
}

func (i *Isbn) IsValid(isbn string) (bool, error) {
	re, err := regexp.Compile(isbnPattern)
	if err != nil {
		return false, err
	}
	if re.MatchString(isbn) == false {
		return false, nil
	}
	return true, nil
}
