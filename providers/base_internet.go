package providers

import (
	"errors"
	"fmt"
	"github.com/rainycape/unidecode"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type InternetInterface interface {
	Email() (string, error)
	SafeEmail() (string, error)
	FreeEmail() (string, error)
	CompanyEmail() (string, error)
	FreeEmailDomain() (string, error)
	SafeEmailDomain() (string, error)
	UserName() (string, error)
	Password(params ...int) (string, error)
	DomainName() (string, error)
	DomainWord() (string, error)
	Url() (string, error)
	Tld() (string, error)
	Ipv4() (string, error)
	Ipv6() (string, error)
	LocalIpv4() (string, error)
	MacAddress() (string, error)
	Transliterate(inputString string) (string, error)
	ToAscii(inputString string) (string, error)
	Slug(numberOfWords int, variableNumberOfWords bool) (string, error)
}

type InternetExtensionAccessors interface {
	SetFreeEmailDomain(params ...[]string)
	GetFreeEmailDomain() []string
	SetUserNameFormats(params ...[]string)
	GetUserNameFormats() []string
	SetEmailFormats(params ...[]string)
	GetEmailFormats() []string
	SetUrlFormats(params ...[]string)
	GetUrlFormats() []string
	SetLocalIPBlocks(params ...[][]string)
	GetLocalIPBlocks() [][]string
	SetTld(params ...[]string)
	GetTld() []string
}

type BaseInternet struct {
	BaseProvider
	freeMailDomain  []string
	tld             []string
	localIpBlocks   [][]string
	userNameFormats []string
	emailFormats    []string
	urlFormats      []string
}

func NewBaseInternet() *BaseInternet {
	baseInternet := &BaseInternet{}
	baseInternet.SetFreeEmailDomain()
	baseInternet.SetTld()
	baseInternet.SetLocalIPBlocks()
	baseInternet.SetUserNameFormats()
	baseInternet.SetEmailFormats()
	baseInternet.SetUrlFormats()
	return baseInternet
}
func (b *BaseInternet) SetFreeEmailDomain(params ...[]string) {
	b.freeMailDomain = []string{"gmail.com", "yahoo.com", "hotmail.com"}
	if len(params) > 0 {
		b.freeMailDomain = params[0]
	}
}

func (b *BaseInternet) GetFreeEmailDomain() []string {
	return b.freeMailDomain
}

func (b *BaseInternet) SetUserNameFormats(params ...[]string) {
	b.userNameFormats = []string{
		"{{LastName}}.{{FirstName}}",
		"{{FirstName}}.{{LastName}}",
		"{{FirstName}}##",
		"?{{LastName}}",
	}
	if len(params) > 0 {
		b.userNameFormats = params[0]
	}
}

func (b *BaseInternet) GetUserNameFormats() []string {
	return b.userNameFormats
}

func (b *BaseInternet) SetEmailFormats(params ...[]string) {
	b.emailFormats = []string{
		"{{UserName}}@{{DomainName}}",
		"{{UserName}}@{{FreeEmailDomain}}",
	}
	if len(params) > 0 {
		b.emailFormats = params[0]
	}
}

func (b *BaseInternet) GetEmailFormats() []string {
	return b.emailFormats
}

func (b *BaseInternet) SetUrlFormats(params ...[]string) {
	b.urlFormats = []string{
		"http://www.{{DomainName}}/",
		"http://{{DomainName}}/",
		"http://www.{{DomainName}}/{{Slug}}",
		"http://www.{{DomainName}}/{{Slug}}",
		"https://www.{{DomainName}}/{{Slug}}",
		"http://www.{{DomainName}}/{{Slug}}.html",
		"http://{{DomainName}}/{{Slug}}",
		"http://{{DomainName}}/{{Slug}}",
		"http://{{DomainName}}/{{Slug}}.html",
		"https://{{DomainName}}/{{Slug}}.html",
	}
	if len(params) > 0 {
		b.urlFormats = params[0]
	}
}

func (b *BaseInternet) GetUrlFormats() []string {
	return b.urlFormats
}

func (b *BaseInternet) SetLocalIPBlocks(params ...[][]string) {
	b.localIpBlocks = [][]string{
		{"10.0.0.0", "10.255.255.255"},
		{"172.16.0.0", "172.31.255.255"},
		{"192.168.0.0", "192.168.255.255"},
	}
	if len(params) > 0 {
		b.localIpBlocks = params[0]
	}
}

func (b *BaseInternet) GetLocalIPBlocks() [][]string {
	return b.localIpBlocks
}

func (b *BaseInternet) SetTld(params ...[]string) {
	b.tld = []string{"com", "com", "com", "com", "com", "com", "biz", "info", "net", "org"}
	if len(params) > 0 {
		b.tld = params[0]
	}
}

func (b *BaseInternet) GetTld() []string {
	return b.tld
}

func (b *BaseInternet) Email() (string, error) {
	format, err := b.RandomElementFromStringSlice(b.GetEmailFormats())
	if err != nil {
		return "", err
	}
	parsed, err := b.Parse(format, b)
	if err != nil {
		return "", err
	}
	return parsed, nil
}

func (b *BaseInternet) SafeEmail() (string, error) {
	//regular expression to match spaces.
	re := regexp.MustCompile(`\s+`)
	userName, err := b.UserName()
	if err != nil {
		return "", err
	}
	domainName, err := b.SafeEmailDomain()
	if err != nil {
		return "", err
	}
	input := userName + "@" + domainName
	email := re.ReplaceAllString(input, "")
	return email, nil
}

func (b *BaseInternet) FreeEmail() (string, error) {
	//regular expression to match spaces.
	re := regexp.MustCompile(`\s+`)
	userName, err := b.UserName()
	if err != nil {
		return "", err
	}
	domainName, err := b.FreeEmailDomain()
	if err != nil {
		return "", err
	}
	input := userName + "@" + domainName
	email := re.ReplaceAllString(input, "")
	return email, nil
}

func (b *BaseInternet) CompanyEmail() (string, error) {
	//regular expression to match spaces.
	re := regexp.MustCompile(`\s+`)
	userName, err := b.UserName()
	if err != nil {
		return "", err
	}
	domainName, err := b.DomainName()
	if err != nil {
		return "", err
	}
	input := userName + "@" + domainName
	email := re.ReplaceAllString(input, "")
	return email, nil
}

func (b *BaseInternet) FreeEmailDomain() (string, error) {
	freeEmailDomain, err := b.RandomElementFromStringSlice(b.GetFreeEmailDomain())
	if err != nil {
		return "", err
	}
	return freeEmailDomain, nil
}

func (b *BaseInternet) SafeEmailDomain() (string, error) {
	domains := []string{
		"example.com",
		"example.org",
		"example.net",
	}
	safeEmailDomain, err := b.RandomElementFromStringSlice(domains)
	if err != nil {
		return "", err
	}
	return safeEmailDomain, nil
}

func (b *BaseInternet) UserName() (string, error) {
	format, err := b.RandomElementFromStringSlice(b.GetUserNameFormats())
	if err != nil {
		return "", err
	}
	parsed, err := b.Parse(format, b)
	if err != nil {
		return "", err
	}
	userName, err := b.Bothify(parsed)
	if err != nil {
		return "", err
	}
	translation, err := b.Transliterate(userName)
	if err != nil {
		return "", err
	}
	userName = strings.ToLower(translation)
	userName = strings.Replace(userName, "..", ".", -1)
	return userName, nil
}

func (b *BaseInternet) Password(params ...int) (string, error) {
	minLength := 6
	maxLength := 20
	if len(params) > 0 {
		minLength = params[0]
	}
	if len(params) > 1 {
		maxLength = params[1]
	}
	if maxLength < minLength {
		tmp := minLength
		minLength = maxLength
		maxLength = tmp
	}
	numberBetween, err := b.NumberBetween()
	if err != nil {
		return "", err
	}
	pattern := strings.Repeat("*", numberBetween)
	password, err := b.Asciify(pattern)
	if err != nil {
		return "", err
	}
	return password, nil
}

func (b *BaseInternet) DomainName() (string, error) {
	tld, err := b.Tld()
	if err != nil {
		return "", err
	}
	domainWord, err := b.DomainWord()
	if err != nil {
		return "", err
	}
	return domainWord + tld, nil
}

func (b *BaseInternet) DomainWord() (string, error) {
	lastName, err := b.BaseProvider.Parse("{{LastName}}", NewBasePerson())
	if err != nil {
		return "", err
	}
	translation, err := b.Transliterate(lastName)
	if err != nil {
		return "", err
	}
	lastName = b.BaseProvider.ToLower(translation)
	lastName = strings.Trim(lastName, "._")
	if lastName == "" {
		return "", errors.New("DomainWord failed")
	}
	lastName = strings.TrimSuffix(lastName, ".")
	return lastName, nil
}

func (b *BaseInternet) Url() (string, error) {
	format, _ := b.RandomElementFromStringSlice(b.GetUrlFormats())
	url, err := b.Parse(format, b)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (b *BaseInternet) Tld() (string, error) {
	tld, err := b.RandomElementFromStringSlice(b.GetTld())
	if err != nil {
		return "", err
	}
	return tld, nil
}

func (b *BaseInternet) Ipv4() (string, error) {
	bm := &BaseMiscellaneous{}
	var number int
	booleanValue, err := bm.Boolean()
	if err != nil {
		return "", err
	}
	number, err = b.NumberBetween(16777216, 2147483647)
	if err != nil {
		return "", err
	}
	if booleanValue == true {
		number, err = b.NumberBetween(-2147483648, -2)
		if err != nil {
			return "", err
		}
	}
	ip, err := helperInstance.Long2ip(uint32(number))
	if err != nil {
		return "", err
	}
	return ip, nil
}

func (b *BaseInternet) Ipv6() (string, error) {
	var res []string
	for i := 0; i < 8; i++ {
		numberBetween, err := b.NumberBetween(0, 65535)
		if err != nil {
			return "", err
		}
		inHexDecimal := helperInstance.DecimalToHexDecimal(int64(numberBetween))
		res = append(res, inHexDecimal)
	}
	return strings.Join(res, ":"), nil
}

func (b *BaseInternet) LocalIpv4() (string, error) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	formats := b.GetLocalIPBlocks()
	randomIndex := rand.Intn(len(formats))
	selectedFormat := formats[randomIndex]
	IpToLongBlock1, err := helperInstance.Ip2long(selectedFormat[0])
	if err != nil {
		return "", err
	}
	IpToLongBlock2, err := helperInstance.Ip2long(selectedFormat[1])
	if err != nil {
		return "", err
	}

	numberBetweenBlocks, err := b.NumberBetween(int(IpToLongBlock1), int(IpToLongBlock2))
	if err != nil {
		return "", err
	}
	localIpV4, err := helperInstance.Long2ip(uint32(numberBetweenBlocks))
	if err != nil {
		return "", err
	}
	return localIpV4, nil
}

func (b *BaseInternet) MacAddress() (string, error) {
	var mac []string
	for i := 0; i < 6; i += 1 {
		//0xff is 255 in hexadecimal
		number, err := b.NumberBetween(0, 0xff)
		if err != nil {
			return "", err
		}
		mac = append(mac, fmt.Sprintf("%2X", number))
	}
	return strings.Join(mac, ":"), nil
}

func (b *BaseInternet) Transliterate(inputString string) (string, error) {
	matched, err := regexp.MatchString("^[A-Za-z0-9_.]+$", inputString)
	if err != nil {
		return "", err
	}
	if !matched {
		return inputString, nil
	}
	return unidecode.Unidecode(inputString), nil
}

// Slug generates random slugs
// Example: 'aut-repellat-commodi-vel-itaque-nihil-id-saepe-nostrum'
func (b *BaseInternet) Slug(numberOfWords int, variableNumberOfWords bool) (string, error) {
	if numberOfWords <= 0 {
		return "", errors.New("numberOfWords must be greater than 0")
	}
	if variableNumberOfWords {
		numberBetween, _ := b.NumberBetween(60, 140)
		numberOfWords = (numberOfWords * numberBetween) + 1
	}
	lorem := &BaseLorem{}
	words, err := lorem.WordsAsList(numberOfWords)
	if err != nil {
		return "", err
	}
	return strings.Join(words, "-"), nil
}

func (b *BaseInternet) ToAscii(inputString string) (string, error) {
	return strconv.QuoteToASCII(inputString), nil
}
