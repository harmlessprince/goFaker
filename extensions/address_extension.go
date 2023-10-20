package extensions

type AddressExtension interface {
	Address() string
	City() string
	PostCode() string
	StreetName() string
	StreetAddress() string
	BuildingNumber() string
	LocalCoordinates() map[string]float64
	Longitude() float64
	Latitude() float64
	CitySuffix() string
	CityName() string
}
