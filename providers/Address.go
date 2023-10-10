package providers

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var myGlobalStruct *Address

type Address struct {
	BaseProvider
	cityFormats          [][]func(*Address) string
	streetNameFormats    [][]func(*Address) string
	streetAddressFormats [][]func(*Address) string
	addressFormats       [][]func(*Address) string
	postcode             []string
	country              []string
	buildingNumber       []string
	citySuffix           []string
	streetSuffix         []string
	cityPrefix           []string
	streetPrefix         []string
}

func init() {
	address := &Address{}
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
}

//<=============Start Setters=============>//

func (a *Address) SetStreetSuffix(param ...[]string) {
	if len(param) > 0 {
		a.streetSuffix = param[0]
	} else {
		a.streetSuffix = []string{
			"Street",
		}
	}

}
func (a *Address) SetPostCodes(param ...[]string) {
	if len(param) > 0 {
		a.postcode = param[0]
	} else {
		a.postcode = []string{
			"#####",
		}
	}
}
func (a *Address) SetBuildingNumber(param ...[]string) {

	if len(param) > 0 {
		a.buildingNumber = param[0]
	} else {
		a.buildingNumber = []string{
			"%#",
		}
	}
}
func (a *Address) SetCitySuffix(param ...[]string) {
	if len(param) > 0 {
		a.citySuffix = param[0]
	} else {
		a.citySuffix = []string{
			"Ville",
		}
	}
}
func (a *Address) SetCityPrefix(param ...[]string) {
	if len(param) > 0 {
		a.cityPrefix = param[0]
	} else {
		a.cityPrefix = []string{
			"Clanton", "Jasper", "GreenVille",
		}
	}
}
func (a *Address) SetStreetPrefix(param ...[]string) {
	if len(param) > 0 {
		a.streetPrefix = param[0]
	} else {
		a.streetPrefix = []string{
			"Decator", "Demopolis", "Florence",
		}
	}
}
func (a *Address) SetCountry(param ...[]string) {
	if len(param) > 0 {
		a.country = param[0]
	} else {
		a.country = []string{
			"Nigeria",
		}
	}

}

func (a *Address) SetCityFormats(param ...[][]func(*Address) string) {
	if len(param) > 0 {
		a.cityFormats = param[0]
	} else {
		a.cityFormats = [][]func(*Address) string{
			{(*Address).CityName, (*Address).CitySuffix},
		}
	}

}
func (a *Address) SetStreetNameFormats(param ...[][]func(*Address) string) {
	if len(param) > 0 {
		a.streetNameFormats = param[0]
	} else {
		a.streetNameFormats = [][]func(*Address) string{
			{(*Address).CitySuffix, (*Address).StreetPrefix},
		}
	}

}
func (a *Address) SetStreetAddressFormats(param ...[][]func(*Address) string) {
	if len(param) > 0 {
		a.streetAddressFormats = param[0]
	} else {
		a.streetAddressFormats = [][]func(*Address) string{
			{(*Address).BuildingNumber, (*Address).StreetName},
		}
	}

}

func (a *Address) SetAddressFormats(param ...[][]func(*Address) string) {
	if len(param) > 0 {
		a.addressFormats = param[0]
	} else {
		a.addressFormats = [][]func(*Address) string{
			{(*Address).StreetAddress, (*Address).PostCode, (*Address).City},
		}
	}

}

//<=============End Setters=============>//

//<=============Start Getters=============>//

func (a *Address) GetStreetSuffix() []string {
	return a.streetSuffix
}

func (a *Address) GetPostCodes() []string {
	return a.postcode
}
func (a *Address) GetBuildingNumber() []string {
	return a.buildingNumber
}
func (a *Address) GetCitySuffix() []string {
	return a.citySuffix
}
func (a *Address) GetCityPrefix() []string {
	return a.cityPrefix
}
func (a *Address) GetStreetPrefix() []string {
	return a.streetPrefix
}

func (a *Address) GetCountry() []string {
	return a.country
}

// GetCityFormats is a method that provides access to the cityFormats.
func (a *Address) GetCityFormats() [][]func(*Address) string {
	return a.cityFormats
}

// GetStreetNameFormats is a method that provides access to the streetNameFormats.
func (a *Address) GetStreetNameFormats() [][]func(*Address) string {

	return a.streetNameFormats
}

// GetStreetAddressFormats is a method that provides access to the streetAddressFormats.
func (a *Address) GetStreetAddressFormats() [][]func(*Address) string {

	return a.streetAddressFormats
}

// GetAddressFormats is a method that provides access to the addressFormats.
func (a *Address) GetAddressFormats() [][]func(*Address) string {
	return a.addressFormats
}

//<=============End Getters=============>//

//<=============Start Implementations=============>//

func (a *Address) CityName() string {
	cityPrefix, err := a.BaseProvider.RandomElementFromStringSlice(a.GetCityPrefix())
	if err != nil {
		panic(err.Error())
		return ""
	}
	return cityPrefix
}
func (a *Address) CityPrefix() string {
	citySuffix, err := a.BaseProvider.RandomElementFromStringSlice(a.GetCitySuffix())
	if err != nil {
		panic(err.Error())
		return ""
	}
	return citySuffix
}
func (a *Address) StreetPrefix() string {
	streetPrefix, err := a.BaseProvider.RandomElementFromStringSlice(a.GetStreetPrefix())
	if err != nil {
		panic(err.Error())
		return ""
	}
	return streetPrefix
}
func (a *Address) CitySuffix() string {
	citySuffix, err := a.BaseProvider.RandomElementFromStringSlice(a.GetCitySuffix())
	if err != nil {
		panic(err.Error())
	}
	return citySuffix
}

func (a *Address) StreetSuffix() string {
	citySuffix, err := a.BaseProvider.RandomElementFromStringSlice(a.GetStreetSuffix())
	if err != nil {
		panic(err.Error())
	}
	return citySuffix
}

func (a *Address) Country() string {
	country, err := a.BaseProvider.RandomElementFromStringSlice(a.GetCountry())
	if err != nil {
		panic(err.Error())
		return ""
	}
	return country
}
func (a *Address) Address() string {
	fmt.Println(a.country)
	formats := a.GetAddressFormats()
	rand.New(rand.NewSource(time.Now().UnixNano()))
	var address string
	randomIndex := rand.Intn(len(formats))
	address = combineMultipleFormats(formats[randomIndex], address, true)
	return address
}

func (a *Address) City() string {
	formats := a.GetCityFormats()
	rand.New(rand.NewSource(time.Now().UnixNano()))
	var city string
	randomIndex := rand.Intn(len(formats))
	city = combineMultipleFormats(formats[randomIndex], city, false)
	return city
}

func (a *Address) PostCode() string {
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

func (a *Address) StreetName() string {
	formats := a.GetStreetNameFormats()
	rand.New(rand.NewSource(time.Now().UnixNano()))
	var streetName string
	randomIndex := rand.Intn(len(formats))
	streetName = combineMultipleFormats(formats[randomIndex], streetName, true)
	return streetName
}

func (a *Address) StreetAddress() string {
	formats := a.GetStreetAddressFormats()
	rand.New(rand.NewSource(time.Now().UnixNano()))
	var streetAddress string
	randomIndex := rand.Intn(len(formats))
	streetAddress = combineMultipleFormats(formats[randomIndex], streetAddress, true)
	return streetAddress
}

func (a *Address) BuildingNumber() string {
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
func (a *Address) LocalCoordinates() map[string]float64 {
	latitude := a.Latitude()
	longitude := a.Longitude()
	coordinate := map[string]float64{
		"latitude":  latitude,
		"longitude": longitude,
	}
	return coordinate
}

func (a *Address) Longitude() float64 {
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

func (a *Address) Latitude() float64 {
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
func combineMultipleFormats(formats []func(*Address) string, item string, space bool) string {
	for i := range formats {
		selectedMethod := formats[i]
		currentName := invokeSelectedMethodAddress(selectedMethod, &Address{})
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

func invokeSelectedMethodAddress(fn func(*Address) string, addressStruct *Address) string {
	return fn(addressStruct)
}
