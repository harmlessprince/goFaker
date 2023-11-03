package providers

import (
	"errors"
	"math/rand"
	"strings"
	"time"
)

type AddressInterface interface {
	Address() (string, error)
	City() (string, error)
	PostCode() (string, error)
	StreetName() (string, error)
	StreetAddress() (string, error)
	BuildingNumber() (string, error)
	LocalCoordinates() (map[string]float64, error)
	Longitude() (float64, error)
	Latitude() (float64, error)
	CitySuffix() (string, error)
	CityName() (string, error)
}
type BaseAddress struct {
	BaseProvider
	cityFormats          []string
	streetNameFormats    []string
	streetAddressFormats []string
	addressFormats       []string
	postcode             []string
	country              []string
	buildingNumber       []string
	citySuffix           []string
	streetSuffix         []string
	cityPrefix           []string
	streetPrefix         []string
}

func NewBaseAddress() *BaseAddress {
	address := &BaseAddress{}
	address.SetStreetSuffix()
	address.SetPostCodes()
	address.SetBuildingNumber()
	address.SetCitySuffix()
	address.SetCityPrefix()
	address.SetCountry()
	address.SetCityFormats()
	address.SetStreetNameFormats()
	address.SetStreetAddressFormats()
	address.SetAddressFormats()
	return address
}

//<=============Start Setters=============>//

func (a *BaseAddress) SetStreetSuffix(param ...[]string) {

	if len(param) > 0 {
		a.streetSuffix = param[0]
	} else {
		a.streetSuffix = []string{
			"Street", "Avenue", "Way", "Drive", "Road",
		}
	}
}
func (a *BaseAddress) SetPostCodes(param ...[]string) {
	if len(param) > 0 {
		a.postcode = param[0]
	} else {
		a.postcode = []string{
			"#####",
		}
	}
}
func (a *BaseAddress) SetBuildingNumber(param ...[]string) {

	if len(param) > 0 {
		a.buildingNumber = param[0]
	} else {
		a.buildingNumber = []string{
			"%#",
		}
	}
}
func (a *BaseAddress) SetCitySuffix(param ...[]string) {
	if len(param) > 0 {
		a.citySuffix = param[0]
	} else {
		a.citySuffix = []string{
			"Ville",
		}
	}
}
func (a *BaseAddress) SetCityPrefix(param ...[]string) {
	if len(param) > 0 {
		a.cityPrefix = param[0]
	} else {
		a.cityPrefix = []string{
			"Clanton", "Jasper", "GreenVille",
		}
	}
}
func (a *BaseAddress) SetStreetPrefix(param ...[]string) {
	if len(param) > 0 {
		a.streetPrefix = param[0]
	} else {
		a.streetPrefix = []string{
			"Decator", "Demopolis", "Florence",
		}
	}
}
func (a *BaseAddress) SetCountry(param ...[]string) {
	if len(param) > 0 {
		a.country = param[0]
	} else {
		a.country = []string{
			"Nigeria",
		}
	}

}

func (a *BaseAddress) SetCityFormats(param ...[]string) {
	if len(param) > 0 {
		a.cityFormats = param[0]
	} else {
		a.cityFormats = []string{
			"{{CityName}} {{CitySuffix}}",
		}
	}

}
func (a *BaseAddress) SetStreetNameFormats(param ...[]string) {
	if len(param) > 0 {
		a.streetNameFormats = param[0]
	} else {
		a.streetNameFormats = []string{
			"{{StreetPrefix}} {{StreetSuffix}}",
		}
	}

}
func (a *BaseAddress) SetStreetAddressFormats(param ...[]string) {
	if len(param) > 0 {
		a.streetAddressFormats = param[0]
	} else {
		a.streetAddressFormats = []string{
			"{{BuildingNumber}}, {{StreetName}}",
		}
	}

}

func (a *BaseAddress) SetAddressFormats(param ...[]string) {
	if len(param) > 0 {
		a.addressFormats = param[0]
	} else {
		a.addressFormats = []string{
			"{{StreetAddress}} {{PostCode}} {{City}}",
		}
	}

}

//<=============End Setters=============>//

//<=============Start Getters=============>//

func (a *BaseAddress) GetStreetSuffix() []string {
	return a.streetSuffix
}

func (a *BaseAddress) GetPostCodes() []string {
	return a.postcode
}
func (a *BaseAddress) GetBuildingNumber() []string {
	return a.buildingNumber
}
func (a *BaseAddress) GetCitySuffix() []string {
	return a.citySuffix
}
func (a *BaseAddress) GetCityPrefix() []string {
	return a.cityPrefix
}
func (a *BaseAddress) GetStreetPrefix() []string {
	return a.streetPrefix
}

func (a *BaseAddress) GetCountry() []string {
	return a.country
}

// GetCityFormats is a method that provides access to the cityFormats.
func (a *BaseAddress) GetCityFormats() []string {
	return a.cityFormats
}

// GetStreetNameFormats is a method that provides access to the streetNameFormats.
func (a *BaseAddress) GetStreetNameFormats() []string {

	return a.streetNameFormats
}

// GetStreetAddressFormats is a method that provides access to the streetAddressFormats.
func (a *BaseAddress) GetStreetAddressFormats() []string {

	return a.streetAddressFormats
}

// GetAddressFormats is a method that provides access to the addressFormats.
func (a *BaseAddress) GetAddressFormats() []string {
	return a.addressFormats
}

//<=============End Getters=============>//

//<=============Start Implementations=============>//

func (a *BaseAddress) CityName() (string, error) {
	cityPrefix, err := a.BaseProvider.RandomElementFromStringSlice(a.GetCityPrefix())
	if err != nil {
		return "", errors.New(err.Error())
	}
	return cityPrefix, nil
}
func (a *BaseAddress) CityPrefix() (string, error) {
	citySuffix, err := a.BaseProvider.RandomElementFromStringSlice(a.GetCitySuffix())
	if err != nil {
		return "", errors.New(err.Error())
	}
	return citySuffix, nil
}
func (a *BaseAddress) StreetPrefix() (string, error) {
	streetPrefix, err := a.BaseProvider.RandomElementFromStringSlice(a.GetStreetPrefix())
	if err != nil {
		return "", errors.New(err.Error())
	}
	return streetPrefix, err
}
func (a *BaseAddress) CitySuffix() (string, error) {
	citySuffix, err := a.BaseProvider.RandomElementFromStringSlice(a.GetCitySuffix())
	if err != nil {
		return "", errors.New(err.Error())
	}
	return citySuffix, nil
}

func (a *BaseAddress) StreetSuffix() (string, error) {
	citySuffix, err := a.BaseProvider.RandomElementFromStringSlice(a.GetStreetSuffix())
	if err != nil {
		return "", errors.New(err.Error())
	}
	return citySuffix, nil
}

func (a *BaseAddress) Country() (string, error) {
	country, err := a.BaseProvider.RandomElementFromStringSlice(a.GetCountry())
	if err != nil {
		return "", errors.New(err.Error())
	}
	return country, nil
}
func (a *BaseAddress) Address() (string, error) {
	formats := a.GetAddressFormats()
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := rand.Intn(len(formats))
	address, err := a.BaseProvider.Parse(formats[randomIndex], a)
	if err != nil {
		return "", errors.New(err.Error())
	}
	return address, nil
}

func (a *BaseAddress) City() (string, error) {
	formats := a.GetCityFormats()
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := rand.Intn(len(formats))
	city, err := a.BaseProvider.Parse(formats[randomIndex], a)
	if err != nil {
		return "", errors.New(err.Error())
	}
	return city, nil
}

func (a *BaseAddress) PostCode() (string, error) {
	randomCode, err := a.BaseProvider.RandomElementFromStringSlice(a.GetPostCodes())
	if err != nil {
		return "", errors.New(err.Error())
	}
	postCode, err := a.BaseProvider.Bothify(randomCode)
	if err != nil {
		return "", errors.New(err.Error())
	}
	return a.BaseProvider.ToLower(postCode), nil
}

func (a *BaseAddress) StreetName() (string, error) {
	formats := a.GetStreetNameFormats()
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := rand.Intn(len(formats))
	streetName, err := a.BaseProvider.Parse(formats[randomIndex], a)
	if err != nil {
		return "", errors.New(err.Error())
	}
	return streetName, nil
}

func (a *BaseAddress) StreetAddress() (string, error) {
	formats := a.GetStreetAddressFormats()
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := rand.Intn(len(formats))
	streetAddress, err := a.BaseProvider.Parse(formats[randomIndex], a)
	if err != nil {
		return "", errors.New(err.Error())
	}
	return streetAddress, nil
}

func (a *BaseAddress) BuildingNumber() (string, error) {
	randomBuildNumber, err := a.BaseProvider.RandomElementFromStringSlice(a.GetBuildingNumber())
	if err != nil {
		return "", errors.New(err.Error())
	}
	generatedBuildNumber, err := a.BaseProvider.Numerify(randomBuildNumber)
	if err != nil {
		return "", errors.New(err.Error())
	}

	return generatedBuildNumber, err
}

// LocalCoordinates returns the latitude and longitude of an address as a map.
// It extracts the latitude and longitude values from the address struct and
// returns them in a map with keys "latitude" and "longitude".
// The returned map has float64 values.
//
// Example usage:
//
//	coordinates := address.LocalCoordinates()
//	fmt.Println(coordinates["latitude"], coordinates["longitude"])
//
// Returns:
//
//	map[string]float64{"latitude": 40.7128, "longitude": -74.0060}
func (a *BaseAddress) LocalCoordinates() (map[string]float64, error) {
	latitude, _ := a.Latitude()
	longitude, _ := a.Longitude()
	coordinate := map[string]float64{
		"latitude":  latitude,
		"longitude": longitude,
	}
	return coordinate, nil
}

func (a *BaseAddress) Longitude() (float64, error) {
	options := RandomFloatOptions{
		NumberOfMaxDecimals: 6,
		Min:                 -180,
		Max:                 180,
	}
	long, err := a.BaseProvider.RandomFloat(options)
	if err != nil {
		return 0, errors.New(err.Error())
	}
	return long, err
}

func (a *BaseAddress) Latitude() (float64, error) {
	options := RandomFloatOptions{
		NumberOfMaxDecimals: 6,
		Min:                 -90,
		Max:                 90,
	}
	lat, err := a.BaseProvider.RandomFloat(options)
	if err != nil {
		return 0, errors.New(err.Error())
	}
	return lat, nil
}

// <=============End Implementations=============>//
func (a *BaseAddress) combineMultipleFormats(formats []func(*BaseAddress) string, item string, space bool) string {
	for i := range formats {
		selectedMethod := formats[i]
		currentName := a.invokeSelectedMethodAddress(selectedMethod)
		if len(item) < 0 {
			item = currentName
		} else {
			if space {
				item = item + " " + currentName
			} else {
				item = item + currentName
			}

		}
	}
	return strings.Trim(item, " ")
}

func (a *BaseAddress) invokeSelectedMethodAddress(fn func(*BaseAddress) string) string {
	return fn(a)
}
