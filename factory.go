package factory

import (
	"github.com/harmlessprince/goFaker/extensions"
	"github.com/harmlessprince/goFaker/providers"
	"github.com/harmlessprince/goFaker/providers/en_NG"
)

func MakeFactory(locality string) extensions.DataGeneratorExtension {
	switch locality {
	case "en_NG":
		return &en_NG.EnNGGenerator{}
	default:
		return &providers.DefaultGenerator{}
	}
}
