package extensions

type MiscellaneousExtension interface {
	Boolean(chanceOfGettingTrue ...int) bool
	Md5() string
	Sha1() string
	Sha256() string
	Locale() string
	CountryCode() string
	CountryISOAlpha3() string
	LanguageCode() string
	CurrencyCode() string
	Emoji() string
}
