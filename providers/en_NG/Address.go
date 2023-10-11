package en_NG

import (
	"github.com/harmlessprince/goFaker/constants"
	"github.com/harmlessprince/goFaker/providers"
)

type NGAddress struct {
	providers.BaseAddress
}

func (a *NGAddress) SetCountry() {
	a.BaseAddress.SetCountry(constants.CountryNamesEnglish)
}
func (a *NGAddress) SetPostCodes() {
	a.BaseAddress.SetPostCodes([]string{"## ###", "#####"})
}

func NewAddress() *NGAddress {
	address := &NGAddress{}
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
