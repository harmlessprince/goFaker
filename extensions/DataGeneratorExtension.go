package extensions

import (
	"time"
)

type DataGeneratorExtension interface {

	// GenerateAddress generates a random address string.
	//
	// This method generates a random address string that can be used for various purposes such as testing, example data, or placeholders.
	//
	// Returns:
	// - A string representing a randomly generated address.
	GenerateAddress() string

	// GenerateCity generates a random city name.
	//
	// Use this method to generate random city names for testing, data generation, or other purposes.
	//
	// Returns:
	// - A string representing a randomly generated city name.
	GenerateCity() string
	// GenerateCountry generates a random country name.
	//
	// This method is used to generate random country names for various applications such as generating test data or example content.
	//
	// Returns:
	// - A string representing a randomly generated country name.
	GenerateCountry() string
	// GeneratePostCode generates a random postal code.
	//
	// Use this method to generate random postal codes that can be used for testing or other purposes.
	//
	// Returns:
	// - A string representing a randomly generated postal code.
	GeneratePostCode() string
	// GenerateStreetName generates a random street name.
	//
	// This method generates a random street name, suitable for test data, sample content, or other applications.
	//
	// Returns:
	// - A string representing a randomly generated street name.
	GenerateStreetName() string
	// GenerateStreetAddress generates a random street address.
	//
	// Use this method to generate random street addresses for various purposes such as generating test data or example content.
	//
	// Returns:
	// - A string representing a randomly generated street address.
	GenerateStreetAddress()
	// GenerateBuildingNumber generates a random building number.
	//
	// This method generates a random building number that can be used for testing, example data, or placeholders.
	//
	// Returns:
	// - A string representing a randomly generated building number.string
	GenerateBuildingNumber() string
	// GenerateLocalCoordinates generates random local coordinates (latitude and longitude).
	//
	// Use this method to generate random latitude and longitude coordinates for testing or other purposes.
	//
	// Returns:
	// - A map[string]float64 with 'latitude' and 'longitude' keys representing the generated coordinates.

	GenerateLocalCoordinates() map[string]float64
	// GenerateLongitude generates a random longitude coordinate.
	//
	// This method generates a random longitude coordinate suitable for testing or other applications.
	//
	// Returns:
	// - A float64 representing a randomly generated longitude coordinate.
	GenerateLongitude() float64
	// GenerateLatitude generates a random latitude coordinate.
	//
	// Use this method to generate random latitude coordinates for various purposes such as generating test data or example content.
	//
	// Returns:
	// - A float64 representing a randomly generated latitude coordinate.
	GenerateLatitude() float64
	// GenerateCitySuffix generates a random city suffix.
	//
	// This method generates a random city suffix that can be used for creating realistic city names or test data.
	//
	// Returns:
	// - A string representing a randomly generated city suffix.
	GenerateCitySuffix() string
	// GenerateCityName generates a random city name.
	//
	// Use this method to generate random city names for testing, data generation, or other purposes.
	//
	// Returns:
	// - A string representing a randomly generated city name.
	GenerateCityName() string

	// GenerateHexColor generates a random hexadecimal color code.
	//
	// Returns:
	// - A string representing a randomly generated hexadecimal color code.
	GenerateHexColor() string
	// GenerateSafeHexColor generates a random safe hexadecimal color code.
	//
	// Returns:
	// - A string representing a randomly generated safe hexadecimal color code.
	GenerateSafeHexColor() string
	// GenerateRgbColorAsArray generates a random RGB color represented as an array of strings.
	//
	// Returns:
	// - An array of strings representing a randomly generated RGB color.
	GenerateRgbColorAsArray() []string

	// GenerateRgbColor generates a random RGB color code.
	//
	// Returns:
	// - A string representing a randomly generated RGB color code.
	GenerateRgbColor() string
	// GenerateRgbCssColor generates a random CSS-compatible RGB color code.
	//
	// Returns:
	// - A string representing a randomly generated CSS-compatible RGB color code.
	GenerateRgbCssColor() string
	// GenerateRgbaCssColor generates a random CSS-compatible RGBA color code.
	//
	// Returns:
	// - A string representing a randomly generated CSS-compatible RGBA color code.
	GenerateRgbaCssColor() string
	// GenerateSafeColorName generates a random safe color name.
	//
	// Returns:
	// - A string representing a randomly generated safe color name.
	GenerateSafeColorName() string
	// GenerateColorName generates a random color name.
	//
	// Returns:
	// - A string representing a randomly generated color name.
	GenerateColorName() string
	GenerateHslColor() string
	GenerateHslColorAsArray() []int
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
	GenerateBoolean(chanceOfGettingTrue ...int) bool
	GenerateMd5() string
	GenerateSha1() string
	GenerateSha256() string
	GenerateLocale() string
	GenerateCountryCode() string
	GenerateCountryISOAlpha3() string
	GenerateLanguageCode() string
	GenerateCurrencyCode() string
	GenerateEmoji() string
	GenerateWord() string
	GenerateWordsAsText(numberOfWords int) string
	GenerateWordsAsList(numberOfWords int) []string
	GenerateSentence(numberOfWords int, variableNumberOfWords bool) string
	GenerateSentencesAsText(numberOfSentences int, variableNumberOfWords bool) string
	GenerateSentencesAsList(numberOfSentences int, variableNumberOfWords bool) []string
	GenerateParagraph(numberOfSentences int, variableNumberOfSentences bool) string
	GenerateParagraphsAsList(numberOfParagraphs ...int) []string
	GenerateParagraphsAsText(numberOfParagraphs ...int) string
	GenerateText(maxNumberOfCharacters ...int) string
	GenerateBloodGroup() string
	GenerateBloodRh() string
	GenerateBloodType() string

	SetDefaultTimeZone(timeZone string)
	GenerateDateTime(max interface{}, timezone string) time.Time
	GenerateDateTimeAD(max interface{}, timezone string) time.Time
	GenerateDateTimeBetween(startDate interface{}, endDate interface{}, timezone string) time.Time
	GenerateDateTimeInterval(dateString string, intervalString string, timezone string) time.Time
	GenerateDateTimeThisWeek(max string, timezone string) time.Time
	GenerateDateTimeThisMonth(max string, timezone string) time.Time
	GenerateDateTimeThisYear(max string, timezone string) time.Time
	GenerateDateTimeThisDecade(max string, timezone string) time.Time
	GenerateDateTimeThisCentury(max string, timezone string) time.Time
	GenerateDate(format string, max interface{}) string
	GenerateTime(format string, max interface{}) string
	GenerateUnixTime(max interface{}) int
	GenerateISO8602(max interface{}) string
	GenerateAmPm(max interface{}) string
	GenerateDayOfMonth(max interface{}) int
	GenerateDayOfWeek(max interface{}) time.Weekday
	GenerateMonth(max interface{}) int
	GenerateMonthName(max interface{}) string
	GenerateYear(max interface{}) string
	GenerateCentury() string
	GenerateTimezone(countryCode string) *time.Location
	//Ean13() string
	//Ean8() string
	//Isbn10() string
	//Isbn13() string
	//Uuid3() string
	//Semver(params ...bool)
	//MimeType() string
	//Extension() string
	//FilePath() string
}
