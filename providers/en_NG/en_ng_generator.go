package en_NG

import "github.com/harmlessprince/goFaker/providers"

var newAddress = NewAddress()
var newPhoneNumber = NewPhoneNumber()
var newPerson = NewEnNGPerson()

type EnNGGenerator struct {
	providers.DefaultGenerator
	EnNGAddress
	EnNGPerson
	EnNGPhoneNumber
}

func (e *EnNGGenerator) Address() (string, error) {
	return newAddress.BaseAddress.Address()
}

func (e *EnNGGenerator) StreetAddress() (string, error) {
	return newAddress.BaseAddress.StreetAddress()
}

func (e *EnNGGenerator) City() (string, error) {
	return newAddress.BaseAddress.City()
}

func (e *EnNGGenerator) Country() (string, error) {
	return newAddress.Country()
}

func (e *EnNGGenerator) PostCode() (string, error) {
	return newAddress.BaseAddress.PostCode()
}

func (e *EnNGGenerator) CityName() (string, error) {
	return newAddress.BaseAddress.CityName()
}

func (e *EnNGGenerator) Name(gender ...string) (string, error) {
	return newPerson.BasePerson.Name(gender...)
}

func (e *EnNGGenerator) FirstName(gender ...string) (string, error) {
	return newPerson.BasePerson.FirstName(gender...)
}

func (e *EnNGGenerator) FirstNameMale() (string, error) {
	return newPerson.BasePerson.FirstNameMale()
}

func (e *EnNGGenerator) FirstNameFemale() (string, error) {
	return newPerson.BasePerson.FirstNameFemale()
}

func (e *EnNGGenerator) LastName() (string, error) {
	return newPerson.BasePerson.LastName()
}

func (e *EnNGGenerator) PhoneNumber() (string, error) {
	return newPhoneNumber.BasePhoneNumber.PhoneNumber()
}
