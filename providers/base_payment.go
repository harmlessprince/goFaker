package providers

import (
	"github.com/harmlessprince/goFaker/calculator"
	"github.com/harmlessprince/goFaker/parsers"
	"log"
	"time"
)

type CreditCardNumberParams struct {
	CreditCardType string
	Formatted      bool
	Separator      string
}
type CreditCardExpirationParam struct {
	Valid bool
}
type CreditCardExpirationStringParam struct {
	Valid                bool
	ExpirationDateFormat string
}

type PaymentInterface interface {
	CreditCardType() (string, error)
	CreditCardNumber(params ...CreditCardNumberParams) (string, error)
	CreditCardExpirationDate(param ...CreditCardExpirationParam) (time.Time, error)
	CreditCardExpirationDateString(param ...CreditCardExpirationStringParam) (string, error)
	CreditCardDetails(param ...CreditCardExpirationParam) (map[string]string, error)
	//iban() string
	//swiftBicNumber() string
}

func defaultCreditCardNumberParams() CreditCardNumberParams {
	return CreditCardNumberParams{
		CreditCardType: "",
		Formatted:      false,
		Separator:      "-",
	}
}
func newCreditCardNumberParams(param CreditCardNumberParams) CreditCardNumberParams {
	defaultParameter := defaultCreditCardNumberParams()
	if param.CreditCardType == "" {
		param.CreditCardType = defaultParameter.CreditCardType
	}
	if param.Separator == "" {
		param.Separator = defaultParameter.Separator
	}
	return param
}
func defaultCreditCardExpirationParam() CreditCardExpirationParam {
	return CreditCardExpirationParam{Valid: true}
}

func defaultCreditCardExpirationStringParam() CreditCardExpirationStringParam {

	return CreditCardExpirationStringParam{
		Valid:                true,
		ExpirationDateFormat: "01/06",
	}
}

func newCreditCardExpirationStringParam(param CreditCardExpirationStringParam) CreditCardExpirationStringParam {
	defaultParameter := defaultCreditCardExpirationStringParam()
	if param.ExpirationDateFormat == "" {
		param.ExpirationDateFormat = defaultParameter.ExpirationDateFormat
	}
	return param
}

type BasePayment struct {
	BaseProvider
	cardVendors []string
	cardParams  map[string][]string
}

func (b *BasePayment) GetCardVendors() []string {
	b.cardVendors = []string{
		"Visa", "Visa", "Visa", "Visa", "Visa",
		"MasterCard", "MasterCard", "MasterCard", "MasterCard", "MasterCard",
		"American Express", "Discover Card", "Visa Retired", "JCB",
	}
	return b.cardVendors
}

func (b *BasePayment) GetCardParams() map[string][]string {
	b.cardParams = map[string][]string{
		"Visa": {
			"4539###########",
			"4556###########",
			"4916###########",
			"4532###########",
			"4929###########",
			"40240071#######",
			"4485###########",
			"4716###########",
			"4##############",
		},
		"Visa Retired": {
			"4539########",
			"4556########",
			"4916########",
			"4532########",
			"4929########",
			"40240071####",
			"4485########",
			"4716########",
			"4###########",
		},
		"MasterCard": {
			"2221###########",
			"23#############",
			"24#############",
			"25#############",
			"26#############",
			"2720###########",
			"51#############",
			"52#############",
			"53#############",
			"54#############",
			"55#############",
		},
		"American Express": {
			"34############",
			"37############",
		},
		"Discover Card": {
			"6011###########",
		},
		"JCB": {
			"3528###########",
			"3589###########",
		},
	}
	return b.cardParams
}

func (b *BasePayment) CreditCardType() (string, error) {
	cardVendor, err := b.BaseProvider.RandomElementFromStringSlice(b.GetCardVendors())
	if err != nil {
		return "", err
	}
	return cardVendor, nil
}

func (b *BasePayment) CreditCardNumber(params ...CreditCardNumberParams) (string, error) {
	param := defaultCreditCardNumberParams()
	if len(params) > 0 {
		param = newCreditCardNumberParams(params[0])
	}
	if param.CreditCardType == "" {
		creditCardType, err := b.CreditCardType()
		if err != nil {
			return "", err
		}
		param.CreditCardType = creditCardType
	}
	mask, err := b.BaseProvider.RandomElementFromStringSlice(b.GetCardParams()[param.CreditCardType])
	if err != nil {
		log.Fatal(err)
	}
	creditCardNumber, err := b.BaseProvider.Numerify(mask)
	if err != nil {
		log.Fatal(err)
	}
	luhn := &calculator.Luhn{}
	creditCardNumber = creditCardNumber + luhn.ComputeCheckDigit(creditCardNumber)

	if param.Formatted == true {
		firstDigits := helperInstance.Substr(creditCardNumber, 0, 4)
		secondDigits := helperInstance.Substr(creditCardNumber, 4, 4)
		thirdDigits := helperInstance.Substr(creditCardNumber, 8, 4)
		fourthDigits := helperInstance.Substr(creditCardNumber, 12, len(creditCardNumber)-12)
		creditCardNumber = firstDigits + param.Separator + secondDigits + param.Separator + thirdDigits + param.Separator + fourthDigits
	}

	return creditCardNumber, nil
}

func (b *BasePayment) CreditCardExpirationDate(params ...CreditCardExpirationParam) (time.Time, error) {
	baseDateTime := &BaseDateTime{}
	param := defaultCreditCardExpirationParam()
	if len(params) > 0 {
		param = params[0]
	}
	if param.Valid == true {
		endTime, _ := parsers.StringToDuration{}.ParseDuration("36months")
		return baseDateTime.DateTimeBetween(time.Now(), time.Now().Add(endTime), "")
	}
	startTime, startTimeError := parsers.StringToDuration{}.ParseDuration("-36months")
	if startTimeError != nil {
		log.Fatal(startTimeError)
	}
	endTime, endTimeError := parsers.StringToDuration{}.ParseDuration("36months")
	if endTimeError != nil {
		log.Fatal(endTimeError)
	}

	return baseDateTime.DateTimeBetween(time.Now().Add(startTime), time.Now().Add(endTime), "")
}

func (b *BasePayment) CreditCardExpirationDateString(params ...CreditCardExpirationStringParam) (string, error) {
	param := defaultCreditCardExpirationStringParam()
	if len(params) > 0 {
		param = newCreditCardExpirationStringParam(params[0])
	}
	expirationDate, err := b.CreditCardExpirationDate(CreditCardExpirationParam{
		Valid: param.Valid,
	})

	if err != nil {
		return "", err
	}

	return expirationDate.Format(param.ExpirationDateFormat), nil
}

func (b *BasePayment) CreditCardDetails(params ...CreditCardExpirationParam) (map[string]string, error) {
	param := defaultCreditCardExpirationParam()
	if len(params) > 0 {
		param = params[0]
	}
	creditCardType, err := b.CreditCardType()
	if err != nil {
		return nil, err
	}
	creditCardNumber, err := b.CreditCardNumber(CreditCardNumberParams{CreditCardType: creditCardType})
	if err != nil {
		return nil, err
	}
	name, err := newBasePerson.Name()
	if err != nil {
		return nil, err
	}
	expirationDate, err := b.CreditCardExpirationDateString(CreditCardExpirationStringParam{Valid: param.Valid})
	if err != nil {
		return nil, err
	}
	return map[string]string{
		"type":           creditCardType,
		"number":         creditCardNumber,
		"name":           name,
		"expirationDate": expirationDate,
	}, nil
}
