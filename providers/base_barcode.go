package providers

import (
	"github.com/harmlessprince/goFaker/calculator"
	"github.com/samber/lo"
	"strconv"
	"strings"
)

type BarcodeInterface interface {
	Ean13() (string, error)
	Ean8() (string, error)
	Isbn10() (string, error)
	Isbn13() (string, error)
}
type BaseBarcode struct {
	BaseProvider
}

func (bc *BaseBarcode) ean(lengthParam ...int) (string, error) {
	length := 13
	if len(lengthParam) > 0 {
		length = lengthParam[0]
	}
	hashes := lo.Times(length-1, func(i int) string {
		return "#"
	})

	code, err := bc.Numerify(strings.Join(hashes, ""))
	if err != nil {
		return "", err
	}
	c := &calculator.Ean{}
	ean, err := c.CheckSum(code)
	return code + strconv.FormatInt(int64(ean), 10), nil
}

func (bc *BaseBarcode) isbnCheckSum(input string) (string, error) {
	isbnCalc := &calculator.Isbn{}
	ean, err := isbnCalc.CheckSum(input)
	if err != nil {
		return "", err
	}
	return ean, nil
}

func (bc *BaseBarcode) Ean13() (string, error) {
	return bc.ean(13)
}

func (bc *BaseBarcode) Ean8() (string, error) {
	return bc.ean(8)
}

func (bc *BaseBarcode) Isbn10() (string, error) {
	hashes := lo.Times(9, func(i int) string {
		return "#"
	})
	code, err := bc.Numerify(strings.Join(hashes, ""))
	if err != nil {
		return "", err
	}
	isbnCalc := &calculator.Isbn{}
	isbnCheckSum, err := isbnCalc.CheckSum(code)
	if err != nil {
		return "", err
	}
	return code + isbnCheckSum, nil
}

func (bc *BaseBarcode) Isbn13() (string, error) {
	numberBtw8To, err := bc.NumberBetween(8, 9)
	if err != nil {
		return "", err
	}
	hashes := lo.Times(9, func(i int) string {
		return "#"
	})
	numerifyCode, err := bc.Numerify(strings.Join(hashes, ""))
	if err != nil {
		return "", err
	}
	code := "97" + strconv.FormatInt(int64(numberBtw8To), 10) + numerifyCode
	eanCalc := &calculator.Ean{}
	ean, err := eanCalc.CheckSum(code)
	if err != nil {
		return "", err
	}
	return code + strconv.FormatInt(int64(ean), 10), nil
}
