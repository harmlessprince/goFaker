package extensions

import "time"

type DateTimeExtension interface {
	SetDefaultTimeZone(timeZone string)
	DateTime(max interface{}, timezone string) time.Time
	DateTimeAD(max interface{}, timezone string) time.Time
	DateTimeBetween(startDate interface{}, endDate interface{}, timezone string) time.Time
	DateTimeInterval(dateString string, intervalString string, timezone string) time.Time
	DateTimeThisWeek(max string, timezone string) time.Time
	DateTimeThisMonth(max string, timezone string) time.Time
	DateTimeThisYear(max string, timezone string) time.Time
	DateTimeThisDecade(max string, timezone string) time.Time
	DateTimeThisCentury(max string, timezone string) time.Time
	Date(format string, max interface{}) string
	Time(format string, max interface{}) string
	UnixTime(max interface{}) int
	ISO8602(max interface{}) string
	AmPm(max interface{}) string
	DayOfMonth(max interface{}) int
	DayOfWeek(max interface{}) time.Weekday
	Month(max interface{}) int
	MonthName(max interface{}) string
	Year(max interface{}) string
	Century() string
	Timezone(countryCode string) *time.Location
}
