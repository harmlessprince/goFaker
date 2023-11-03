package goFaker

import (
	"github.com/harmlessprince/goFaker/providers"
	"github.com/harmlessprince/goFaker/providers/en_NG"
)

type MakeFactoryParam struct {
	Locality string
}

func defaultMakeFactory() MakeFactoryParam {
	return MakeFactoryParam{Locality: ""}
}
func MakeFactory(locality ...MakeFactoryParam) providers.DataGeneratorExtension {
	localityCode := defaultMakeFactory()
	if len(locality) > 0 {
		localityCode = locality[0]
	}
	switch localityCode.Locality {
	case "en_NG":
		return &en_NG.EnNGGenerator{}
	default:
		return &providers.DefaultGenerator{}
	}
}
