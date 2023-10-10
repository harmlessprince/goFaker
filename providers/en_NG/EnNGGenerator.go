package en_NG

var addressInstance Address

type EnNGGenerator struct {
}

func (e *EnNGGenerator) GenerateAddress() string {
	//TODO implement me
	return addressInstance.Address.Address()
}

func (e *EnNGGenerator) City() string {
	//TODO implement me
	panic("implement me")
}

func (e *EnNGGenerator) PostCode() string {
	//TODO implement me
	panic("implement me")
}

func (e *EnNGGenerator) StreetName() string {
	//TODO implement me
	panic("implement me")
}

func (e *EnNGGenerator) StreetAddress() string {
	//TODO implement me
	panic("implement me")
}

func (e *EnNGGenerator) BuildingNumber() string {
	//TODO implement me
	panic("implement me")
}

func (e *EnNGGenerator) LocalCoordinates() map[string]float64 {
	//TODO implement me
	panic("implement me")
}

func (e *EnNGGenerator) Longitude() float64 {
	//TODO implement me
	panic("implement me")
}

func (e *EnNGGenerator) Latitude() float64 {
	//TODO implement me
	panic("implement me")
}

func (e *EnNGGenerator) CitySuffix() string {
	//TODO implement me
	panic("implement me")
}

func (e *EnNGGenerator) CityName() string {
	//TODO implement me
	panic("implement me")
}
