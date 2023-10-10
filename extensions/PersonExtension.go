package extensions

type PersonInterface interface {
	Name(gender ...string) string

	FirstName(gender ...string) string
	FirstNameMale() string
	FirstNameFemale() string

	LastName() string

	Title(gender ...string) string
	TitleMale() string
	TitleFemale() string
}
