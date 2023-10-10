package extensions

type PhoneNumberExtension interface {
	PhoneNumber() string
	E164PhoneNumber() string
	Imei() string
}
