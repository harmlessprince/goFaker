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

func (e *EnNGGenerator) GenerateStreetName() string {
	return newAddress.BaseAddress.StreetName()
}

func (e *EnNGGenerator) GenerateStreetAddress() string {
	return newAddress.BaseAddress.StreetAddress()
}

func (e *EnNGGenerator) GenerateBuildingNumber() string {
	return newAddress.BaseAddress.BuildingNumber()
}

func (e *EnNGGenerator) GenerateLocalCoordinates() map[string]float64 {
	return newAddress.BaseAddress.LocalCoordinates()
}

func (e *EnNGGenerator) GenerateLongitude() float64 {
	return newAddress.BaseAddress.Longitude()
}

func (e *EnNGGenerator) GenerateLatitude() float64 {
	return newAddress.BaseAddress.Latitude()
}

func (e *EnNGGenerator) GenerateCitySuffix() string {
	return newAddress.BaseAddress.CitySuffix()
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

func (e *EnNGGenerator) GenerateTitle(gender ...string) string {
	return e.EnNGPerson.Title(gender...)
}

func (e *EnNGGenerator) GenerateTitleMale() string {
	return newPerson.BasePerson.TitleMale()
}

func (e *EnNGGenerator) GenerateTitleFemale() string {
	return newPerson.BasePerson.TitleFemale()
}

func (e *EnNGGenerator) GeneratePhoneNumber() string {
	return newPhoneNumber.BasePhoneNumber.PhoneNumber()
}

func (e *EnNGGenerator) GenerateE164PhoneNumber() string {
	return newPhoneNumber.BasePhoneNumber.E164PhoneNumber()
}

func (e *EnNGGenerator) GenerateImei() string {
	return newPhoneNumber.BasePhoneNumber.Imei()
}
