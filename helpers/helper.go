package helpers

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"net"
	"os"
	"regexp"
	"strconv"
	"time"
)

var Seed = time.Now().UnixNano()

type Helper struct {
}

func (h Helper) RoundFloat(value float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	result := math.Round(value*ratio) / ratio
	return result
}

// RandomElementsKeys selects a random set of unique indices from a given slice of elements
// and returns those indices as a slice of integers. It ensures that the selected indices
// are unique and within the bounds of the input slice.
//
// Parameters:
// - slice: An interface{} representing the input slice of elements.
// - size: The number of unique indices to be selected from the input slice.
//
// Returns:
//   - A slice of integers representing the selected unique indices.
//   - An error if the input slice is not valid, has zero length, or if size is greater than
//     the length of the input slice.
func (h Helper) RandomElementsKeys(slice interface{}, size int, allowDuplicates bool) (interface{}, error) {
	// Type assertion to convert the interface{} slice to a []interface{}
	elements, ok := slice.([]interface{})

	if !ok {
		return slice, errors.New("argument for parameter slice needs to be slice")
	}
	if len(elements) <= 0 {
		return nil, errors.New("slice must be greater than zero")
	}

	if len(elements) < size {
		return slice, errors.New("size is greater than the length of the slice")
	}

	rand.New(rand.NewSource(Seed))

	selectedKeys := make([]int, 0, size)

	if allowDuplicates == false {
		selectedKeyMap := make(map[int]struct{})
		for len(selectedKeyMap) < size {
			randomIndex := rand.Intn(len(elements))
			_, exists := selectedKeyMap[randomIndex]
			if !exists {
				selectedKeyMap[randomIndex] = struct{}{}
				selectedKeys = append(selectedKeys, randomIndex)

			}
		}
	} else {
		for index := 0; index < size; index++ {
			randomIndex := rand.Intn(len(elements))
			selectedKeys = append(selectedKeys, randomIndex)
		}
	}
	return selectedKeys, nil
}
func (h Helper) RandomElement(slice interface{}) (interface{}, error) {
	elements, ok := slice.([]interface{})

	if !ok {
		return slice, errors.New("invalid slice type")
	}
	if len(elements) <= 0 {
		return nil, errors.New("slice must be greater than zero")
	}

	randomKey, err := h.RandomElementsKeys(elements, 1, false)
	if err != nil {
		return nil, err
	}

	return elements[randomKey.([]int)[0]], err
}
func (h Helper) RandomNumber() int {
	rand.New(rand.NewSource(Seed))
	return rand.Int()
}

// RandomNumberBetween generates a random integer between the specified minimum and maximum values (inclusive).
// It returns the generated random integer and an error if the minimum value is greater than or equal to the maximum value.
//
// Parameters:
//   - min: The minimum value (inclusive) for the random number.
//   - max: The maximum value (inclusive) for the random number.
//
// Returns:
//   - int: The generated random integer.
//   - error: An error if min is greater than or equal to max, otherwise nil.
//
// Example usage:
//
//	randInt, err := Helper{}.RandomNumberBetween(1, 100)
//	if err != nil {
//	    fmt.Println("Error:", err)
//	} else {
//	    fmt.Println("Random number:", randInt)
//	}
func (h Helper) RandomNumberBetween(min int, max int) (int, error) {
	if min >= max {
		return 0, errors.New("min value must be greater than max value")
	}
	rand.New(rand.NewSource(Seed))

	return rand.Intn(max-min) + min, nil
}

func (h Helper) LargestRandomNumber() int {
	return math.MaxInt
}

func (h Helper) SmallestRandomNumber() int {
	return math.MinInt
}

func (h Helper) FormatFloatToDecimalPlaces(myFloat float64, decimalPlaces int) (float64, error) {
	formatString := fmt.Sprintf("%%.%df", decimalPlaces)
	formattedValue := fmt.Sprintf(formatString, myFloat)
	var result float64
	_, err := fmt.Sscanf(formattedValue, "%f", &result)
	if err != nil {
		// Handle the error if necessary
		return 0.0, err // Return a default value on error
	}

	return result, nil
}

//
//func (h Helper) mapOrSliceValues() {
//
//}

// ReverseString reverses a given string.
func (h Helper) ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// ReplaceAtIndex replaces a character in a string at a specific index.
func (h Helper) ReplaceAtIndex(input string, index int, replacement string) string {
	return input[:index] + replacement + input[index+len(replacement):]
}

func (h Helper) PregReplaceCallback(re *regexp.Regexp, input string, callback func() (string, error)) string {
	return re.ReplaceAllStringFunc(input, func(match string) string {
		matched, _ := callback()
		return matched
	})
}
func (h Helper) PregReplaceCallbackNoError(re *regexp.Regexp, input string, callback func(item ...string) string) string {
	return re.ReplaceAllStringFunc(input, func(match string) string {

		matched := callback()
		fmt.Println(match)
		return matched
	})
}

func (h Helper) ArraySum(digits []int) int {
	sum := 0
	for _, num := range digits {
		sum += num
	}
	return sum
}

func (h Helper) SliceOfStringsToInt(numberString []string) []int {
	digits := make([]int, 0, len(numberString))
	for _, char := range numberString {
		digit, err := strconv.Atoi(char)
		if err != nil {
			panic(fmt.Sprintf("Error converting character to int: %v\n", err))
		}
		digits = append(digits, digit)
	}
	return digits
}

func (h Helper) StringToInt(numberString string) []int {
	digits := make([]int, 0, len(numberString))
	for _, char := range numberString {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			panic(fmt.Sprintf("Error converting character to int: %v\n", err))
		}
		digits = append(digits, digit)
	}
	return digits
}
func (h Helper) DecimalToHexDecimal(number int64) string {
	return strconv.FormatInt(number, 16)
}
func (h Helper) HexDecimalToDecimal(str string) int64 {
	value, err := strconv.ParseInt(str, 16, 0)
	if err != nil {
		panic(err.Error())
		return 0
	}
	return value
}

func (h Helper) Ip2long(ip string) (uint32, error) {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return 0, errors.New("invalid IP address")
	}
	ipBytes := parsedIP.To4()
	if ipBytes == nil {
		return 0, errors.New("IPv6 address not supported")
	}
	return binary.BigEndian.Uint32(ipBytes), nil
}

func (h Helper) Long2ip(ipInt uint32) (string, error) {
	ipBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(ipBytes, ipInt)
	ip := net.IP(ipBytes)
	ip = ip.To4()
	if ip == nil {
		return "", errors.New("invalid IP integer")
	}
	return ip.String(), nil
}

func (h Helper) StrPadLeft(input string, padLength int, padString string) string {
	output := ""
	inputLen := len(input)
	if inputLen >= padLength {
		return input
	}
	padCount := padLength - inputLen
	for i := 1; i <= padCount; i = i + len(padString) {
		output += padString
	}
	return output + input
}
func (h Helper) StrPadRight(input string, padLength int, padString string) string {
	output := input
	inputLen := len(input)

	if inputLen >= padLength {
		return input
	}

	padCount := padLength - inputLen
	padStringLen := len(padString)

	for i := 0; i < padCount; i += padStringLen {
		output += padString
	}

	// Trim any excess padding to ensure the result is exactly padLength long
	if len(output) > padLength {
		output = output[:padLength]
	}

	return output
}

// Substr returns the portion of string specified by the start and length parameters.
// params
//   - string $string — The input string.
//   - int $offset
//   - int $length
func (h Helper) Substr(str string, start int, length int) string {
	return str[start : start+length]
}

func (h Helper) IsWritable(path string) (bool, error) {
	// Try to create and delete a temporary file in the directory
	testFilePath := path + string(os.PathSeparator) + "test.tmp"
	file, err := os.Create(testFilePath)
	if err != nil {
		return false, err
	}
	err = file.Close()
	if err != nil {
		return false, err
	}
	err = os.Remove(testFilePath)
	if err != nil {
		return false, err
	}
	return true, nil
}
func (h Helper) IsDirectory(path string) (bool, error) {
	// Check if the path exists and is a directory
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	if !fileInfo.IsDir() {
		return false, fmt.Errorf("%s is not a directory", path)
	}
	return true, nil
}
