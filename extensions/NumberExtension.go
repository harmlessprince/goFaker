package extensions

type NumberExtension interface {
	NumberBetween(min int, max int) int
	RandomDigit() int
	RandomDigitNot(except int) int
	RandomDigitNotZero() int
	RandomFloat(numberOfDecimalPlaces int, params ...float64) float64
	RandomNumber(numberOfDigits int, strict bool) int
}
