package en_NG

import (
	"github.com/harmlessprince/goFaker/constants"
	"github.com/harmlessprince/goFaker/providers"
)

type EnNGAddress struct {
	providers.BaseAddress
}

func (a *EnNGAddress) SetCountry() {
	a.BaseAddress.SetCountry(constants.CountryNamesEnglish)
}
func (a *EnNGAddress) SetPostCodes() {
	a.BaseAddress.SetPostCodes([]string{"## ###", "#####"})
}

func NewAddress() *EnNGAddress {
	address := &EnNGAddress{}
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
