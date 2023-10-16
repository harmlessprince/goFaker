package main

import (
	"fmt"
	"github.com/harmlessprince/goFaker/extensions"
	"github.com/harmlessprince/goFaker/providers"
	"github.com/harmlessprince/goFaker/providers/en_NG"
)

func main() {
	locality := "en_NG"
	generator := GeneratorFactory(locality)
	fmt.Println(generator.GenerateDate("", "2020-08-10 15:04:05"))
	fmt.Println(generator.GenerateCountryCode())
}

func GeneratorFactory(locality string) extensions.DataGeneratorExtension {
	switch locality {
	case "en_NG":
		return &en_NG.EnNGGenerator{}
	default:
		return &providers.DefaultGenerator{}
	}
}
