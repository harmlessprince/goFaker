package providers

var addressGenerator = NewAddress()

type DefaultGenerator struct {
	BaseAddress
	BasePerson
	BasePhoneNumber
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
	return e.BasePerson.FirstName()
}

func (e *DefaultGenerator) GenerateFirstNameMale() string {
	return e.BasePerson.FirstNameMale()
}

func (e *DefaultGenerator) GenerateFirstNameFemale() string {
	return e.BasePerson.FirstNameFemale()
}

func (e *DefaultGenerator) GenerateLastName() string {
	return e.BasePerson.LastName()
}

func (e *DefaultGenerator) GenerateTitle(gender ...string) string {
	return e.BasePerson.Title(gender...)
}

func (e *DefaultGenerator) GenerateTitleMale() string {
	return e.BasePerson.TitleMale()
}

func (e *DefaultGenerator) GenerateTitleFemale() string {
	return e.BasePerson.TitleFemale()
}

func (e *DefaultGenerator) GeneratePhoneNumber() string {
	return e.BasePhoneNumber.PhoneNumber()
}

func (e *DefaultGenerator) GenerateE164PhoneNumber() string {
	return e.BasePhoneNumber.E164PhoneNumber()
}

func (e *DefaultGenerator) GenerateImei() string {
	return e.BasePhoneNumber.Imei()
}
