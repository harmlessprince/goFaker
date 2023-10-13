package main

import (
	"fmt"
	"github.com/harmlessprince/goFaker/providers"
	"github.com/harmlessprince/goFaker/providers/en_NG"
)

func main() {

	//data := &en_NG.EnNGGenerator{}
	//input := "{{GenerateAddress}}.{{GenerateCity}}"
	//result, err := providers.BaseProvider{}.Parse(input, data)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}

	locality := "en_NG"
	generator := GeneratorFactory(locality)
	fmt.Println(generator.GenerateRgbaCssColor())
}

func GeneratorFactory(locality string) providers.DataGenerator {
	switch locality {
	case "en_NG":
		return &en_NG.EnNGGenerator{}
	default:
		return &providers.DefaultGenerator{}
	}
}
