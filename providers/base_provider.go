package providers

import (
	"errors"
	"fmt"
	"github.com/harmlessprince/goFaker/custom_errors"
	"github.com/harmlessprince/goFaker/helpers"
	"github.com/samber/lo"
	"math"
	"math/rand"
	"reflect"
	"regexp"
	"strings"
	"time"
)

var helperInstance helpers.Helper

type BaseProvider struct {
}

// IntCallbackFunc for int
type IntCallbackFunc func() (int, error)

// StringCallbackFunc for string
type StringCallbackFunc func() (string, error)

// RandomFloatOptions struct to hold optional parameters with default values
type RandomFloatOptions struct {
	NumberOfMaxDecimals uint
	Min                 int
	Max                 int
}

// DefaultRandomFloatOptions returns default options for RandomFloat
func DefaultRandomFloatOptions() RandomFloatOptions {
	return RandomFloatOptions{
		NumberOfMaxDecimals: 0,
		Min:                 0,
		Max:                 0,
	}
}

func (b BaseProvider) RandomDigit() (uint, error) {

	randomDigit, err := helperInstance.RandomNumberBetween(0, 9)
	if err != nil {
		return 0, err
	}
	return uint(randomDigit), err
}

func (b BaseProvider) RandomDigitNotZero() (int, error) {
	randomDigit, err := helperInstance.RandomNumberBetween(1, 9)
	if err != nil {
		return 0, err
	}
	return randomDigit, err
}

func (b BaseProvider) RandomDigitNot(except uint) (int, error) {
	randomDigit, err := helperInstance.RandomNumberBetween(0, 8)
	if err != nil {
		return 0, err
	}

	if uint(randomDigit) >= except {
		randomDigit = randomDigit + 1
	}

	return randomDigit, err
}

func (b BaseProvider) RandomNumber(numberOfDigits uint, strict bool) (int, error) {
	if numberOfDigits <= 0 {
		number, err := b.RandomDigitNotZero()
		if err != nil {
			return 0, err
		}
		return number, nil
	}
	max := int(math.Pow(10, float64(numberOfDigits))) - 1
	if max > helperInstance.LargestRandomNumber() {
		return 0, &custom_errors.InvalidArgumentError{Message: "RandomNumber() can only generate numbers up to math.MaxInt"}
	}
	if strict {
		number, err := helperInstance.RandomNumberBetween(int(math.Pow(10, float64(numberOfDigits-1))), max)
		if err != nil {
			return 0, err
		}
		return number, nil
	}
	number, err := helperInstance.RandomNumberBetween(0, max)
	if err != nil {
		return 0, err
	}
	return number, nil
}

func (b BaseProvider) RandomFloat(options ...RandomFloatOptions) (float64, error) {
	opts := DefaultRandomFloatOptions()
	if len(options) > 0 {
		opts = options[0]
	}

	if opts.NumberOfMaxDecimals <= 0 {
		digit, err := b.RandomDigit()
		if err != nil {
			return 0, err
		}
		return float64(digit), nil
	}

	if opts.Max <= 0 {
		max, err := b.RandomNumber(0, false)
		if err != nil {
			return 0, err
		}
		if opts.Min > max {
			max = opts.Min
		}
	}

	if opts.Min > opts.Max {
		tmp := opts.Min
		opts.Min = opts.Max
		opts.Max = tmp
	}
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randomFloat := float64(opts.Min) + (rand.Float64() * float64(opts.Max-opts.Min))
	return helperInstance.RoundFloat(randomFloat, opts.NumberOfMaxDecimals), nil
}

// NumberBetween generates a random integer between the specified minimum and maximum values (inclusive).
// If no minimum and maximum values are provided, it defaults to generating a random integer
// between 0 and the maximum value representable by an int (2147483647).
//
// Parameters:
//   - params: An optional variadic parameter list where the first parameter represents the minimum value,
//     and the second parameter represents the maximum value.
//
// Returns:
// - An integer representing the randomly generated value between the specified range.
// - An error if there was an issue generating the random number.
func (b BaseProvider) NumberBetween(params ...int) (int, error) {
	min := 0
	max := 2147483647
	if len(params) > 0 {
		min = params[0]
	}
	if len(params) > 1 {
		max = params[1]
	}
	if max < min {
		tmp := min
		min = max
		max = tmp
	}
	value, err := helperInstance.RandomNumberBetween(min, max)
	if err != nil {
		return 0, err
	}
	return value, nil
}

// RandomLetter generates a random lowercase English alphabet letter.
// It uses a random number generator to generate a random integer between
// the ASCII values of 'a' (97) and 'z' (122) and then converts it to a string
// representing the corresponding alphabet character.
//
// Returns:
// - A string containing a random lowercase English alphabet letter.
// - An error if there was an issue generating the random number.
func (b BaseProvider) RandomLetter() (string, error) {
	randomNumber, err := helperInstance.RandomNumberBetween(97, 122)
	if err != nil {
		return "", err
	}
	character := fmt.Sprintf("%c", rune(randomNumber))
	return character, nil
}

// RandomAscii generates a random ASCII character within the range of ASCII values from 33 to 126 (inclusive).
// It uses a random number generator to select a random integer within this range and then converts it
// to a string representing the corresponding ASCII character.
//
// Returns:
// - A string containing a random ASCII character.
// - An error if there was an issue generating the random number.
func (b BaseProvider) RandomAscii() (string, error) {
	randomNumber, err := helperInstance.RandomNumberBetween(33, 126)
	if err != nil {
		return "", err
	}
	character := fmt.Sprintf("%c", rune(randomNumber))
	return character, nil
}
func (b BaseProvider) RandomElementsForSliceStrings(slice []string, count int, allowDuplicates bool) ([]string, error) {

	if len(slice) <= 0 {
		slice = []string{"a", "b", "c"}
	}
	numberOfElements := len(slice)
	randomElements := make([]string, 0)
	if allowDuplicates == false && count > 0 && numberOfElements < count {
		return randomElements, errors.New(fmt.Sprintf("Cannot get %d elements, only %d in array", count, numberOfElements))
	}

	var keys []int
	for key := range slice {
		keys = append(keys, key)
	}
	elementHasBeenSelectedAlready := make(map[int]bool)
	numberOfRandomElements := 0

	return getRandomElements(numberOfRandomElements, randomElements, numberOfElements, allowDuplicates, elementHasBeenSelectedAlready, keys, count, slice)
}
func getRandomElements(
	numberOfRandomElements int,
	randomElements []string, numberOfElements int,
	allowDuplicates bool,
	elementHasBeenSelectedAlready map[int]bool,
	keys []int,
	count int,
	slice any,
) ([]string, error) {

	for numberOfRandomElements <= count {
		index, err := helperInstance.RandomNumberBetween(0, numberOfElements)
		if err != nil {
			return randomElements, err
		}
		if allowDuplicates == false {
			if elementHasBeenSelectedAlready[index] {
				continue
			}
			elementHasBeenSelectedAlready[index] = true
			numberOfRandomElements++
		}
		key := keys[index]
		randomElements = append(randomElements, slice.([]string)[key])
		numberOfRandomElements += 1
	}
	return randomElements, nil
}
func (b BaseProvider) RandomElements(slice interface{}, count int, allowDuplicates bool) (interface{}, error) {
	elements, ok := slice.([]interface{})

	if !ok {
		return slice, errors.New("argument for parameter slice needs to be slice")
	}
	if len(elements) <= 0 {
		elements = []interface{}{"a", "b", "c"}
	}
	numberOfElements := len(elements)

	if allowDuplicates == false && count > 0 && numberOfElements < count {
		return 0, errors.New(fmt.Sprintf("Cannot get %d elements, only %d in array", count, numberOfElements))
	}
	var keys []int
	for key := range elements {
		keys = append(keys, key)
	}

	elementHasBeenSelectedAlready := make(map[int]bool)
	numberOfRandomElements := 0
	randomElements := make([]interface{}, 0)
	for numberOfRandomElements <= count {
		index, err := helperInstance.RandomNumberBetween(0, numberOfElements)
		if err != nil {
			return 0, err
		}
		if allowDuplicates == false {
			if elementHasBeenSelectedAlready[index] {
				continue
			}
			elementHasBeenSelectedAlready[index] = true
			numberOfRandomElements++
		}
		key := keys[index]
		randomElements = append(randomElements, elements[key])
		numberOfRandomElements += 1
	}
	return randomElements, nil
}

func (b BaseProvider) RandomElementFromSlice(slice interface{}) (interface{}, error) {
	elements, ok := slice.([]interface{})

	if !ok {
		return slice, errors.New("argument for parameter slice needs to be slice")
	}
	if len(elements) <= 0 {
		elements = []interface{}{"a", "b", "c"}
	}
	randomElements, err := b.RandomElements(elements, 1, false)
	if err != nil {
		return nil, err
	}
	return randomElements.([]interface{})[0], nil
}
func (b BaseProvider) RandomElementFromStringSlice(slice []string) (string, error) {
	elements := slice
	if len(elements) <= 0 {
		elements = []string{"a", "b", "c"}
	}
	randomElements, err := b.RandomElementsForSliceStrings(elements, 1, false)
	if err != nil {
		return "", err
	}
	return randomElements[0], nil
}
func (b BaseProvider) RandomElementFromMap(elements map[interface{}]interface{}) (interface{}, error) {

	if len(elements) <= 0 {
		elements[1] = "a"
		elements[2] = "b"
		elements[3] = "c"
	}
	randomElements, err := b.RandomElements(elements, 1, false)
	if err != nil {
		return nil, err
	}
	return randomElements.([]interface{})[0], nil
}

// RandomKeyFromSlice returns a random key (index) from a slice of elements.
//
// Parameters:
//   - slice: A slice of interface{} containing elements.
//
// Returns:
//   - int: A random key (index) from the slice.
//   - error: An error if the provided argument is not a slice or if the slice is empty.
//
// Example usage:
//
//	mySlice := []interface{}{"apple", "banana", "cherry", "date"}
//	randomKey, err := BaseProvider{}.RandomKeyFromSlice(mySlice)
//	if err != nil {
//	    fmt.Println("Error:", err)
//	} else {
//	    fmt.Println("Random key:", randomKey)
//	}
func (b BaseProvider) RandomKeyFromSlice(slice interface{}) (int, error) {
	elements, ok := slice.([]interface{})

	if !ok {
		return 0, errors.New("argument for parameter slice needs to be slice")
	}

	if len(elements) <= 0 {
		return 0, errors.New("length of slice must be greater than 0")
	}
	var keys []int
	for key := range elements {
		keys = append(keys, key)
	}

	randomNumber, err := helperInstance.RandomNumberBetween(0, len(elements))
	if err != nil {
		return 0, err
	}

	return keys[randomNumber], nil
}

// RandomKeyFromMap returns a random key from a map of elements.
//
// Parameters:
//   - elements: A map of interface{} to interface{} containing elements.
//
// Returns:
//   - interface{}: A random key from the map.
//   - error: An error if the provided map is empty.
//
// Example usage:
//
//	myMap := map[interface{}]interface{}{1: "apple", 2: "banana", 3: "cherry"}
//	randomKey, err := BaseProvider{}.RandomKeyFromMap(myMap)
//	if err != nil {
//	    fmt.Println("Error:", err)
//	} else {
//	    fmt.Println("Random key:", randomKey)
//	}
func (b BaseProvider) RandomKeyFromMap(elements map[interface{}]interface{}) (interface{}, error) {
	if len(elements) <= 0 {
		elements[1] = "a"
		elements[2] = "b"
		elements[3] = "c"
	}
	keys := lo.Keys[interface{}, interface{}](elements)
	randomNumber, err := helperInstance.RandomNumberBetween(0, len(elements))
	if err != nil {
		return 0, err
	}
	return keys[randomNumber], nil
}

// ShuffleSlice shuffles the elements of a slice in a random order.
//
// Parameters:
//   - elements: A slice of interface{} containing the elements to shuffle.
//
// Returns:
//   - []interface{}: A new slice with the elements shuffled in a random order.
//
// Example usage:
//
//	originalSlice := []interface{}{1, 2, 3, 4, 5}
//	shuffledSlice := BaseProvider{}.ShuffleSlice(originalSlice)
//	fmt.Println("Shuffled slice:", shuffledSlice)
func (b BaseProvider) ShuffleSlice(elements []interface{}) []interface{} {
	return lo.Shuffle(elements)
}

// ShuffleString shuffles the characters of a string in a random order.
//
// Parameters:
//   - elements: A string containing the characters to shuffle.
//
// Returns:
//   - string: A new string with the characters shuffled in a random order.
//
// Example usage:
//
//	originalString := "hello"
//	shuffledString := BaseProvider{}.ShuffleString(originalString)
//	fmt.Println("Shuffled string:", shuffledString)
func (b BaseProvider) ShuffleString(elements string) string {
	splitString := strings.Split(elements, "")
	shuffledSplitString := lo.Shuffle(splitString)
	return strings.Join(shuffledSplitString[:], "")
}
func (b BaseProvider) replaceWildcardInt(input string, wildcard string, callback IntCallbackFunc) (string, error) {

	pos := strings.Index(input, wildcard)

	if pos == -1 {
		return input, nil
	}

	last := strings.LastIndex(input, wildcard) + 1

	var builder strings.Builder
	builder.WriteString(input[:pos])

	for i := pos; i < last; i++ {
		builder.WriteByte(input[i])
		if string(input[i]) == wildcard {
			result, err := callback()
			if err != nil {
				return "", err
			}
			builder.WriteString(fmt.Sprintf("%d", result))
		}
	}

	builder.WriteString(input[last:])
	return builder.String(), nil
}

func (b BaseProvider) replaceWildcardString(input string, wildcard string, callback StringCallbackFunc) (string, error) {
	pos := strings.Index(input, wildcard)

	if pos == -1 {
		return input, nil
	}

	last := strings.LastIndex(input, wildcard) + 1

	var builder strings.Builder
	builder.WriteString(input[:pos])

	for i := pos; i < last; i++ {
		builder.WriteByte(input[i])
		if string(input[i]) == wildcard {
			result, err := callback()
			if err != nil {
				return "", err
			}
			builder.WriteString(result)
		}
	}

	builder.WriteString(input[last:])
	return builder.String(), nil
}

// Numerify Replaces all hash sign ("#") occurrences with a random number
// Numerify Replaces all percentage sign ("%") occurrences with a not null number
//
// Parameters:
//   - item:  String that needs to bet parsed
//
// Returns:
//   - string: A new string with both # and % replaced with random numbers.
//
// Example usage:
//
//	originalString := "hello##3"
//	result := BaseProvider{}.Numerify(originalString)
//	fmt.Println("result string:", result)
func (b BaseProvider) Numerify(item string) (string, error) {
	var toReplace []int
	if len(item) <= 0 {
		item = "###"
	}
	position := strings.Index(item, "#")

	if position != -1 {
		// Find the last occurrence of '#' from the given position
		last := strings.LastIndex(item, "#") + 1
		// Slice the string to search within the specified range
		//substring := item[position:last]
		for i := position; i < last; i++ {

			if item[i] == '#' {
				toReplace = append(toReplace, i)
			}
		}
		numberOfReplacements := len(toReplace)

		// Determine the maximum number of random digits to generate at once
		maxAtOnce := len(fmt.Sprint(helperInstance.LargestRandomNumber())) - 1

		//var numbers string
		var numbers strings.Builder
		i := 0
		for i < numberOfReplacements {

			size := math.Min(float64(numberOfReplacements-1), float64(maxAtOnce))

			randomNumber, err := b.RandomNumber(uint(size), false)
			if err != nil {
				return "", err
			}
			randomNumberStr := fmt.Sprintf("%d", randomNumber)

			numbers.WriteString(randomNumberStr)
			if size > 0 {
				i += int(size)
			} else {
				i += 1
			}

		}
		//Convert the 'numbers' string to a slice of runes for replacement
		numberRunes := []rune(numbers.String())

		// Replace characters in the original string based on positions in 'toReplace'
		for i := 0; i < numberOfReplacements; i++ {
			item = helperInstance.ReplaceAtIndex(item, toReplace[i], string(numberRunes[i]))
		}
	}

	parsedItem, err := b.replaceWildcardInt(item, "%", b.RandomDigitNotZero)
	if err != nil {
		return "", err
	}

	return parsedItem, nil
}

// Lexify Replaces all question mark ('?') occurrences with a random letter
//
// Parameters:
//   - item:  String that needs to be parsed
//
// Returns:
//   - string: A new string with all ? replaced with random letters.
//
// Example usage:
//
//	originalString := "hello ???"
//	result := BaseProvider{}.Lexify(originalString)
//	fmt.Println("result string:", result)
func (b BaseProvider) Lexify(item ...string) (string, error) {
	itemToBeParsed := "???"
	if len(item) > 0 {
		itemToBeParsed = item[0]
	}
	parsedString, err := b.replaceWildcardString(itemToBeParsed, "?", b.RandomLetter)
	if err != nil {
		return "", err
	}
	return parsedString, nil
}

// Bothify Replaces hash signs ('#') and question marks ('?') with random numbers and letters respectively
// Bothify Replaces An asterisk ('*') is replaced with either a random number or a random letter
// Parameters:
//   - item:  String that needs to be parsed
//
// Returns:
//   - string: A new string with ?, # and * replaced with respective characters.
//
// Example usage:
//
//	originalString := "hello ?#*"
//	result := BaseProvider{}.Bothify(originalString)
//	fmt.Println("result string:", result)
func (b BaseProvider) Bothify(item ...string) (string, error) {
	itemToBeParsed := "## ?? **"
	if len(item) > 0 {
		itemToBeParsed = item[0]
	}
	parsedString, err := b.replaceWildcardString(itemToBeParsed, "*", func() (string, error) {
		number, err := helperInstance.RandomNumberBetween(0, 2)
		if err != nil {
			return "", err
		}
		if number == 1 {
			return "#", nil
		} else {
			return "?", nil
		}
	})

	if err != nil {
		return "", err
	}

	parsedString, err = b.Numerify(parsedString)
	if err != nil {
		return "", err
	}

	parsedString, err = b.Lexify(parsedString)
	if err != nil {
		return "", err
	}

	return parsedString, nil
}

// Asciify Replaces * signs with random numbers and letters and special characters
//
// Parameters:
//   - item:  String that needs to be parsed
//
// Returns:
//   - string: A new string with * replaced with respective characters.
//
// Example usage:
//
//	originalString := "****"
//	result, _ := BaseProvider{}.Asciify(originalString)
//	fmt.Println("result string:", result)
func (b BaseProvider) Asciify(item ...string) (string, error) {
	itemToBeParsed := "****"
	if len(item) > 0 {
		itemToBeParsed = item[0]
	}
	re := regexp.MustCompile(`\*`)
	return helperInstance.PregReplaceCallback(re, itemToBeParsed, b.RandomAscii), nil
}

func (b BaseProvider) ToLower(item string) string {
	return strings.ToLower(item)
}

func (b BaseProvider) ToUpper(item string) string {
	return strings.ToUpper(item)
}

// Parse takes an input string, extracts placeholders enclosed in `{{` and `}}`, and replaces them with values obtained by calling methods on the provided data structure.
// It returns the updated string and an error if any issues occur during parsing.
func (b BaseProvider) Parse(input string, data interface{}, args ...map[string]interface{}) (string, error) {
	parsed := input
	structValue := reflect.ValueOf(data)
	methodMap := make(map[string]reflect.Value)

	// Retrieve the type of the provided data structure and create a map of method names to their associated reflect.Value.
	structType := reflect.TypeOf(data)
	for i := 0; i < structType.NumMethod(); i++ {
		method := structType.Method(i)
		methodMap[method.Name] = structValue.MethodByName(method.Name)
	}
	//fmt.Println("structValue: ", structValue)
	//fmt.Println("structType: ", structType)
	//fmt.Println("kind", structValue.Kind())
	//fmt.Println("NMethod: ", structType.NumMethod())
	for {
		start := strings.Index(parsed, "{{")
		end := strings.Index(parsed, "}}")

		if start == -1 || end == -1 {
			break
		}
		// Extract the placeholder name enclosed in `{{` and `}}`.
		fieldName := parsed[start+2 : end]

		field := structValue.MethodByName(fieldName)
		// Check if the method does not exist and is  not a function.
		if field.IsValid() == false && field.Type().Kind() != reflect.Func {
			return "", fmt.Errorf("method '%s' does not exist on the struct", fieldName)
		}
		methodArgs := make([]reflect.Value, 0)
		// Check if an argument exists for the current method, and add it if found.
		if len(args) > 0 {
			if arg, argExists := args[0][fieldName]; argExists {
				methodArgs = append(methodArgs, reflect.ValueOf(arg))
			}
		}
		var result []reflect.Value
		if len(methodArgs) > 0 {
			result = field.Call(methodArgs)
		} else {
			result = field.Call(nil)
		}

		if len(result) > 0 {
			parsed = strings.Replace(parsed, "{{"+fieldName+"}}", result[0].Interface().(string), 1)
		}

	}

	return parsed, nil
}
