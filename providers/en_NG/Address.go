package en_NG

import (
	"github.com/harmlessprince/goFaker/constants"
	"github.com/harmlessprince/goFaker/providers"
)

type Address struct {
	providers.Address
}

func (a *Address) SetCountry() {
	address := providers.Address{}
	address.SetCountry(constants.CountryNamesEnglish)
}
