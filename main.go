package main

import (
	"fmt"
	"github.com/harmlessprince/goFaker/providers"
	"github.com/harmlessprince/goFaker/providers/en_NG"
)

func main() {
	locality := "en_NG"
	generator := GeneratorFactory(locality)
	fmt.Println(generator.Date("", "2020-08-10 15:04:05"))
}

func GeneratorFactory(locality string) providers.DataGenerator {
	switch locality {
	case "en_NG":
		return &en_NG.EnNGGenerator{}
	default:
		return &providers.DefaultGenerator{}
	}
}
