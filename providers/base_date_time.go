package providers

import (
	"errors"
	"fmt"
	"github.com/harmlessprince/goFaker/constants"
	"github.com/harmlessprince/goFaker/parsers"
	"math"
	"strconv"
	"time"
)

type DateTimeInterface interface {
	SetDefaultTimeZone(timeZone string)
	DateTime(max interface{}, timezone string) (time.Time, error)
	DateTimeAD(max interface{}, timezone string) (time.Time, error)
	DateTimeBetween(startDate interface{}, endDate interface{}, timezone string) (time.Time, error)
	DateTimeInterval(dateString string, intervalString string, timezone string) (time.Time, error)
	DateTimeThisWeek(max string, timezone string) (time.Time, error)
	DateTimeThisMonth(max string, timezone string) (time.Time, error)
	DateTimeThisYear(max string, timezone string) (time.Time, error)
	DateTimeThisDecade(max string, timezone string) (time.Time, error)
	DateTimeThisCentury(max string, timezone string) (time.Time, error)
	Date(format string, max interface{}) (string, error)
	Time(format string, max interface{}) (string, error)
	UnixTime(max interface{}) (int, error)
	ISO8602(max interface{}) (string, error)
	AmPm(max interface{}) (string, error)
	DayOfMonth(max interface{}) (int, error)
	DayOfWeek(max interface{}) (time.Weekday, error)
	Month(max interface{}) (int, error)
	MonthName(max interface{}) (string, error)
	Year(max interface{}) (string, error)
	Century() (string, error)
	Timezone(countryCode string) (*time.Location, error)
}
type BaseDateTime struct {
	BaseProvider
	defaultTimeZone string
}

func (bdt *BaseDateTime) SetDefaultTimeZone(timeZone string) {
	bdt.defaultTimeZone = timeZone
}
func (bdt *BaseDateTime) getDefaultTimeZone() string {
	return bdt.defaultTimeZone
}

func (bdt *BaseDateTime) getMaxTimestamp(max interface{}) (int64, error) {
	switch maxType := max.(type) {
	case int:
		return int64(maxType), nil
	case float64:
		return int64(maxType), nil
	case string:
		if max == "" {
			return time.Now().UnixNano(), nil
		}
		timestamp, err := time.Parse(time.DateTime, maxType)
		if err != nil {
			return 0, err
		}
		return timestamp.Unix(), nil
	case time.Time:
		return maxType.Unix(), nil
	default:
		return 0, fmt.Errorf("unsupported type for max: %T", max)
	}
}
func (bdt *BaseDateTime) setTimezone(dt time.Time, timezone string) (time.Time, error) {
	location, err := time.LoadLocation(bdt.resolveTimezone(timezone))
	if err != nil {
		return time.Time{}, err
	}
	return dt.In(location), nil
}

// resolveTimezone
// Parameters:
//   - timezone : string|""
//
// Returns:
//   - string
func (bdt *BaseDateTime) resolveTimezone(timezone string) string {
	if timezone == "" {
		if bdt.defaultTimeZone == "" {
			return time.Local.String()
		}
		return bdt.getDefaultTimeZone()
	}

	return timezone
}

// DateTime Get a  Time object between January 1, 1970, and now.
//
// Parameters:
//   - max: time.Time|int|string  maximum timestamp used as random end limit, default to time.Now()
//   - timezone: string time zone location in which the date time should be set, default to BaseDateTime.defaultTimezone,
//
// Returns:
//   - Time.time : description
//
// Example usage:
//
//	baseDateTime.DateTime("2024-01-01", "Africa/lagos")
//
// See https://www.ibm.com/docs/en/ts7650g-protectier/3.4.0?topic=reference-worldwide-time-zone-codes
func (bdt *BaseDateTime) DateTime(max interface{}, timezone string) (time.Time, error) {
	unixTimestamp, err := bdt.UnixTime(max)
	if err != nil {
		return time.Time{}, err
	}
	return bdt.setTimezone(time.Unix(int64(unixTimestamp), 0), timezone)
}

// DateTimeAD Get a Time object between January 1, 1970, and now.
//
// Parameters:
//   - max: time.Time|int|string  maximum timestamp used as random end limit, default to time.Now()
//   - timezone: *time.Location  time zone location in which the date time should be set, default to BaseDateTime.defaultTimezone,
//
// Returns:
//   - Time.time : description
//
// Example usage:
//
//	baseDateTime.DateTimeAD("2024-01-01", "Africa/lagos")
//
// See https://www.ibm.com/docs/en/ts7650g-protectier/3.4.0?topic=reference-worldwide-time-zone-codes
func (bdt *BaseDateTime) DateTimeAD(max interface{}, timezone string) (time.Time, error) {
	maxTimeStamp, err := bdt.getMaxTimestamp(max)
	if err != nil {
		return time.Time{}, err
	}
	numberBTw, err := bdt.NumberBetween(math.MinInt64, int(maxTimeStamp))
	if err != nil {
		return time.Time{}, err
	}
	return bdt.setTimezone(time.Unix(int64(numberBTw), 0), timezone)
}

// DateTimeBetween Get a Time object based on a random date between two given dates.
// Accepts date strings that can be recognized by Time.Parse().
//
// Parameters:
//   - max: time.Time|int|string  maximum timestamp used as random end limit, default to time.Now()
//   - format: string e.g   "15:04:05"
//
// Returns:
//   - int : "2020292"
//
// Example usage:
//   - baseDateTime.DateTimeBetween("2026-01-02", "2036-01-02", "Africa/lagos")
//
// See https://www.ibm.com/docs/en/ts7650g-protectier/3.4.0?topic=reference-worldwide-time-zone-codes
func (bdt *BaseDateTime) DateTimeBetween(startDate interface{}, endDate interface{}, timezone string) (time.Time, error) {

	var startTimestamp int64
	var endTimeStamp int64
	switch startDateType := startDate.(type) {
	case time.Time:
		startTimestamp = startDateType.Unix()
	case string:
		timestamp, err := time.Parse(time.DateTime, startDateType)
		if err != nil {
			return time.Time{}, err
		}
		startTimestamp = timestamp.Unix()
	default:
		thirtyYearsAgo := time.Now().AddDate(-30, 0, 0)
		startTimestamp = thirtyYearsAgo.Unix()
	}

	switch endDateType := endDate.(type) {
	case time.Time:
		endTimeStamp = endDateType.Unix()
	case string:
		timestamp, err := time.Parse(time.DateTime, endDateType)
		if err != nil {
			return time.Time{}, err
		}
		endTimeStamp = timestamp.Unix()
	default:
		var err error
		endTimeStamp, err = bdt.getMaxTimestamp(endDate)
		if err != nil {
			return time.Time{}, err
		}
	}

	if startTimestamp > endTimeStamp {
		return time.Time{}, errors.New("start date must be anterior to end date")
	}
	timestampBtw, errBtw := bdt.NumberBetween(int(startTimestamp), int(endTimeStamp))
	if errBtw != nil {
		return time.Time{}, errBtw
	}
	return bdt.setTimezone(time.Unix(int64(timestampBtw), 0), timezone)
}

// DateTimeInterval
// Get a Time object based on a random date between one date given and an interval.
// Accepts date strings that can be recognized by helpers.UnitMap.
//
// Parameters:
//   - dateString: string maximum timestamp used as random end limit, default to time.Now()
//   - intervalString: string e.g   "4years"
//   - timezone: string e.g. "Africa/lagos"
//
// Returns:
//   - time.Time
//
// Example usage:
//   - baseDateTime.DateTimeInterval(time.Now().Format(time.RFC3339), "-5d", "Africa/lagos")
//
// See https://www.ibm.com/docs/en/ts7650g-protectier/3.4.0?topic=reference-worldwide-time-zone-codes
func (bdt *BaseDateTime) DateTimeInterval(dateString string, intervalString string, timezone string) (time.Time, error) {

	if dateString == "" {
		interval, err := parsers.StringToDuration{}.ParseDuration("-30years")
		if err != nil {
			return time.Time{}, err
		}
		dateString = time.Now().Add(interval).Format(time.RFC3339)
	}
	datetime, err := time.Parse(time.RFC3339, dateString)
	if err != nil {
		return time.Time{}, err
	}

	interval, err := parsers.StringToDuration{}.ParseDuration(intervalString)
	if err != nil {
		return time.Time{}, err
	}
	otherDatetime := datetime.Add(interval)
	begin := datetime
	end := otherDatetime
	if datetime.After(otherDatetime) {
		begin, end = otherDatetime, datetime
	}
	return bdt.DateTimeBetween(begin, end, timezone)
}

// DateTimeThisWeek
// Get a date time object somewhere within a week.
//
// Parameters:
//   - max: time.Time|int|string  maximum timestamp used as random end limit, default to time.Now()
//   - timezone: string time zone location in which the date time should be set, default to BaseDateTime.defaultTimezone,
//
// Returns:
//   - time.Time
//
// Example usage:
//   - baseDateTime.DateTimeThisCentury(time.Now().Format(time.DateTime), "")
//
// See https://www.ibm.com/docs/en/ts7650g-protectier/3.4.0?topic=reference-worldwide-time-zone-codes
func (bdt *BaseDateTime) DateTimeThisWeek(max string, timezone string) (time.Time, error) {
	if max == "" {
		max = time.Now().Format(time.DateTime)
	}
	startDateDuration, err := parsers.StringToDuration{}.ParseDuration("-1w")
	if err != nil {
		return time.Time{}, err
	}
	startDateTime := time.Now().Add(startDateDuration)
	return bdt.DateTimeBetween(startDateTime, max, timezone)
}

// DateTimeThisMonth
// Get a date time object somewhere within a month.
//
// Parameters:
//   - max: time.Time|int|string  maximum timestamp used as random end limit, default to time.Now()
//   - timezone: string time zone location in which the date time should be set, default to BaseDateTime.defaultTimezone,
//
// Returns:
//   - time.Time
//
// Example usage:
//   - baseDateTime.DateTimeThisMonth(time.Now().Format(time.DateTime), "")
//
// See https://www.ibm.com/docs/en/ts7650g-protectier/3.4.0?topic=reference-worldwide-time-zone-codes
func (bdt *BaseDateTime) DateTimeThisMonth(max string, timezone string) (time.Time, error) {
	if max == "" {
		max = time.Now().Format(time.DateTime)
	}
	startDateDuration, err := parsers.StringToDuration{}.ParseDuration("-1months")
	if err != nil {
		return time.Time{}, err
	}
	startDateTime := time.Now().Add(startDateDuration)
	return bdt.DateTimeBetween(startDateTime, max, timezone)
}

// DateTimeThisYear
// Get a date time object somewhere inside the current year.
//
// Parameters:
//   - max: time.Time|int|string  maximum timestamp used as random end limit, default to time.Now()
//   - timezone: string time zone location in which the date time should be set, default to BaseDateTime.defaultTimezone,
//
// Returns:
//   - time.Time
//
// Example usage:
//   - baseDateTime.DateTimeThisYear(time.Now().Format(time.DateTime), "")
//
// See https://www.ibm.com/docs/en/ts7650g-protectier/3.4.0?topic=reference-worldwide-time-zone-codes
func (bdt *BaseDateTime) DateTimeThisYear(max string, timezone string) (time.Time, error) {
	if max == "" {
		max = time.Now().Format(time.DateTime)
	}
	now := time.Now()
	startDateTime := time.Date(now.Year(), time.January, 1, 0, 0, 0, 0, now.Location())
	return bdt.DateTimeBetween(startDateTime, max, timezone)
}

// DateTimeThisDecade
// Get a date time object somewhere within a century.
//
// Parameters:
//   - max: time.Time|int|string  maximum timestamp used as random end limit, default to time.Now()
//   - timezone: string time zone location in which the date time should be set, default to BaseDateTime.defaultTimezone,
//
// Returns:
//   - time.Time
//
// Example usage:
//   - baseDateTime.DateTimeThisDecade(time.Now().Format(time.DateTime), "")
//
// See https://www.ibm.com/docs/en/ts7650g-protectier/3.4.0?topic=reference-worldwide-time-zone-codes
func (bdt *BaseDateTime) DateTimeThisDecade(max string, timezone string) (time.Time, error) {
	if max == "" {
		max = time.Now().Format(time.DateTime)
	}
	startDateDuration, err := parsers.StringToDuration{}.ParseDuration("-10years")
	if err != nil {
		return time.Time{}, err
	}
	startDateTime := time.Now().Add(startDateDuration)
	return bdt.DateTimeBetween(startDateTime, max, timezone)
}

// DateTimeThisCentury
// Get a date time object somewhere within a century.
//
// Parameters:
//   - max: time.Time|int|string  maximum timestamp used as random end limit, default to time.Now()
//   - timezone: string time zone location in which the date time should be set, default to BaseDateTime.defaultTimezone,
//
// Returns:
//   - time.Time
//
// Example usage:
//   - baseDateTime.DateTimeThisCentury(time.Now().Format(time.DateTime), "")
//
// See https://www.ibm.com/docs/en/ts7650g-protectier/3.4.0?topic=reference-worldwide-time-zone-codes
func (bdt *BaseDateTime) DateTimeThisCentury(max string, timezone string) (time.Time, error) {
	if max == "" {
		max = time.Now().Format(time.DateTime)
	}
	startDateDuration, err := parsers.StringToDuration{}.ParseDuration("-100years")
	if err != nil {
		return time.Time{}, err
	}
	startDateTime := time.Now().Add(startDateDuration)
	return bdt.DateTimeBetween(startDateTime, max, timezone)
}

// Date return a date string between January 1, 1970, and now
//
// Parameters:
//   - max: time.Time|int|string  maximum timestamp used as random end limit, default to time.Now()
//   - format: string e.g  "2006-01-02"
//
// Returns:
//   - string : "2020-01-02"
//
// Example usage:
//
//	baseDateTime.Date("2006-01-01", "")
func (bdt *BaseDateTime) Date(format string, max interface{}) (string, error) {
	if format == "" {
		format = "2006-01-02" //Y-m-d
	}

	dateTime, err := bdt.DateTime(max, "")
	if err != nil {
		return "", err
	}
	return dateTime.Format(format), nil
}

// Time return a time string (24h format by default)
//
// Parameters:
//   - max: time.Time|int|string  maximum timestamp used as random end limit, default to time.Now()
//   - format: string e.g   "15:04:05"
//
// Returns:
//   - string : "2020-01-02"
//
// Example usage:
//   - baseDateTime.Time("15:04:05", "")
func (bdt *BaseDateTime) Time(format string, max interface{}) (string, error) {
	if format == "" {
		format = time.TimeOnly
	}
	dateTime, err := bdt.DateTime(max, "")
	if err != nil {
		return "", err
	}
	return dateTime.Format(format), nil
}

// UnixTime Get a timestamp between January 1, 1970, and now
//
// Parameters:
//   - max: time.Time|int|string  maximum timestamp used as random end limit, default to time.Now()
//   - format: string e.g   "15:04:05"
//
// Returns:
//   - int : "2020292"
//
// Example usage:
//   - baseDateTime.UnixTime("")
func (bdt *BaseDateTime) UnixTime(max interface{}) (int, error) {
	maxTimeStamp, errMaxT := bdt.getMaxTimestamp(max)
	if errMaxT != nil {
		return 0, errMaxT
	}
	numberBtw, errBtw := bdt.NumberBetween(0, int(maxTimeStamp))
	if errBtw != nil {
		return 0, errBtw
	}
	return numberBtw, nil
}

func (bdt *BaseDateTime) ISO8602(max interface{}) (string, error) {
	return bdt.Date(time.RFC3339, max)
}

// AmPm
// Get a date time object somewhere within a month.
//
// Parameters:
//   - max: time.Time|int|string  maximum timestamp used as random end limit, default to time.Now()
//
// Returns:
//   - time.Time
//
// Example usage:
//   - baseDateTime.DateTimeThisMonth("")
func (bdt *BaseDateTime) AmPm(max interface{}) (string, error) {
	if max == "" {
		max = time.Now().Format(time.DateTime)
	}
	dateTime, err := bdt.DateTime(max, "")
	if err != nil {
		return "", err
	}
	return dateTime.Format("PM"), nil
}

// DayOfMonth
// Get a month number.
//
// Parameters:
//   - max: time.Time|int|string  maximum timestamp used as random end limit, default to time.Now()
//
// Returns:
//   - int e.g 12
//
// Example usage:
//   - baseDateTime.DayOfMonth("")
func (bdt *BaseDateTime) DayOfMonth(max interface{}) (int, error) {
	if max == "" {
		max = time.Now().Format(time.DateTime)
	}
	dateTime, err := bdt.DateTime(max, "")
	if err != nil {
		return 0, err
	}
	return dateTime.Day(), nil
}

// DayOfWeek
// Get a week name.
//
// Parameters:
//   - max: time.Time|int|string  maximum timestamp used as random end limit, default to time.Now()
//
// Returns:
//   - int e.g 12
//
// Example usage:
//   - baseDateTime.DayOfWeek("")
func (bdt *BaseDateTime) DayOfWeek(max interface{}) (time.Weekday, error) {
	if max == "" {
		max = time.Now().Format(time.DateTime)
	}
	dateTime, err := bdt.DateTime(max, "")
	if err != nil {
		return 0, err
	}
	return dateTime.Weekday(), nil
}

// Month
// Get a month number.
//
// Parameters:
//   - max: time.Time|int|string  maximum timestamp used as random end limit, default to time.Now()
//
// Returns:
//   - int e.g 1
//
// Example usage:
//   - baseDateTime.Month("")
func (bdt *BaseDateTime) Month(max interface{}) (int, error) {
	if max == "" {
		max = time.Now().Format(time.DateTime)
	}
	dateTime, err := bdt.DateTime(max, "")
	if err != nil {
		return 0, err
	}
	return int(dateTime.Month()), nil
}

// MonthName
// Get a month name.
//
// Parameters:
//   - max: time.Time|int|string  maximum timestamp used as random end limit, default to time.Now()
//
// Returns:
//   - int e.g 12
//
// Example usage:
//   - baseDateTime.MonthName("")
func (bdt *BaseDateTime) MonthName(max interface{}) (string, error) {
	if max == "" {
		max = time.Now().Format(time.DateTime)
	}
	dateTime, err := bdt.DateTime(max, "")
	if err != nil {
		return "", err
	}
	return dateTime.Month().String(), nil
}

// Year
// Get a year from date time.
//
// Parameters:
//   - max: time.Time|int|string  maximum timestamp used as random end limit, default to time.Now()
//
// Returns:
//   - int e.g 12
//
// Example usage:
//   - baseDateTime.Year("")
func (bdt *BaseDateTime) Year(max interface{}) (string, error) {
	if max == "" {
		max = time.Now().Format(time.DateTime)
	}
	dateTime, err := bdt.DateTime(max, "")
	if err != nil {
		return "", err
	}
	return strconv.Itoa(dateTime.Year()), nil
}

func (bdt *BaseDateTime) Century() (string, error) {
	century, err := bdt.RandomElementFromStringSlice(constants.Century)
	if err != nil {
		return "", err
	}
	return century, nil
}

func (bdt *BaseDateTime) Timezone(countryCode string) (*time.Location, error) {
	timeZones := []string{
		"America/New_York",
		"America/Los_Angeles",
		"America/Chicago",
		"Europe/London",
		"Europe/Paris",
		"Europe/Berlin",
		"Asia/Tokyo",
		"Asia/Dubai",
		"Asia/Kolkata",
		"Australia/Sydney",
		"Australia/Melbourne",
		"Australia/Perth",
		"Africa/Cairo",
		"Africa/Nairobi",
		"Africa/Lagos",
	}
	timezoneLocation, _ := time.LoadLocation(countryCode)
	if countryCode == "" {
		randomTimeZone, err := bdt.RandomElementFromStringSlice(timeZones)
		if err != nil {
			return nil, err
		}
		timezoneLocation, err = time.LoadLocation(randomTimeZone)
		if err != nil {
			return nil, err
		}
	}
	return timezoneLocation, nil
}
