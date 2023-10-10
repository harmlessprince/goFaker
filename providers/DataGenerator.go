package providers

type DataGenerator interface {
	GenerateAddress() string
	City() string
	PostCode() string
	StreetName() string
	StreetAddress() string
	BuildingNumber() string
	LocalCoordinates() map[string]float64
	Longitude() float64
	Latitude() float64
	CitySuffix() string
	CityName() string
	//Ean13() string
	//Ean8() string
	//Isbn10() string
	//Isbn13() string
	//BloodType() string
	//BloodRh() string
	//BloodGroup() string
	//HexColor() string
	//SafeHexColor() string
	//RgbColorAsArray() string
	//RgbColor() string
	//RgbCssColor() string
	//RgbaCssColor() string
	//SafeColorName() string
	//ColorName() string
	//HslColor() string
	//HslColorAsArray() []int
	//Company() string
	//CompanySuffix() string
	//JobTitle() string
	//Name(gender ...string) string
	//FirstName(gender ...string) string
	//FirstNameMale() string
	//FirstNameFemale() string
	//LastName() string
	//Title(gender ...string) string
	//TitleMale() string
	//TitleFemale() string
	//PhoneNumber() string
	//E164PhoneNumber() string
	//Imei() string
	//Uuid3() string
	//Semver(params ...bool)
	//MimeType() string
	//Extension() string
	//FilePath() string
	//DateTime(params ...string) time.Time
	//DateTimeAD(params ...string) time.Time
	//DateTimeBetween(params ...string) time.Time
	//DateTimeInterval(params ...string) time.Time
	//DateTimeThisWeek(params ...string) time.Time
	//DateTimeThisMonth(params ...string) time.Time
	//DateTimeThisYear(params ...string) time.Time
	//DateTimeThisDecade(params ...string) time.Time
	//DateTimeThisCentury(params ...string) time.Time
	//Date(params ...string) string
	//Time(params ...string) string
	//UnixTime(until ...string) string
	//ISO8602(until ...string) string
	//AmPm(until ...string) string
	//DayOfMonth(until ...string) string
	//DayOfWeek(until ...string) string
	//Month(until ...string) string
	//MonthName(until ...string) string
	//Year(until ...string) string
	//Century() string
	//Timezone(countryCode ...string) string
}
