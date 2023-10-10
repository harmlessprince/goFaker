package extensions

import "time"

type DateTimeExtension interface {
	// DateTime Get a  Time object between January 1, 1970, and `until` (defaults to "now").
	//
	// Parameters:
	//   - name: description.
	// Returns:
	//
	//	- type: description
	//
	// Example usage:
	//
	DateTime(params ...string) time.Time
	// DateTimeAD Get a  Time object between January 1, 1970, and `until` (defaults to "now").
	//
	// Parameters:
	//   - name: description.
	// Returns:
	//
	//	- type: description
	//
	// Example usage:
	//
	DateTimeAD(params ...string) time.Time

	// DateTimeBetween Get a  Time object between January 1, 1970, and `until` (defaults to "now").
	//
	// Parameters:
	//   - name: description.
	// Returns:
	//
	//	- type: description
	//
	// Example usage:
	//
	DateTimeBetween(params ...string) time.Time

	// DateTimeInterval Get a  Time object between January 1, 1970, and `until` (defaults to "now").
	//
	// Parameters:
	//   - name: description.
	// Returns:
	//
	//	- type: description
	//
	// Example usage:
	//
	DateTimeInterval(params ...string) time.Time
	// DateTimeThisWeek Get a  Time object between January 1, 1970, and `until` (defaults to "now").
	//
	// Parameters:
	//   - name: description.
	// Returns:
	//
	//	- type: description
	//
	// Example usage:
	//
	DateTimeThisWeek(params ...string) time.Time
	// DateTimeThisMonth Get a  Time object between January 1, 1970, and `until` (defaults to "now").
	//
	// Parameters:
	//   - name: description.
	// Returns:
	//
	//	- type: description
	//
	// Example usage:
	//
	DateTimeThisMonth(params ...string) time.Time
	// DateTimeThisYear Get a  Time object between January 1, 1970, and `until` (defaults to "now").
	//
	// Parameters:
	//   - name: description.
	// Returns:
	//
	//	- type: description
	//
	// Example usage:
	//
	DateTimeThisYear(params ...string) time.Time
	// DateTimeThisDecade Get a  Time object between January 1, 1970, and `until` (defaults to "now").
	//
	// Parameters:
	//   - name: description.
	// Returns:
	//
	//	- type: description
	//
	// Example usage:
	//
	DateTimeThisDecade(params ...string) time.Time
	// DateTimeThisCentury Get a  Time object between January 1, 1970, and `until` (defaults to "now").
	//
	// Parameters:
	//   - name: description.
	// Returns:
	//
	//	- type: description
	//
	// Example usage:
	//
	DateTimeThisCentury(params ...string) time.Time

	// Date Get a  Time object between January 1, 1970, and `until` (defaults to "now").
	//
	// Parameters:
	//   - name: description.
	// Returns:
	//
	//	- type: description
	//
	// Example usage:
	//
	Date(params ...string) string
	// Time Get a  Time object between January 1, 1970, and `until` (defaults to "now").
	//
	// Parameters:
	//   - name: description.
	// Returns:
	//
	//	- type: description
	//
	// Example usage:
	//
	Time(params ...string) string

	// UnixTime Get a  Time object between January 1, 1970, and `until` (defaults to "now").
	//
	// Parameters:
	//   - name: description.
	// Returns:
	//
	//	- type: description
	//
	// Example usage:
	//
	UnixTime(until ...string) string

	// ISO8602 Get a  Time object between January 1, 1970, and `until` (defaults to "now").
	//
	// Parameters:
	//   - name: description.
	// Returns:
	//
	//	- type: description
	//
	// Example usage:
	//
	ISO8602(until ...string) string

	// AmPm Get a  Time object between January 1, 1970, and `until` (defaults to "now").
	//
	// Parameters:
	//   - name: description.
	// Returns:
	//
	//	- type: description
	//
	// Example usage:
	//
	AmPm(until ...string) string

	// DayOfMonth Get a Time object between January 1, 1970, and `until` (defaults to "now").
	//
	// Parameters:
	//   - name: description.
	// Returns:
	//
	//	- type: description
	//
	// Example usage:
	//
	DayOfMonth(until ...string) string

	// DayOfWeek Get a  Time object between January 1, 1970, and `until` (defaults to "now").
	//
	// Parameters:
	//   - name: description.
	// Returns:
	//
	//	- type: description
	//
	// Example usage:
	//
	DayOfWeek(until ...string) string
	// Month Get a  Time object between January 1, 1970, and `until` (defaults to "now").
	//
	// Parameters:
	//   - name: description.
	// Returns:
	//
	//	- type: description
	//
	// Example usage:
	//
	Month(until ...string) string
	// MonthName Get a  Time object between January 1, 1970, and `until` (defaults to "now").
	//
	// Parameters:
	//   - name: description.
	// Returns:
	//
	//	- type: description
	//
	// Example usage:
	//
	MonthName(until ...string) string
	// Year Get a  Time object between January 1, 1970, and `until` (defaults to "now").
	//
	// Parameters:
	//   - name: description.
	// Returns:
	//
	//	- type: description
	//
	// Example usage:
	//
	Year(until ...string) string
	// Century Get a  Time object between January 1, 1970, and `until` (defaults to "now").
	//
	// Parameters:
	//   - name: description.
	// Returns:
	//
	//	- type: description
	//
	// Example usage:
	//
	Century() string
	// Timezone Get a  Time object between January 1, 1970, and `until` (defaults to "now").
	//
	// Parameters:
	//   - name: description.
	// Returns:
	//
	//	- type: description
	//
	// Example usage:
	//
	Timezone(countryCode ...string) string
}
