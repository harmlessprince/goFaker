package providers

import (
	"fmt"
	"github.com/rainycape/unidecode"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

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

func (b *BaseInternet) Email() string {
	fomart, err := b.RandomElementFromStringSlice(b.GetEmailFormats())
	if err != nil {
		log.Fatal(err.Error())
		return ""
	}
	parsed, parsedErr := b.Parse(fomart, b)
	if parsedErr != nil {
		log.Fatal(parsedErr.Error())
		return ""
	}
	return parsed
}

func (b *BaseInternet) SafeEmail() string {
	//regular expression to match spaces.
	re := regexp.MustCompile(`\s+`)
	input := b.UserName() + b.SafeEmailDomain()
	email := re.ReplaceAllString(input, "")
	return email
}

func (b *BaseInternet) FreeEmail() string {
	//regular expression to match spaces.
	re := regexp.MustCompile(`\s+`)
	input := b.UserName() + b.FreeEmailDomain()
	email := re.ReplaceAllString(input, "")
	return email
}

func (b *BaseInternet) CompanyEmail() string {
	//regular expression to match spaces.
	re := regexp.MustCompile(`\s+`)
	input := b.UserName() + "@" + b.DomainName()
	email := re.ReplaceAllString(input, "")
	return email
}

func (b *BaseInternet) FreeEmailDomain() string {
	freeEmailDomain, err := b.RandomElementFromStringSlice(b.GetFreeEmailDomain())
	if err != nil {
		log.Fatal(err.Error())
		return ""
	}
	return freeEmailDomain
}

func (b *BaseInternet) SafeEmailDomain() string {
	domains := []string{
		"example.com",
		"example.org",
		"example.net",
	}
	safeEmailDomain, err := b.RandomElementFromStringSlice(domains)
	if err != nil {
		log.Fatal(err.Error())
		return ""
	}
	return safeEmailDomain
}

func (b *BaseInternet) UserName() string {
	format, randomErr := b.RandomElementFromStringSlice(b.GetUserNameFormats())
	if randomErr != nil {
		log.Fatal(randomErr.Error())
		return ""
	}
	parsed, parseErr := b.Parse(format, b)
	if parseErr != nil {
		log.Fatal(parseErr.Error())
		return ""
	}
	userName, bothifyErr := b.Bothify(parsed)
	if bothifyErr != nil {
		log.Fatal(bothifyErr.Error())
		return ""
	}
	userName = strings.ToLower(b.Transliterate(userName))
	userName = strings.Replace(userName, "..", ".", -1)
	return userName
}

func (b *BaseInternet) Password(params ...int) string {
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
		log.Fatal(err.Error())
		return ""
	}
	pattern := strings.Repeat("*", numberBetween)
	password, asciifyErr := b.Asciify(pattern)
	if asciifyErr != nil {
		log.Fatal(asciifyErr.Error())
		return ""
	}
	return password
}

func (b *BaseInternet) DomainName() string {
	return b.DomainWord() + b.Tld()
}

func (b *BaseInternet) DomainWord() string {
	lastName, err := b.BaseProvider.Parse("{{LastName}}", newPerson)
	if err != nil {
		log.Fatal(err.Error())
		return ""
	}
	lastName = b.BaseProvider.ToLower(b.Transliterate(lastName))
	lastName = strings.Trim(lastName, "._")
	if lastName == "" {
		log.Fatal("DomainWord failed")
		return ""
	}
	lastName = strings.TrimSuffix(lastName, ".")
	return lastName
}

func (b *BaseInternet) Url() string {
	format, _ := b.RandomElementFromStringSlice(b.GetUrlFormats())
	url, err := b.Parse(format, b)
	if err != nil {
		log.Fatal(err.Error())
		return ""
	}
	return url
}

func (b *BaseInternet) Tld() string {
	tld, err := b.RandomElementFromStringSlice(b.GetTld())
	if err != nil {
		log.Fatal(err.Error())
		return ""
	}
	return tld
}

func (b *BaseInternet) Ipv4() string {
	bm := &BaseMiscellaneous{}
	var number int
	var err error
	booleanValue := bm.Boolean()
	number, err = b.NumberBetween(16777216, 2147483647)
	if booleanValue == true {
		number, err = b.NumberBetween(-2147483648, -2)
	}
	if err != nil {
		log.Fatal(err.Error())
	}
	ip, errLong2ip := helperInstance.Long2ip(uint32(number))
	if errLong2ip != nil {
		log.Fatal(errLong2ip.Error())
	}
	return ip
}

func (b *BaseInternet) Ipv6() string {
	var res []string
	for i := 0; i < 8; i++ {
		numberBetween, err := b.NumberBetween(0, 65535)
		if err != nil {
			log.Fatal(err.Error())
		}
		inHexDecimal := helperInstance.DecimalToHexDecimal(int64(numberBetween))
		res = append(res, inHexDecimal)
	}
	return strings.Join(res, ":")
}

func (b *BaseInternet) LocalIpv4() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	formats := b.GetLocalIPBlocks()
	randomIndex := rand.Intn(len(formats))
	selectedFormat := formats[randomIndex]
	IpToLongBlock1, err1 := helperInstance.Ip2long(selectedFormat[0])
	if err1 != nil {
		log.Fatal(err1.Error())
	}
	IpToLongBlock2, err2 := helperInstance.Ip2long(selectedFormat[1])
	if err2 != nil {
		log.Fatal(err2.Error())
	}

	numberBetweenBlocks, err3 := b.NumberBetween(int(IpToLongBlock1), int(IpToLongBlock2))
	if err3 != nil {
		log.Fatal(err3.Error())
	}
	localIpV4, err5 := helperInstance.Long2ip(uint32(numberBetweenBlocks))
	if err5 != nil {
		log.Fatal(err5.Error())
	}
	return localIpV4
}

func (b *BaseInternet) MacAddress() string {
	var mac []string
	for i := 0; i < 6; i += 1 {
		//0xff is 255 in hexadecimal
		number, err := b.NumberBetween(0, 0xff)
		if err != nil {
			log.Fatal(err.Error())
		}
		mac = append(mac, fmt.Sprintf("%2X", number))
	}
	return strings.Join(mac, ":")
}

func (b *BaseInternet) Transliterate(inputString string) string {
	matched, err := regexp.MatchString("^[A-Za-z0-9_.]+$", inputString)
	if err != nil {
		log.Fatal(err.Error())
	}
	if !matched {
		return inputString
	}
	return unidecode.Unidecode(inputString)
}

// Slug generates random slugs
// Example: 'aut-repellat-commodi-vel-itaque-nihil-id-saepe-nostrum'
func (b *BaseInternet) Slug(numberOfWords int, variableNumberOfWords bool) string {
	if numberOfWords <= 0 {
		return ""
	}
	if variableNumberOfWords {
		numberBetween, _ := b.NumberBetween(60, 140)
		numberOfWords = (numberOfWords * numberBetween) + 1
	}
	lorem := &BaseLorem{}
	words := lorem.WordsAsList(numberOfWords)
	return strings.Join(words, "-")
}

func (b *BaseInternet) ToAscii(inputString string) string {
	return strconv.QuoteToASCII(inputString)
}

func (b *BaseInternet) UserAgent() string {
	//TODO Implement user agent
	panic("implement me")
	//randNum := randIntRange(0, 4)
	//switch randNum {
	//case 0:
	//	return ChromeUserAgent()
	//case 1:
	//	return FirefoxUserAgent()
	//case 2:
	//	return SafariUserAgent()
	//case 3:
	//	return OperaUserAgent()
	//default:
	//	return ChromeUserAgent()
	//}
}
