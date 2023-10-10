package providers

type DefaultGenerator struct {
	Address
	Person
	PhoneNumber
}

func (e *DefaultGenerator) GenerateAddress() string {
	//TODO implement me
	return e.Address.Address()
}

func (e *DefaultGenerator) City() string {
	//TODO implement me
	panic("implement me")
}

func (e *DefaultGenerator) PostCode() string {
	//TODO implement me
	panic("implement me")
}

func (e *DefaultGenerator) StreetName() string {
	//TODO implement me
	panic("implement me")
}

func (e *DefaultGenerator) StreetAddress() string {
	//TODO implement me
	panic("implement me")
}

func (e *DefaultGenerator) BuildingNumber() string {
	//TODO implement me
	panic("implement me")
}

func (e *DefaultGenerator) LocalCoordinates() map[string]float64 {
	//TODO implement me
	panic("implement me")
}

func (e *DefaultGenerator) Longitude() float64 {
	//TODO implement me
	panic("implement me")
}

func (e *DefaultGenerator) Latitude() float64 {
	//TODO implement me
	panic("implement me")
}

func (e *DefaultGenerator) CitySuffix() string {
	//TODO implement me
	panic("implement me")
}

func (e *DefaultGenerator) CityName() string {
	//TODO implement me
	panic("implement me")
}
