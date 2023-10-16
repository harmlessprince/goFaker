package main

import (
	"fmt"
)

func main() {
	locality := "en_NG"
	generator := MakeFactory(locality)
	fmt.Println(generator.GenerateDate("", "2020-08-10 15:04:05"))
	fmt.Println(generator.GenerateBloodType())
	fmt.Println(generator.GeneratePhoneNumber())
	fmt.Println(generator.GenerateAddress())
	fmt.Println(generator.GenerateCountryCode())
}
