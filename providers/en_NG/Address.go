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

func (a *EnNGAddress) SetCitySuffix() {
	cityNames := []string{
		"Ikeja", "Agege", "Iyana Ipaja", "Ohafia", "Umu Nneochi", "Jimeta",
		"Mubi", "Numan", "Yola", "Uyo", "Oron", "Akwa", "Onitsha", "Azare",
		"Bauchi", "Misau", "Dikwa", "Ede", "Ikire", "Ila", "Oshogbo", "Ibadan", "Oyo",
		"Gusau", "Kaura Namoda", "Mushin", "Shomolu", "Gombe", "Kumo", "Enugu", "Nsuka",
		"Sapele", "Burutu", "Ughelli", "Warri", "Calabar", "Ogoja",
	}
	a.BaseAddress.SetCitySuffix(cityNames)
}
func (a *EnNGAddress) SetStreetPrefix(param ...[]string) {
	streetNames := []string{"Broad", "Bode Thomas", "Adeola Odeku",
		"Allen Avenue", "Toyin",
		"Adetokubo Ademola", "Adeniran Ogunsanya",
		"Ozumba Mbadiwe", "Campbell", "Forsythe",
		"Balogun", "Ahmadu Bello", "Isaac John", "Ogunlana",
		"Cameron", "Reeve", "Agarawu", "Adeniji Adele", "Agodo", "Agoro",
		"Akanni", "Alakoro Marina", "Ashogbon", "Epe",
	}
	a.BaseAddress.SetStreetPrefix(streetNames)
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
	address.SetStreetPrefix()
	return address
}
