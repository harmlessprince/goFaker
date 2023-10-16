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
	GenerateStreetAddress() string
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
	// GenerateHslColor generates a random HSL color code.
	//
	// Returns:
	// - A string representing a randomly generated HSL color code.
	GenerateHslColor() string
	// GenerateHslColorAsArray generates a random HSL color represented as an array of integers.
	//
	// Returns:
	// - An array of integers representing a randomly generated HSL color.
	GenerateHslColorAsArray() []int
	// GenerateCompany generates a random company name.
	//
	// Returns:
	// - A string representing a randomly generated company name.
	GenerateCompany() string
	// GenerateCompanySuffix generates a random company suffix.
	//
	// Returns:
	// - A string representing a randomly generated company suffix.
	GenerateCompanySuffix() string
	// GenerateJobTitle generates a random job title.
	//
	// Returns:
	// - A string representing a randomly generated job title.

	GenerateJobTitle() string
	// GenerateName generates a random full name.
	//
	// Parameters:
	// - gender (optional): Specify the gender of the generated name (e.g., "male" or "female").
	//
	// Returns:
	// - A string representing a randomly generated full name.
	GenerateName(gender ...string) string
	// GenerateFirstName generates a random first name.
	//
	// Parameters:
	// - gender (optional): Specify the gender of the generated first name (e.g., "male" or "female").
	//
	// Returns:
	// - A string representing a randomly generated first name.
	GenerateFirstName(gender ...string) string
	// GenerateFirstNameMale generates a random male first name.
	//
	// Returns:
	// - A string representing a randomly generated male first name.
	GenerateFirstNameMale() string
	// GenerateFirstNameFemale generates a random female first name.
	//
	// Returns:
	// - A string representing a randomly generated female first name.
	GenerateFirstNameFemale() string
	// GenerateLastName generates a random last name.
	//
	// Returns:
	// - A string representing a randomly generated last name.
	GenerateLastName() string
	// GenerateTitle generates a random job title.
	//
	// Parameters:
	// - gender (optional): Specify the gender of the generated title (e.g., "male" or "female").
	//
	// Returns:
	// - A string representing a randomly generated job title.
	GenerateTitle(gender ...string) string
	// GenerateTitleMale generates a random male job title.
	//
	// Returns:
	// - A string representing a randomly generated male job title.
	GenerateTitleMale() string
	// GenerateTitleFemale generates a random female job title.
	//
	// Returns:
	// - A string representing a randomly generated female job title.
	GenerateTitleFemale() string
	// GeneratePhoneNumber generates a random phone number.
	//
	// Returns:
	// - A string representing a randomly generated phone number.
	GeneratePhoneNumber() string
	// GenerateE164PhoneNumber generates a random E.164 formatted phone number.
	//
	// Returns:
	// - A string representing a randomly generated E.164 formatted phone number.

	GenerateE164PhoneNumber() string
	// GenerateImei generates a random IMEI (International Mobile Equipment Identity) number.
	//
	// Returns:
	// - A string representing a randomly generated IMEI number.
	GenerateImei() string
	// GenerateEmail generates a random email address.
	//
	// Returns:
	// - A string representing a randomly generated email address.

	GenerateEmail() string
	// GenerateSafeEmail generates a random safe email address.
	//
	// Returns:
	// - A string representing a randomly generated safe email address.
	GenerateSafeEmail() string
	// GenerateFreeEmail generates a random free email address.
	//
	// Returns:
	// - A string representing a randomly generated free email address.
	GenerateFreeEmail() string
	// GenerateCompanyEmail generates a random company email address.
	//
	// Returns:
	// - A string representing a randomly generated company email address.
	GenerateCompanyEmail() string
	// GenerateFreeEmailDomain generates a random domain for free email addresses.
	//
	// Returns:
	// - A string representing a randomly generated domain for free email addresses.
	GenerateFreeEmailDomain() string
	// GenerateSafeEmailDomain generates a random safe domain for email addresses.
	//
	// Returns:
	// - A string representing a randomly generated safe domain for email addresses.
	GenerateSafeEmailDomain() string

	// GenerateUserName generates a random username.
	//
	// Returns:
	// - A string representing a randomly generated username.
	GenerateUserName() string
	// GeneratePassword generates a random password.
	//
	// Parameters:
	// - params (optional): Specify the desired length of the generated password.
	//
	// Returns:
	// - A string representing a randomly generated password.
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

	// SetDefaultTimeZone sets the default time zone for the date and time generation.
	//
	// Parameters:
	// - timeZone: The time zone to set as the default for subsequent date and time generation.
	SetDefaultTimeZone(timeZone string)

	// GenerateDateTime generates a random date and time within a specified time frame.
	//
	// Parameters:
	// - max: The maximum date and time (can be a string or time.Time).
	// - timezone: The time zone to use for the generated date and time.
	//
	// Returns:
	// - A time.Time representing a randomly generated date and time.
	GenerateDateTime(max interface{}, timezone string) time.Time

	// GenerateDateTimeAD generates a random date and time within the Common Era (CE) or AD time frame.
	//
	// Parameters:
	// - max: The maximum date and time (can be a string or time.Time).
	// - timezone: The time zone to use for the generated date and time.
	//
	// Returns:
	// - A time.Time representing a randomly generated date and time in the Common Era.
	GenerateDateTimeAD(max interface{}, timezone string) time.Time

	// GenerateDateTimeBetween generates a random date and time within a specified date range.
	//
	// Parameters:
	// - startDate: The start date and time (can be a string or time.Time).
	// - endDate: The end date and time (can be a string or time.Time).
	// - timezone: The time zone to use for the generated date and time.
	//
	// Returns:
	// - A time.Time representing a randomly generated date and time within the specified range.
	GenerateDateTimeBetween(startDate interface{}, endDate interface{}, timezone string) time.Time

	// GenerateDateTimeInterval generates a random date and time at regular intervals from a specified base date.
	//
	// Parameters:
	// - dateString: The base date as a string.
	// - intervalString: The interval as a string (e.g., "1d" for one day).
	// - timezone: The time zone to use for the generated date and time.
	//
	// Returns:
	// - A time.Time representing a randomly generated date and time at specified intervals from the base date.
	GenerateDateTimeInterval(dateString string, intervalString string, timezone string) time.Time

	// GenerateDateTimeThisWeek generates a random date and time within the current week.
	//
	// Parameters:
	// - max: The maximum date and time (can be a string or time.Time).
	// - timezone: The time zone to use for the generated date and time.
	//
	// Returns:
	// - A time.Time representing a randomly generated date and time within the current week.
	GenerateDateTimeThisWeek(max string, timezone string) time.Time

	// GenerateDateTimeThisMonth generates a random date and time within the current month.
	//
	// Parameters:
	// - max: The maximum date and time (can be a string or time.Time).
	// - timezone: The time zone to use for the generated date and time.
	//
	// Returns:
	// - A time.Time representing a randomly generated date and time within the current month.
	GenerateDateTimeThisMonth(max string, timezone string) time.Time

	// GenerateDateTimeThisYear generates a random date and time within the current year.
	//
	// Parameters:
	// - max: The maximum date and time (can be a string or time.Time).
	// - timezone: The time zone to use for the generated date and time.
	//
	// Returns:
	// - A time.Time representing a randomly generated date and time within the current year.
	GenerateDateTimeThisYear(max string, timezone string) time.Time

	// GenerateDateTimeThisDecade generates a random date and time within the current decade.
	//
	// Parameters:
	// - max: The maximum date and time (can be a string or time.Time).
	// - timezone: The time zone to use for the generated date and time.
	//
	// Returns:
	// - A time.Time representing a randomly generated date and time within the current decade.
	GenerateDateTimeThisDecade(max string, timezone string) time.Time

	// GenerateDateTimeThisCentury generates a random date and time within the current century.
	//
	// Parameters:
	// - max: The maximum date and time (can be a string or time.Time).
	// - timezone: The time zone to use for the generated date and time.
	//
	// Returns:
	// - A time.Time representing a randomly generated date and time within the current century.
	GenerateDateTimeThisCentury(max string, timezone string) time.Time

	// GenerateDate generates a random date in the specified format.
	//
	// Parameters:
	// - format: The format of the date (e.g., "2006-01-02").
	// - max: The maximum date (can be a string or time.Time).
	//
	// Returns:
	// - A string representing a randomly generated date.
	GenerateDate(format string, max interface{}) string

	// GenerateTime generates a random time in the specified format.
	//
	// Parameters:
	// - format: The format of the time (e.g., "15:04:05").
	// - max: The maximum time (can be a string or time.Time).
	//
	// Returns:
	// - A string representing a randomly generated time.
	GenerateTime(format string, max interface{}) string

	// GenerateUnixTime generates a random UNIX timestamp.
	//
	// Parameters:
	// - max: The maximum UNIX timestamp.
	//
	// Returns:
	// - An integer representing a randomly generated UNIX timestamp.
	GenerateUnixTime(max interface{}) int

	// GenerateISO8602 generates a random ISO 8601 formatted date and time.
	//
	// Parameters:
	// - max: The maximum date and time (can be a string or time.Time).
	//
	// Returns:
	// - A string representing a randomly generated ISO 8601 formatted date and time.
	GenerateISO8602(max interface{}) string

	// GenerateAmPm generates a random AM or PM.
	//
	// Parameters:
	// - max: The maximum value.
	//
	// Returns:
	// - A string representing "AM" or "PM."
	GenerateAmPm(max interface{}) string

	// GenerateDayOfMonth generates a random day of the month (1-31).
	//
	// Parameters:
	// - max: The maximum day of the month.
	//
	// Returns:
	// - An integer representing a randomly generated day of the month.
	GenerateDayOfMonth(max interface{}) int

	// GenerateDayOfWeek generates a random day of the week.
	//
	// Parameters:
	// - max: The maximum value.
	//
	// Returns:
	// - A time.Weekday representing a randomly generated day of the week.
	GenerateDayOfWeek(max interface{}) time.Weekday

	// GenerateMonth generates a random month (1-12).
	//
	// Parameters:
	// - max: The maximum month.
	//
	// Returns:
	// - An integer representing a randomly generated month.
	GenerateMonth(max interface{}) int

	// GenerateMonthName generates a random month name.
	//
	// Parameters:
	// - max: The maximum value.
	//
	// Returns:
	// - A string representing a randomly generated month name.
	GenerateMonthName(max interface{}) string

	// GenerateYear generates a random year.
	//
	// Parameters:
	// - max: The maximum year.
	//
	// Returns:
	// - A string representing a randomly generated year.
	GenerateYear(max interface{}) string

	// GenerateCentury generates a random century.
	//
	// Returns:
	// - A string representing a randomly generated century.
	GenerateCentury() string

	// GenerateTimezone generates a random time zone based on the given country code.
	//
	// Parameters:
	// - countryCode: The country code used to determine

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
