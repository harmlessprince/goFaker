package main

import (
	"fmt"
	"github.com/harmlessprince/goFaker"
)

func main() {
	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println("Recovered. Error:\n", r)
	//		fmt.Println("go on")
	//	}
	//}()
	//baseText := providers.NewBaseText()
	//
	//fmt.Println(baseText.RealText())

	factory := goFaker.MakeFactory(goFaker.MakeFactoryParam{Locality: "en_NG"})
	//factory := providers.DefaultGenerator{}
	fmt.Println(factory.SentencesAsText(3, false))
	fmt.Println(factory.SentencesAsList(3, false))
	fmt.Println(factory.Paragraph(2, true))
}
