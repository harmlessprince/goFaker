package en_NG

import "github.com/harmlessprince/goFaker/providers"

var newAddress = NewAddress()
var newPhoneNumber = NewPhoneNumber()
var newPerson = NewPerson()

type EnNGGenerator struct {
	providers.DefaultGenerator
	EnNGAddress
	EnNGPerson
	EnNGPhoneNumber
}

func (e *EnNGGenerator) GenerateAddress() string {
	return newAddress.BaseAddress.Address()
}

func (e *EnNGGenerator) GenerateCity() string {
	return newAddress.BaseAddress.City()
}

func (e *EnNGGenerator) GenerateCountry() string {
	return newAddress.Country()
}

func (e *EnNGGenerator) GeneratePostCode() string {
	return newAddress.BaseAddress.PostCode()
}

func (e *EnNGGenerator) GenerateCityName() string {
	return newAddress.BaseAddress.CityName()
}

func (e *EnNGGenerator) GenerateName(gender ...string) string {
	return newPerson.BasePerson.Name()
}

func (e *EnNGGenerator) GenerateFirstName(gender ...string) string {
	return newPerson.BasePerson.FirstName(gender...)
}

func (e *EnNGGenerator) GenerateFirstNameMale() string {
	return newPerson.BasePerson.FirstNameMale()
}

func (e *EnNGGenerator) GenerateFirstNameFemale() string {
	return newPerson.BasePerson.FirstNameFemale()
}

func (e *EnNGGenerator) GenerateLastName() string {
	return newPerson.BasePerson.LastName()
}

func (e *EnNGGenerator) GeneratePhoneNumber() string {
	return newPhoneNumber.BasePhoneNumber.PhoneNumber()
}
