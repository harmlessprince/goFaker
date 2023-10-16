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

	// GenerateDomainName generates a random domain name.
	//
	// Returns:
	// - A string representing a randomly generated domain name.
	GenerateDomainName() string

	// GenerateDomainWord generates a random word suitable for use in a domain name.
	//
	// Returns:
	// - A string representing a randomly generated word for a domain name.
	GenerateDomainWord() string

	// GenerateUrl generates a random URL.
	//
	// Returns:
	// - A string representing a randomly generated URL.
	GenerateUrl() string

	// GenerateTld generates a random top-level domain (TLD) name.
	//
	// Returns:
	// - A string representing a randomly generated top-level domain name.
	GenerateTld() string

	// GenerateIpv4 generates a random IPv4 address.
	//
	// Returns:
	// - A string representing a randomly generated IPv4 address.
	GenerateIpv4() string

	// GenerateIpv6 generates a random IPv6 address.
	//
	// Returns:
	// - A string representing a randomly generated IPv6 address.
	GenerateIpv6() string

	// GenerateLocalIpv4 generates a random local IPv4 address.
	//
	// Returns:
	// - A string representing a randomly generated local IPv4 address.
	GenerateLocalIpv4() string

	// GenerateMacAddress generates a random MAC address.
	//
	// Returns:
	// - A string representing a randomly generated MAC address.
	GenerateMacAddress() string

	// GenerateTransliterate transliterates a given input string.
	//
	// Parameters:
	// - inputString: The input string to be transliterated.
	//
	// Returns:
	// - A string representing the transliterated version of the input string.
	GenerateTransliterate(inputString string) string

	// GenerateToAscii converts a given input string to its ASCII representation.
	//
	// Parameters:
	// - inputString: The input string to be converted to ASCII.
	//
	// Returns:
	// - A string representing the ASCII representation of the input string.
	GenerateToAscii(inputString string) string

	// GenerateSlug generates a random URL-friendly slug.
	//
	// Parameters:
	// - numberOfWords: The number of words in the slug.
	// - variableNumberOfWords: Specify if the number of words can vary.
	//
	// Returns:
	// - A string representing a randomly generated slug.
	GenerateSlug(numberOfWords int, variableNumberOfWords bool) string

	// GenerateBoolean generates a random boolean value.
	//
	// Parameters:
	// - chanceOfGettingTrue (optional): Specify the probability of getting 'true' as a percentage (default is 50%).
	//
	// Returns:
	// - A boolean value, either true or false.
	GenerateBoolean(chanceOfGettingTrue ...int) bool

	// GenerateMd5 generates a random MD5 hash.
	//
	// Returns:
	// - A string representing a randomly generated MD5 hash.
	GenerateMd5() string

	// GenerateSha1 generates a random SHA-1 hash.
	//
	// Returns:
	// - A string representing a randomly generated SHA-1 hash.
	GenerateSha1() string

	// GenerateSha256 generates a random SHA-256 hash.
	//
	// Returns:
	// - A string representing a randomly generated SHA-256 hash.
	GenerateSha256() string

	// GenerateLocale generates a random locale code.
	//
	// Returns:
	// - A string representing a randomly generated locale code.
	GenerateLocale() string

	// GenerateCountryCode generates a random country code.
	//
	// Returns:
	// - A string representing a randomly generated country code.
	GenerateCountryCode() string

	// GenerateCountryISOAlpha3 generates a random ISO 3166-1 alpha-3 country code.
	//
	// Returns:
	// - A string representing a randomly generated ISO 3166-1 alpha-3 country code.
	GenerateCountryISOAlpha3() string

	// GenerateLanguageCode generates a random language code.
	//
	// Returns:
	// - A string representing a randomly generated language code.
	GenerateLanguageCode() string

	// GenerateCurrencyCode generates a random currency code.
	//
	// Returns:
	// - A string representing a randomly generated currency code.
	GenerateCurrencyCode() string

	// GenerateEmoji generates a random emoji.
	//
	// Returns:
	// - A string representing a randomly generated emoji.
	GenerateEmoji() string

	// GenerateWord generates a random word.
	//
	// Returns:
	// - A string representing a randomly generated word.
	GenerateWord() string

	// GenerateWordsAsText generates a string of random words.
	//
	// Parameters:
	// - numberOfWords: The number of words to generate.
	//
	// Returns:
	// - A string representing a space-separated list of randomly generated words.
	GenerateWordsAsText(numberOfWords int) string

	// GenerateWordsAsList generates a list of random words.
	//
	// Parameters:
	// - numberOfWords: The number of words to generate.
	//
	// Returns:
	// - A slice of strings representing randomly generated words.
	GenerateWordsAsList(numberOfWords int) []string

	// GenerateSentence generates a random sentence.
	//
	// Parameters:
	// - numberOfWords: The number of words in the sentence.
	// - variableNumberOfWords: Specify if the number of words can vary.
	//
	// Returns:
	// - A string representing a randomly generated sentence.
	GenerateSentence(numberOfWords int, variableNumberOfWords bool) string

	// GenerateSentencesAsText generates a string of random sentences.
	//
	// Parameters:
	// - numberOfSentences: The number of sentences to generate.
	// - variableNumberOfWords: Specify if the number of words can vary in each sentence.
	//
	// Returns:
	// - A string representing a space-separated list of randomly generated sentences.
	GenerateSentencesAsText(numberOfSentences int, variableNumberOfWords bool) string

	// GenerateSentencesAsList generates a list of random sentences.
	//
	// Parameters:
	// - numberOfSentences: The number of sentences to generate.
	// - variableNumberOfWords: Specify if the number of words can vary in each sentence.
	//
	// Returns:
	// - A slice of strings representing randomly generated sentences.
	GenerateSentencesAsList(numberOfSentences int, variableNumberOfWords bool) []string

	// GenerateParagraph generates a random paragraph.
	//
	// Parameters:
	// - numberOfSentences: The number of sentences in the paragraph.
	// - variableNumberOfSentences: Specify if the number of sentences can vary.
	//
	// Returns:
	// - A string representing a randomly generated paragraph.
	GenerateParagraph(numberOfSentences int, variableNumberOfSentences bool) string

	// GenerateParagraphsAsList generates a list of random paragraphs.
	//
	// Parameters:
	// - numberOfParagraphs: The number of paragraphs to generate.
	//
	// Returns:
	// - A slice of strings representing randomly generated paragraphs.
	GenerateParagraphsAsList(numberOfParagraphs ...int) []string

	// GenerateParagraphsAsText generates a string of random paragraphs.
	//
	// Parameters:
	// - numberOfParagraphs: The number of paragraphs to generate.
	//
	// Returns:
	// - A string representing a space-separated list of randomly generated paragraphs.
	GenerateParagraphsAsText(numberOfParagraphs ...int) string

	// GenerateText generates a random text with a specified maximum number of characters.
	//
	// Parameters:
	// - maxNumberOfCharacters (optional): The maximum number of characters for the generated text.
	//
	// Returns:
	// - A string representing randomly generated text.
	GenerateText(maxNumberOfCharacters ...int) string

	// GenerateBloodGroup generates a random blood group, including the ABO blood type (A, B, AB, O) and the Rh factor (+ or -).
	//
	// Returns:
	// - A string representing a randomly generated blood group, such as "A+", "B-", "AB+", "O-", etc.
	GenerateBloodGroup() string
	// GenerateBloodRh generates a random blood Rh factor (+ or -).
	//
	// Returns:
	// - A string representing a randomly generated blood Rh factor, either "+" or "-".
	GenerateBloodRh() string

	// GenerateBloodType generates a random blood type (A, B, AB, O).
	//
	// Returns:
	// - A string representing a randomly generated blood type (A, B, AB, O).
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
