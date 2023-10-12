package extensions

type InternetExtension interface {
	Email() string
	SafeEmail() string
	FreeEmail() string
	CompanyEmail() string
	FreeEmailDomain() string
	SafeEmailDomain() string
	UserName() string
	Password(params ...int) string
	DomainName() string
	DomainWord() string
	Url() string
	Tld() string
	Ipv4() string
	Ipv6() string
	LocalIpv4() string
	MacAddress() string
	Transliterate(inputString string) string
	ToAscii(inputString string) string
	Slug(numberOfWords int, variableNumberOfWords bool) string
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
