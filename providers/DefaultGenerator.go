package providers

var addressGenerator = NewAddress()
var newPhoneNumber = NewBasePhoneNumber()
var newPerson = NewBasePerson()
var newBaseCompany = NewBaseCompany()
var newBaseIternet = NewBaseInternet()

type DefaultGenerator struct {
	BaseAddress
	BasePerson
	BasePhoneNumber
	BaseCompany
	BaseLorem
	BaseInternet
	BaseMiscellaneous
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
	return newBaseIternet.Email()
}

func (e *DefaultGenerator) GenerateSafeEmail() string {
	return newBaseIternet.SafeEmail()
}

func (e *DefaultGenerator) GenerateFreeEmail() string {
	return newBaseIternet.FreeEmail()
}

func (e *DefaultGenerator) GenerateCompanyEmail() string {
	return newBaseIternet.CompanyEmail()
}

func (e *DefaultGenerator) GenerateFreeEmailDomain() string {
	return newBaseIternet.FreeEmailDomain()
}

func (e *DefaultGenerator) GenerateSafeEmailDomain() string {
	return newBaseIternet.SafeEmailDomain()
}

func (e *DefaultGenerator) GenerateUserName() string {
	return newBaseIternet.UserName()
}

func (e *DefaultGenerator) GeneratePassword(params ...int) string {
	return newBaseIternet.Password(params...)
}

func (e *DefaultGenerator) GenerateDomainName() string {
	return newBaseIternet.DomainName()
}

func (e *DefaultGenerator) GenerateDomainWord() string {
	return newBaseIternet.DomainWord()
}

func (e *DefaultGenerator) GenerateUrl() string {
	return newBaseIternet.Url()
}

func (e *DefaultGenerator) GenerateTld() string {
	return newBaseIternet.Tld()
}

func (e *DefaultGenerator) GenerateIpv4() string {
	return newBaseIternet.Ipv4()
}

func (e *DefaultGenerator) GenerateIpv6() string {
	return newBaseIternet.Ipv6()
}

func (e *DefaultGenerator) GenerateLocalIpv4() string {
	return newBaseIternet.LocalIpv4()
}

func (e *DefaultGenerator) GenerateMacAddress() string {
	return newBaseIternet.MacAddress()
}

func (e *DefaultGenerator) GenerateTransliterate(inputString string) string {
	return newBaseIternet.Transliterate(inputString)
}

func (e *DefaultGenerator) GenerateToAscii(inputString string) string {
	return newBaseIternet.ToAscii(inputString)
}

func (e *DefaultGenerator) GenerateSlug(numberOfWords int, variableNumberOfWords bool) string {
	return newBaseIternet.Slug(numberOfWords, variableNumberOfWords)
}
