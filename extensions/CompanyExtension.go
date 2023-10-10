package extensions

type CompanyExtension interface {
	Company() string
	CompanySuffix() string
	JobTitle() string
}
