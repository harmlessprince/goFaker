package en_NG

var newAddress = NewAddress()

var newPhoneNumber = NewPhoneNumber()
var newPerson = NewPerson()

type EnNGGenerator struct {
	NGAddress
	EnNGPerson
	EnNGPhoneNumber
}

func (e *EnNGGenerator) GenerateAddress() string {
	return newAddress.Address.Address()
}

func (e *EnNGGenerator) GenerateCity() string {
	return newAddress.Address.City()
}

func (e *EnNGGenerator) GeneratePostCode() string {
	return newAddress.Address.PostCode()
}

func (e *EnNGGenerator) GenerateStreetName() string {
	return newAddress.Address.StreetName()
}

func (e *EnNGGenerator) GenerateStreetAddress() string {
	return newAddress.Address.StreetAddress()
}

func (e *EnNGGenerator) GenerateBuildingNumber() string {
	return newAddress.Address.BuildingNumber()
}

func (e *EnNGGenerator) GenerateLocalCoordinates() map[string]float64 {
	return newAddress.Address.LocalCoordinates()
}

func (e *EnNGGenerator) GenerateLongitude() float64 {
	return newAddress.Address.Longitude()
}

func (e *EnNGGenerator) GenerateLatitude() float64 {
	return newAddress.Address.Latitude()
}

func (e *EnNGGenerator) GenerateCitySuffix() string {
	return newAddress.Address.CitySuffix()
}

func (e *EnNGGenerator) GenerateCityName() string {
	return newAddress.Address.CityName()
}

func (e *EnNGGenerator) GenerateName(gender ...string) string {
	return newPerson.Person.Name()
}

func (e *EnNGGenerator) GenerateFirstName(gender ...string) string {
	return newPerson.Person.FirstName(gender...)
}

func (e *EnNGGenerator) GenerateFirstNameMale() string {
	return newPerson.Person.FirstNameMale()
}

func (e *EnNGGenerator) GenerateFirstNameFemale() string {
	return newPerson.Person.FirstNameFemale()
}

func (e *EnNGGenerator) GenerateLastName() string {
	return newPerson.Person.LastName()
}

func (e *EnNGGenerator) GenerateTitle(gender ...string) string {
	return e.EnNGPerson.Title(gender...)
}

func (e *EnNGGenerator) GenerateTitleMale() string {
	return newPerson.Person.TitleMale()
}

func (e *EnNGGenerator) GenerateTitleFemale() string {
	return newPerson.Person.TitleFemale()
}

func (e *EnNGGenerator) GeneratePhoneNumber() string {
	return newPhoneNumber.PhoneNumber.PhoneNumber()
}

func (e *EnNGGenerator) GenerateE164PhoneNumber() string {
	return newPhoneNumber.PhoneNumber.E164PhoneNumber()
}

func (e *EnNGGenerator) GenerateImei() string {
	return newPhoneNumber.PhoneNumber.Imei()
}
