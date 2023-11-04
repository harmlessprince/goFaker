package providers

import (
	"fmt"
)

var addressGenerator = NewBaseAddress()
var newPhoneNumber = NewBasePhoneNumber()
var newBasePerson = NewBasePerson()
var newBaseCompany = NewBaseCompany()
var newBaseInternet = NewBaseInternet()

type DataGeneratorExtension interface {
	AddressInterface
	PersonInterface
	DateTimeInterface
	ColorInterface
	CompanyInterface
	ImageInterface
	InternetInterface
	LoremInterface
	BloodInterface
	MiscellaneousInterface
	PaymentInterface
	PhoneNumberInterface
	TextInterface
	UserAgentInterface
	UuidInterface
	NumberInterface
	BarcodeInterface
}
type DefaultGenerator struct {
	BaseAddress
	BaseBarcode
	BaseColor
	BaseCompany
	BaseDateTime
	BaseImage
	BaseInternet
	BaseLorem
	BaseMedical
	BaseMiscellaneous
	BasePayment
	BasePerson
	BasePhoneNumber
	BaseProvider
	BaseUuid
	BaseUserAgent
	BaseText
}

// Address generates a random address string.
//
// This method generates a random address string that can be used for various purposes such as testing, example data, or placeholders.
//
// Returns:
// - A string representing a randomly generated address.
func (e *DefaultGenerator) Address() (string, error) {
	return addressGenerator.Address()
}
func (e *DefaultGenerator) Country() (string, error) {
	return addressGenerator.Country()
}

func (e *DefaultGenerator) City() (string, error) {
	return addressGenerator.City()
}

func (e *DefaultGenerator) PostCode() (string, error) {
	return addressGenerator.PostCode()
}

func (e *DefaultGenerator) StreetName() (string, error) {
	return addressGenerator.StreetName()
}

func (e *DefaultGenerator) StreetAddress() (string, error) {
	return addressGenerator.StreetAddress()
}

func (e *DefaultGenerator) BuildingNumber() (string, error) {
	return addressGenerator.BuildingNumber()
}

func (e *DefaultGenerator) LocalCoordinates() (map[string]float64, error) {
	return addressGenerator.LocalCoordinates()
}

func (e *DefaultGenerator) Longitude() (float64, error) {
	return addressGenerator.Longitude()
}

func (e *DefaultGenerator) Latitude() (float64, error) {
	return addressGenerator.Latitude()
}

func (e *DefaultGenerator) CitySuffix() (string, error) {
	return addressGenerator.CitySuffix()
}

func (e *DefaultGenerator) CityName() (string, error) {
	return addressGenerator.CityName()
}

func (e *DefaultGenerator) Name(gender ...string) (string, error) {
	return newBasePerson.Name(gender...)
}

func (e *DefaultGenerator) FirstName(gender ...string) (string, error) {
	fmt.Println(newBasePerson)
	return newBasePerson.FirstName(gender...)
}

func (e *DefaultGenerator) FirstNameMale() (string, error) {
	return newBasePerson.FirstNameMale()
}

func (e *DefaultGenerator) FirstNameFemale() (string, error) {
	return newBasePerson.FirstNameFemale()
}

func (e *DefaultGenerator) LastName() (string, error) {
	return newBasePerson.LastName()
}

func (e *DefaultGenerator) Title(gender ...string) (string, error) {
	return newBasePerson.Title(gender...)
}

func (e *DefaultGenerator) TitleMale() (string, error) {
	return newBasePerson.TitleMale()
}

func (e *DefaultGenerator) TitleFemale() (string, error) {
	return newBasePerson.TitleFemale()
}

func (e *DefaultGenerator) PhoneNumber() (string, error) {
	return newPhoneNumber.PhoneNumber()
}

func (e *DefaultGenerator) E164PhoneNumber() (string, error) {
	return newPhoneNumber.E164PhoneNumber()
}

func (e *DefaultGenerator) Imei() (string, error) {
	return newPhoneNumber.Imei()
}

func (e *DefaultGenerator) Company() (string, error) {
	return newBaseCompany.Company()
}

func (e *DefaultGenerator) CompanySuffix() (string, error) {
	return newBaseCompany.CompanySuffix()
}

func (e *DefaultGenerator) JobTitle() (string, error) {
	return newBaseCompany.JobTitle()
}

func (e *DefaultGenerator) Email() (string, error) {
	return newBaseInternet.Email()
}

func (e *DefaultGenerator) SafeEmail() (string, error) {
	return newBaseInternet.SafeEmail()
}

func (e *DefaultGenerator) FreeEmail() (string, error) {
	return newBaseInternet.FreeEmail()
}

func (e *DefaultGenerator) CompanyEmail() (string, error) {
	return newBaseInternet.CompanyEmail()
}

func (e *DefaultGenerator) FreeEmailDomain() (string, error) {
	return newBaseInternet.FreeEmailDomain()
}

func (e *DefaultGenerator) SafeEmailDomain() (string, error) {
	return newBaseInternet.SafeEmailDomain()
}

func (e *DefaultGenerator) UserName() (string, error) {
	return newBaseInternet.UserName()
}

func (e *DefaultGenerator) Password(params ...int) (string, error) {
	return newBaseInternet.Password(params...)
}

func (e *DefaultGenerator) DomainName() (string, error) {
	return newBaseInternet.DomainName()
}

func (e *DefaultGenerator) DomainWord() (string, error) {
	return newBaseInternet.DomainWord()
}

func (e *DefaultGenerator) Url() (string, error) {
	return newBaseInternet.Url()
}

func (e *DefaultGenerator) Tld() (string, error) {
	return newBaseInternet.Tld()
}

func (e *DefaultGenerator) Ipv4() (string, error) {
	return newBaseInternet.Ipv4()
}

func (e *DefaultGenerator) Ipv6() (string, error) {
	return newBaseInternet.Ipv6()
}

func (e *DefaultGenerator) LocalIpv4() (string, error) {
	return newBaseInternet.LocalIpv4()
}

func (e *DefaultGenerator) MacAddress() (string, error) {
	return newBaseInternet.MacAddress()
}

func (e *DefaultGenerator) Transliterate(inputString string) (string, error) {
	return newBaseInternet.Transliterate(inputString)
}

func (e *DefaultGenerator) ToAscii(inputString string) (string, error) {
	return newBaseInternet.ToAscii(inputString)
}

func (e *DefaultGenerator) Slug(numberOfWords int, variableNumberOfWords bool) (string, error) {
	return newBaseInternet.Slug(numberOfWords, variableNumberOfWords)
}
