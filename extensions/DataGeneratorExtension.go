package extensions

import (
	"time"
)

type DataGeneratorExtension interface {
	// GenerateAddress
	//
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
	GenerateHexColor() string
	GenerateSafeHexColor() string
	GenerateRgbColorAsArray() []string
	GenerateRgbColor() string
	GenerateRgbCssColor() string
	GenerateRgbaCssColor() string
	GenerateSafeColorName() string
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
	//Uuid3() string
	//Semver(params ...bool)
	//MimeType() string
	//Extension() string
	//FilePath() string
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
}
