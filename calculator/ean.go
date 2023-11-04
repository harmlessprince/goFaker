package calculator

import (
	"regexp"
	"strconv"
	"strings"
)

const eanPattern = `^(?:\d{8}|\d{13})$`

type Ean struct {
}

func (e *Ean) CheckSum(digits string) (int, error) {
	if _, err := strconv.ParseInt(digits, 10, 64); err != nil {
		return 0, err
	}
	sequence := []int{1, 3}
	if len(digits)+1 == 8 {
		sequence = []int{3, 1}
	}
	sums := 0
	splitDigits := strings.Split(digits, "")
	for i, digit := range splitDigits {
		digitAsInt, _ := strconv.ParseInt(digit, 10, 64)
		sums += int(digitAsInt) * sequence[i%2]
	}
	return (10 - sums%10) % 10, nil
}

func (e *Ean) IsValid(ean string) (bool, error) {
	re, err := regexp.Compile(eanPattern)
	if err != nil {
		return false, err
	}
	if re.MatchString(ean) == false {
		return false, nil
	}
	return true, nil
}
