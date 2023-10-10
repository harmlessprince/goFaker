package main

import (
	"fmt"
	"github.com/harmlessprince/goFaker/providers"
	"github.com/harmlessprince/goFaker/providers/en_NG"
)

func main() {
	address := providers.Address{}
	//
	fmt.Println(address.Address())
	//locality := "en_NG"
	//generator := GeneratorFactory(locality)
	//
	//fmt.Println(generator.Address())
}

//type Generator interface {
//	Country() string
//}

func GeneratorFactory(locality string) providers.DataGenerator {
	switch locality {
	case "en_NG":
		return &en_NG.EnNGGenerator{}
	default:
		panic("Not Implemented")
		return &providers.DefaultGenerator{}
	}
}
