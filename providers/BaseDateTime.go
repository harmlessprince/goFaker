package providers

import (
	"fmt"
	"github.com/harmlessprince/goFaker/constants"
	"github.com/harmlessprince/goFaker/helpers"
	"math"
	"strconv"
	"time"
)

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

func (bdt *BaseDateTime) getMaxTimestamp(max interface{}) (int, error) {
	switch maxType := max.(type) {
	case int:
		return maxType, nil
	case float64:
		return int(maxType), nil
	case string:
		if max == "" {
			return int(time.Now().UnixNano()), nil
		}
		timestamp, err := time.Parse(time.DateTime, maxType)
		if err != nil {
			panic(err)
		}
		return int(timestamp.Unix()), nil
	case time.Time:
		return int(maxType.Unix()), nil
	default:
		return 0, fmt.Errorf("unsupported type for max: %T", max)
	}
}
func (bdt *BaseDateTime) setTimezone(dt time.Time, timezone string) time.Time {
	location, err := time.LoadLocation(bdt.resolveTimezone(timezone))
	if err != nil {
		panic(err)
	}
	return dt.In(location)
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
func (bdt *BaseDateTime) DateTime(max interface{}, timezone string) time.Time {
	unixTimestamp := bdt.UnixTime(max)
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
func (bdt *BaseDateTime) DateTimeAD(max interface{}, timezone string) time.Time {
	maxTimeStamp, err := bdt.getMaxTimestamp(max)
	if err != nil {
		panic(err)
	}
	numberBTw, errNumbBtw := bdt.NumberBetween(math.MinInt64, maxTimeStamp)
	if errNumbBtw != nil {
		panic(errNumbBtw)
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
func (bdt *BaseDateTime) DateTimeBetween(startDate interface{}, endDate interface{}, timezone string) time.Time {
	var startTimestamp int64
	switch startDateType := startDate.(type) {
	case time.Time:
		startTimestamp = startDateType.Unix()
	case string:
		timestamp, err := time.Parse(time.DateTime, startDateType)
		if err != nil {
			panic(err)
			return time.Time{}
		}
		startTimestamp = timestamp.Unix()
	default:
		thirtyYearsAgo := time.Now().AddDate(-30, 0, 0)
		startTimestamp = thirtyYearsAgo.Unix()
	}
	endTimeStamp, err := bdt.getMaxTimestamp(endDate)
	if err != nil {
		panic(err)
		return time.Time{}
	}
	if startTimestamp > int64(endTimeStamp) {
		panic("start date must be anterior to end date")
		return time.Time{}
	}
	timestampBtw, errBtw := bdt.NumberBetween(int(startTimestamp), endTimeStamp)
	if errBtw != nil {
		panic(errBtw)
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
func (bdt *BaseDateTime) DateTimeInterval(dateString string, intervalString string, timezone string) time.Time {

	if dateString == "" {
		interval, err := helpers.StringToDuration{}.ParseDuration("-30years")
		if err != nil {
			panic(err)
			return time.Time{}
		}
		dateString = time.Now().Add(interval).Format(time.RFC3339)
	}
	datetime, err := time.Parse(time.RFC3339, dateString)
	if err != nil {
		panic(err)
		return time.Time{}
	}

	interval, err := helpers.StringToDuration{}.ParseDuration(intervalString)
	if err != nil {
		panic(err)
		return time.Time{}
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
func (bdt *BaseDateTime) DateTimeThisWeek(max string, timezone string) time.Time {
	if max == "" {
		max = time.Now().Format(time.DateTime)
	}
	startDateDuration, err := helpers.StringToDuration{}.ParseDuration("-1w")
	if err != nil {
		panic(err)
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
func (bdt *BaseDateTime) DateTimeThisMonth(max string, timezone string) time.Time {
	if max == "" {
		max = time.Now().Format(time.DateTime)
	}
	startDateDuration, err := helpers.StringToDuration{}.ParseDuration("-1months")
	if err != nil {
		panic(err)
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
func (bdt *BaseDateTime) DateTimeThisYear(max string, timezone string) time.Time {
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
func (bdt *BaseDateTime) DateTimeThisDecade(max string, timezone string) time.Time {
	if max == "" {
		max = time.Now().Format(time.DateTime)
	}
	startDateDuration, err := helpers.StringToDuration{}.ParseDuration("-10years")
	if err != nil {
		panic(err)
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
func (bdt *BaseDateTime) DateTimeThisCentury(max string, timezone string) time.Time {
	if max == "" {
		max = time.Now().Format(time.DateTime)
	}
	startDateDuration, err := helpers.StringToDuration{}.ParseDuration("-100years")
	if err != nil {
		panic(err)
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
func (bdt *BaseDateTime) Date(format string, max interface{}) string {
	if format == "" {
		format = "2006-01-02" //Y-m-d
	}

	dateTime := bdt.DateTime(max, "")
	return dateTime.Format(format)
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
func (bdt *BaseDateTime) Time(format string, max interface{}) string {
	if format == "" {
		format = time.TimeOnly
	}
	return bdt.DateTime(max, "").Format(format)
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
func (bdt *BaseDateTime) UnixTime(max interface{}) int {
	maxTimeStamp, errMaxT := bdt.getMaxTimestamp(max)
	if errMaxT != nil {
		panic(errMaxT)
	}
	numberBtw, errBtw := bdt.NumberBetween(0, maxTimeStamp)
	if errBtw != nil {
		panic(errBtw)
	}
	return numberBtw
}

func (bdt *BaseDateTime) ISO8602(max interface{}) string {
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
func (bdt *BaseDateTime) AmPm(max interface{}) string {
	if max == "" {
		max = time.Now().Format(time.DateTime)
	}
	return bdt.DateTime(max, "").Format("PM")
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
func (bdt *BaseDateTime) DayOfMonth(max interface{}) int {
	if max == "" {
		max = time.Now().Format(time.DateTime)
	}
	return bdt.DateTime(max, "").Day()
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
func (bdt *BaseDateTime) DayOfWeek(max interface{}) time.Weekday {
	if max == "" {
		max = time.Now().Format(time.DateTime)
	}
	return bdt.DateTime(max, "").Weekday()
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
func (bdt *BaseDateTime) Month(max interface{}) int {
	if max == "" {
		max = time.Now().Format(time.DateTime)
	}
	return int(bdt.DateTime(max, "").Month())
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
func (bdt *BaseDateTime) MonthName(max interface{}) string {
	if max == "" {
		max = time.Now().Format(time.DateTime)
	}
	return bdt.DateTime(max, "").Month().String()
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
func (bdt *BaseDateTime) Year(max interface{}) string {
	if max == "" {
		max = time.Now().Format(time.DateTime)
	}
	return strconv.Itoa(bdt.DateTime(max, "").Year())
}

func (bdt *BaseDateTime) Century() string {
	century, err := bdt.RandomElementFromStringSlice(constants.Century)
	if err != nil {
		panic(err)
	}
	return century
}

func (bdt *BaseDateTime) Timezone(countryCode string) *time.Location {
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
	var timezoneLocation *time.Location
	if countryCode == "" {
		randomTimeZone, err := bdt.RandomElementFromStringSlice(timeZones)
		if err != nil {
			panic(err)
		}
		timezoneLocation, err = time.LoadLocation(randomTimeZone)
		if err != nil {
			panic(err)
		}
	} else {
		timezoneLocation, _ = time.LoadLocation(countryCode)
	}
	return timezoneLocation
}
