package main

import (
	"fmt"
	"github.com/harmlessprince/goFaker/providers"
)

func main() {
	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println("Recovered. Error:\n", r)
	//		fmt.Println("go on")
	//	}
	//}()
	provider := &providers.BaseBarcode{}

	fmt.Println(provider.Isbn10())

	//factory := goFaker.MakeFactory(goFaker.MakeFactoryParam{Locality: "en_NG"})
	////factory := providers.DefaultGenerator{}
	//fmt.Println(factory.SentencesAsText(3, false))
	//fmt.Println(factory.SentencesAsList(3, false))
	//fmt.Println(factory.Paragraph(2, true))
	//fmt.Println(factory.ImageUrl())
}
