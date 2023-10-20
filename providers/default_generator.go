package providers

import "time"

var addressGenerator = NewAddress()
var newPhoneNumber = NewBasePhoneNumber()
var newPerson = NewBasePerson()
var newBaseCompany = NewBaseCompany()
var newBaseInternet = NewBaseInternet()

type DefaultGenerator struct {
	BaseAddress
	BasePerson
	BasePhoneNumber
	BaseCompany
	BaseLorem
	BaseInternet
	BaseMiscellaneous
	BaseMedical
	BaseColor
	BaseDateTime
	BaseUuid
}

func (e *DefaultGenerator) GenerateAddress() string {
	return addressGenerator.Address()
}
func (e *DefaultGenerator) GenerateCountry() string {
	return addressGenerator.Country()
}

func (e *DefaultGenerator) GenerateCity() string {
	return addressGenerator.City()
}

func (e *DefaultGenerator) GeneratePostCode() string {
	return addressGenerator.PostCode()
}

func (e *DefaultGenerator) GenerateStreetName() string {
	return addressGenerator.StreetName()
}

func (e *DefaultGenerator) GenerateStreetAddress() string {
	return addressGenerator.StreetAddress()
}

func (e *DefaultGenerator) GenerateBuildingNumber() string {
	return addressGenerator.BuildingNumber()
}

func (e *DefaultGenerator) GenerateLocalCoordinates() map[string]float64 {
	return addressGenerator.LocalCoordinates()
}

func (e *DefaultGenerator) GenerateLongitude() float64 {
	return addressGenerator.Longitude()
}

func (e *DefaultGenerator) GenerateLatitude() float64 {
	return addressGenerator.Latitude()
}

func (e *DefaultGenerator) GenerateCitySuffix() string {
	return addressGenerator.CitySuffix()
}

func (e *DefaultGenerator) GenerateCityName() string {
	return addressGenerator.CityName()
}

func (e *DefaultGenerator) GenerateName(gender ...string) string {
	return e.BasePerson.Name()
}

func (e *DefaultGenerator) GenerateFirstName(gender ...string) string {
	return newPerson.FirstName()
}

func (e *DefaultGenerator) GenerateFirstNameMale() string {
	return newPerson.FirstNameMale()
}

func (e *DefaultGenerator) GenerateFirstNameFemale() string {
	return newPerson.FirstNameFemale()
}

func (e *DefaultGenerator) GenerateLastName() string {
	return newPerson.LastName()
}

func (e *DefaultGenerator) GenerateTitle(gender ...string) string {
	return newPerson.Title(gender...)
}

func (e *DefaultGenerator) GenerateTitleMale() string {
	return newPerson.TitleMale()
}

func (e *DefaultGenerator) GenerateTitleFemale() string {
	return newPerson.TitleFemale()
}

func (e *DefaultGenerator) GeneratePhoneNumber() string {
	return newPhoneNumber.PhoneNumber()
}

func (e *DefaultGenerator) GenerateE164PhoneNumber() string {
	return newPhoneNumber.E164PhoneNumber()
}

func (e *DefaultGenerator) GenerateImei() string {
	return newPhoneNumber.Imei()
}

func (e *DefaultGenerator) GenerateCompany() string {
	return newBaseCompany.Company()
}

func (e *DefaultGenerator) GenerateCompanySuffix() string {
	return newBaseCompany.CompanySuffix()
}

func (e *DefaultGenerator) GenerateJobTitle() string {
	return newBaseCompany.JobTitle()
}

func (e *DefaultGenerator) GenerateEmail() string {
	return newBaseInternet.Email()
}

func (e *DefaultGenerator) GenerateSafeEmail() string {
	return newBaseInternet.SafeEmail()
}

func (e *DefaultGenerator) GenerateFreeEmail() string {
	return newBaseInternet.FreeEmail()
}

func (e *DefaultGenerator) GenerateCompanyEmail() string {
	return newBaseInternet.CompanyEmail()
}

func (e *DefaultGenerator) GenerateFreeEmailDomain() string {
	return newBaseInternet.FreeEmailDomain()
}

func (e *DefaultGenerator) GenerateSafeEmailDomain() string {
	return newBaseInternet.SafeEmailDomain()
}

func (e *DefaultGenerator) GenerateUserName() string {
	return newBaseInternet.UserName()
}

func (e *DefaultGenerator) GeneratePassword(params ...int) string {
	return newBaseInternet.Password(params...)
}

func (e *DefaultGenerator) GenerateDomainName() string {
	return newBaseInternet.DomainName()
}

func (e *DefaultGenerator) GenerateDomainWord() string {
	return newBaseInternet.DomainWord()
}

func (e *DefaultGenerator) GenerateUrl() string {
	return newBaseInternet.Url()
}

func (e *DefaultGenerator) GenerateTld() string {
	return newBaseInternet.Tld()
}

func (e *DefaultGenerator) GenerateIpv4() string {
	return newBaseInternet.Ipv4()
}

func (e *DefaultGenerator) GenerateIpv6() string {
	return newBaseInternet.Ipv6()
}

func (e *DefaultGenerator) GenerateLocalIpv4() string {
	return newBaseInternet.LocalIpv4()
}

func (e *DefaultGenerator) GenerateMacAddress() string {
	return newBaseInternet.MacAddress()
}

func (e *DefaultGenerator) GenerateTransliterate(inputString string) string {
	return newBaseInternet.Transliterate(inputString)
}

func (e *DefaultGenerator) GenerateToAscii(inputString string) string {
	return newBaseInternet.ToAscii(inputString)
}

func (e *DefaultGenerator) GenerateSlug(numberOfWords int, variableNumberOfWords bool) string {
	return newBaseInternet.Slug(numberOfWords, variableNumberOfWords)
}

func (e *DefaultGenerator) GenerateBoolean(chanceOfGettingTrue ...int) bool {
	return e.BaseMiscellaneous.Boolean(chanceOfGettingTrue...)
}

func (e *DefaultGenerator) GenerateMd5() string {
	return e.BaseMiscellaneous.Md5()
}

func (e *DefaultGenerator) GenerateSha1() string {
	return e.BaseMiscellaneous.Sha1()
}

func (e *DefaultGenerator) GenerateSha256() string {
	return e.BaseMiscellaneous.Sha256()
}

func (e *DefaultGenerator) GenerateLocale() string {
	return e.BaseMiscellaneous.Locale()
}

func (e *DefaultGenerator) GenerateCountryCode() string {
	return e.BaseMiscellaneous.CountryCode()
}

func (e *DefaultGenerator) GenerateCountryISOAlpha3() string {
	return e.BaseMiscellaneous.CountryISOAlpha3()
}

func (e *DefaultGenerator) GenerateLanguageCode() string {
	return e.BaseMiscellaneous.LanguageCode()
}

func (e *DefaultGenerator) GenerateCurrencyCode() string {
	return e.BaseMiscellaneous.CurrencyCode()
}

func (e *DefaultGenerator) GenerateEmoji() string {
	return e.BaseMiscellaneous.Emoji()
}

func (e *DefaultGenerator) GenerateWord() string {
	return e.BaseLorem.Word()
}

func (e *DefaultGenerator) GenerateWordsAsText(numberOfWords int) string {
	return e.BaseLorem.WordsAsText(numberOfWords)
}

func (e *DefaultGenerator) GenerateWordsAsList(numberOfWords int) []string {
	return e.BaseLorem.WordsAsList(numberOfWords)
}

func (e *DefaultGenerator) GenerateSentence(numberOfWords int, variableNumberOfWords bool) string {
	return e.BaseLorem.Sentence(numberOfWords, variableNumberOfWords)
}

func (e *DefaultGenerator) GenerateSentencesAsText(numberOfSentences int, variableNumberOfWords bool) string {
	return e.BaseLorem.SentencesAsText(numberOfSentences, variableNumberOfWords)
}

func (e *DefaultGenerator) GenerateSentencesAsList(numberOfSentences int, variableNumberOfWords bool) []string {
	return e.BaseLorem.SentencesAsList(numberOfSentences, variableNumberOfWords)
}

func (e *DefaultGenerator) GenerateParagraph(numberOfSentences int, variableNumberOfSentences bool) string {
	return e.BaseLorem.Paragraph(numberOfSentences, variableNumberOfSentences)
}

func (e *DefaultGenerator) GenerateParagraphsAsList(numberOfParagraphs ...int) []string {
	return e.BaseLorem.ParagraphsAsList(numberOfParagraphs...)
}

func (e *DefaultGenerator) GenerateParagraphsAsText(numberOfParagraphs ...int) string {
	return e.BaseLorem.ParagraphsAsText(numberOfParagraphs...)
}

func (e *DefaultGenerator) GenerateText(maxNumberOfCharacters ...int) string {
	return e.BaseLorem.Text(maxNumberOfCharacters...)
}

func (e *DefaultGenerator) GenerateBloodGroup() string {
	return e.BaseMedical.BloodGroup()
}

func (e *DefaultGenerator) GenerateBloodRh() string {
	return e.BaseMedical.BloodRh()
}

func (e *DefaultGenerator) GenerateBloodType() string {
	return e.BaseMedical.BloodType()
}

func (e *DefaultGenerator) GenerateHexColor() string {
	return e.HexColor()
}

func (e *DefaultGenerator) GenerateSafeHexColor() string {
	return e.SafeHexColor()
}

func (e *DefaultGenerator) GenerateRgbColorAsArray() []string {
	return e.RgbColorAsArray()
}

func (e *DefaultGenerator) GenerateRgbColor() string {
	return e.RgbColor()
}

func (e *DefaultGenerator) GenerateRgbCssColor() string {
	return e.RgbCssColor()
}

func (e *DefaultGenerator) GenerateRgbaCssColor() string {
	return e.RgbaCssColor()
}

func (e *DefaultGenerator) GenerateSafeColorName() string {
	return e.SafeColorName()
}

func (e *DefaultGenerator) GenerateColorName() string {
	return e.ColorName()
}

func (e *DefaultGenerator) GenerateHslColor() string {
	return e.HslColor()
}

func (e *DefaultGenerator) GenerateHslColorAsArray() []int {
	return e.HslColorAsArray()
}

func (e *DefaultGenerator) GenerateDateTime(max interface{}, timezone string) time.Time {
	return e.DateTime(max, timezone)
}

func (e *DefaultGenerator) GenerateDateTimeAD(max interface{}, timezone string) time.Time {
	return e.DateTimeAD(max, timezone)
}

func (e *DefaultGenerator) GenerateDateTimeBetween(startDate interface{}, endDate interface{}, timezone string) time.Time {
	return e.DateTimeBetween(startDate, endDate, timezone)
}

func (e *DefaultGenerator) GenerateDateTimeInterval(dateString string, intervalString string, timezone string) time.Time {
	return e.DateTimeInterval(dateString, intervalString, timezone)
}

func (e *DefaultGenerator) GenerateDateTimeThisWeek(max string, timezone string) time.Time {
	return e.DateTimeThisWeek(max, timezone)
}

func (e *DefaultGenerator) GenerateDateTimeThisMonth(max string, timezone string) time.Time {
	return e.DateTimeThisMonth(max, timezone)
}

func (e *DefaultGenerator) GenerateDateTimeThisYear(max string, timezone string) time.Time {
	return e.DateTimeThisYear(max, timezone)
}

func (e *DefaultGenerator) GenerateDateTimeThisDecade(max string, timezone string) time.Time {
	return e.DateTimeThisDecade(max, timezone)
}

func (e *DefaultGenerator) GenerateDateTimeThisCentury(max string, timezone string) time.Time {
	return e.DateTimeThisCentury(max, timezone)
}

func (e *DefaultGenerator) GenerateDate(format string, max interface{}) string {
	return e.Date(format, max)
}

func (e *DefaultGenerator) GenerateTime(format string, max interface{}) string {
	return e.Time(format, max)
}

func (e *DefaultGenerator) GenerateUnixTime(max interface{}) int {
	return e.UnixTime(max)
}

func (e *DefaultGenerator) GenerateISO8602(max interface{}) string {
	return e.ISO8602(max)
}

func (e *DefaultGenerator) GenerateAmPm(max interface{}) string {
	return e.AmPm(max)
}

func (e *DefaultGenerator) GenerateDayOfMonth(max interface{}) int {
	return e.DayOfMonth(max)
}

func (e *DefaultGenerator) GenerateDayOfWeek(max interface{}) time.Weekday {
	return e.DayOfWeek(max)
}

func (e *DefaultGenerator) GenerateMonth(max interface{}) int {
	return e.Month(max)
}

func (e *DefaultGenerator) GenerateMonthName(max interface{}) string {
	return e.MonthName(max)
}

func (e *DefaultGenerator) GenerateYear(max interface{}) string {
	return e.Year(max)
}

func (e *DefaultGenerator) GenerateCentury() string {
	return e.Century()
}

func (e *DefaultGenerator) GenerateTimezone(countryCode string) *time.Location {
	return e.Timezone(countryCode)
}
