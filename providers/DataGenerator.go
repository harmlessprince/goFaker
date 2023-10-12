package providers

type DataGenerator interface {
	//extensions.InternetExtension
	GenerateAddress() string
	GenerateCity() string
	GenerateCountry() string
	GeneratePostCode() string
	GenerateStreetName() string
	GenerateStreetAddress() string
	GenerateBuildingNumber() string
	GenerateLocalCoordinates() map[string]float64
	GenerateLongitude() float64
	GenerateLatitude() float64
	GenerateCitySuffix() string
	GenerateCityName() string
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
	GenerateCompany() string
	GenerateCompanySuffix() string
	GenerateJobTitle() string
	GenerateName(gender ...string) string
	GenerateFirstName(gender ...string) string
	GenerateFirstNameMale() string
	GenerateFirstNameFemale() string
	GenerateLastName() string
	GenerateTitle(gender ...string) string
	GenerateTitleMale() string
	GenerateTitleFemale() string
	GeneratePhoneNumber() string
	GenerateE164PhoneNumber() string
	GenerateImei() string
	GenerateEmail() string
	GenerateSafeEmail() string
	GenerateFreeEmail() string
	GenerateCompanyEmail() string
	GenerateFreeEmailDomain() string
	GenerateSafeEmailDomain() string
	GenerateUserName() string
	GeneratePassword(params ...int) string
	GenerateDomainName() string
	GenerateDomainWord() string
	GenerateUrl() string
	GenerateTld() string
	GenerateIpv4() string
	GenerateIpv6() string
	GenerateLocalIpv4() string
	GenerateMacAddress() string
	GenerateTransliterate(inputString string) string
	GenerateToAscii(inputString string) string
	GenerateSlug(numberOfWords int, variableNumberOfWords bool) string
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
