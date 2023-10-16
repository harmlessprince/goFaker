package goFaker

import (
	"github.com/harmlessprince/goFaker/extensions"
	"github.com/harmlessprince/goFaker/providers"
	"github.com/harmlessprince/goFaker/providers/en_NG"
)

func MakeFactory(locality ...string) extensions.DataGeneratorExtension {
	var localityCode string
	if len(locality) > 0 {
		localityCode = locality[0]
	}
	switch localityCode {
	case "en_NG":
		return &en_NG.EnNGGenerator{}
	default:
		return &providers.DefaultGenerator{}
	}
}
