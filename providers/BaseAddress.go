package providers

import (
	"math/rand"
	"strings"
	"time"
)

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

func NewAddress() *BaseAddress {
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
			"{{CityName CitySuffix}}",
		}
	}

}
func (a *BaseAddress) SetStreetNameFormats(param ...[]string) {
	if len(param) > 0 {
		a.streetNameFormats = param[0]
	} else {
		a.streetNameFormats = []string{
			"{{StreetName}} {{StreetSuffix}}",
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

func (a *BaseAddress) CityName() string {
	cityPrefix, err := a.BaseProvider.RandomElementFromStringSlice(a.GetCityPrefix())
	if err != nil {
		panic(err.Error())
		return ""
	}
	return cityPrefix
}
func (a *BaseAddress) CityPrefix() string {
	citySuffix, err := a.BaseProvider.RandomElementFromStringSlice(a.GetCitySuffix())
	if err != nil {
		panic(err.Error())
		return ""
	}
	return citySuffix
}
func (a *BaseAddress) StreetPrefix() string {
	streetPrefix, err := a.BaseProvider.RandomElementFromStringSlice(a.GetStreetPrefix())
	if err != nil {
		panic(err.Error())
		return ""
	}
	return streetPrefix
}
func (a *BaseAddress) CitySuffix() string {
	citySuffix, err := a.BaseProvider.RandomElementFromStringSlice(a.GetCitySuffix())
	if err != nil {
		panic(err.Error())
	}
	return citySuffix
}

func (a *BaseAddress) StreetSuffix() string {
	citySuffix, err := a.BaseProvider.RandomElementFromStringSlice(a.GetStreetSuffix())
	if err != nil {
		panic(err.Error())
	}
	return citySuffix
}

func (a *BaseAddress) Country() string {
	country, err := a.BaseProvider.RandomElementFromStringSlice(a.GetCountry())
	if err != nil {
		panic(err.Error())
		return ""
	}
	return country
}
func (a *BaseAddress) Address() string {
	formats := a.GetAddressFormats()
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := rand.Intn(len(formats))
	address, err := a.BaseProvider.Parse(formats[randomIndex], a)
	if err != nil {
		panic(err.Error())
		return ""
	}
	return address
}

func (a *BaseAddress) City() string {
	formats := a.GetCityFormats()
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := rand.Intn(len(formats))
	city, err := a.BaseProvider.Parse(formats[randomIndex], a)
	if err != nil {
		panic(err.Error())
		return ""
	}
	return city
}

func (a *BaseAddress) PostCode() string {
	randomCode, err := a.BaseProvider.RandomElementFromStringSlice(a.GetPostCodes())
	if err != nil {
		panic(err.Error())
		return ""
	}
	postCode, err := a.BaseProvider.Bothify(randomCode)
	if err != nil {
		panic(err.Error())
		return ""
	}
	return a.BaseProvider.ToLower(postCode)
}

func (a *BaseAddress) StreetName() string {
	formats := a.GetStreetNameFormats()
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := rand.Intn(len(formats))
	streetName, err := a.BaseProvider.Parse(formats[randomIndex], a)
	if err != nil {
		panic(err.Error())
		return ""
	}
	return streetName
}

func (a *BaseAddress) StreetAddress() string {
	formats := a.GetStreetAddressFormats()
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := rand.Intn(len(formats))
	streetAddress, err := a.BaseProvider.Parse(formats[randomIndex], a)
	if err != nil {
		panic(err.Error())
		return ""
	}
	return streetAddress
}

func (a *BaseAddress) BuildingNumber() string {
	randomBuildNumber, err := a.BaseProvider.RandomElementFromStringSlice(a.GetBuildingNumber())
	if err != nil {
		panic(err.Error())
		return ""
	}
	generatedBuildNumber, err := a.BaseProvider.Numerify(randomBuildNumber)
	if err != nil {
		panic(err.Error())
		return ""
	}

	return generatedBuildNumber
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
func (a *BaseAddress) LocalCoordinates() map[string]float64 {
	latitude := a.Latitude()
	longitude := a.Longitude()
	coordinate := map[string]float64{
		"latitude":  latitude,
		"longitude": longitude,
	}
	return coordinate
}

func (a *BaseAddress) Longitude() float64 {
	options := RandomFloatOptions{
		NumberOfMaxDecimals: 6,
		Min:                 -180,
		Max:                 180,
	}
	long, err := a.BaseProvider.RandomFloat(options)
	if err != nil {
		panic(err.Error())
		return 0.0
	}
	return long
}

func (a *BaseAddress) Latitude() float64 {
	options := RandomFloatOptions{
		NumberOfMaxDecimals: 6,
		Min:                 -90,
		Max:                 90,
	}
	lat, err := a.BaseProvider.RandomFloat(options)
	if err != nil {
		panic(err.Error())
		return 0.0
	}
	return lat
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
