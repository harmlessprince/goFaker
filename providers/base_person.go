package providers

type PersonInterface interface {
	Name(gender ...string) (string, error)
	FirstName(gender ...string) (string, error)
	FirstNameMale() (string, error)
	FirstNameFemale() (string, error)
	LastName() (string, error)
	Title(gender ...string) (string, error)
	TitleMale() (string, error)
	TitleFemale() (string, error)
}

const GenderFemale string = "female"
const GenderMale string = "male"

type BasePerson struct {
	BaseProvider
	titleFormats      []string
	firstNameFormats  []string
	maleNameFormats   []string
	femaleNameFormats []string
	maleFirstNames    []string
	femaleFirstNames  []string
	lastNames         []string
	maleTitles        []string
	femaleTitles      []string
	genderMale        string
	genderFemale      string
}

func NewBasePerson() *BasePerson {
	p := &BasePerson{
		titleFormats: []string{
			"{{TitleMale}}",
			"{{TitleFemale}}",
		},
		firstNameFormats: []string{
			"{{FirstNameMale}}",
			"{{FirstNameFemale}}",
		},
		maleNameFormats:   []string{"{{FirstNameMale}} {{LastName}}"},
		femaleNameFormats: []string{"{{FirstNameFemale}} {{LastName}}"},
		maleFirstNames:    []string{"John", "Kane"},
		femaleFirstNames:  []string{"Jane", "Clara"},
		lastNames:         []string{"Harry"},
		maleTitles:        []string{"Mr.", "Dr.", "Prof."},
		femaleTitles:      []string{"Mrs.", "Ms.", "Miss", "Dr.", "Prof."},
		genderMale:        GenderMale,
		genderFemale:      GenderFemale,
	}
	return p
}

// GetTitleFormats is a method that provides access to the titleFormats.
func (p *BasePerson) GetTitleFormats() []string {
	return p.titleFormats
}

// SetFirstNameFormats is a method that provides access to set the firstNameFormats.
func (p *BasePerson) SetFirstNameFormats(param []string) {
	p.firstNameFormats = param
}

// GetFirstNameFormats is a method that provides access to the firstNameFormats.
func (p *BasePerson) GetFirstNameFormats() []string {
	return p.firstNameFormats
}

// GetMaleNameFormats is a method that provides access to the maleNameFormats.
func (p *BasePerson) GetMaleNameFormats() []string {
	return p.maleNameFormats
}
func (p *BasePerson) SetMaleNameFormats(param []string) {
	p.maleNameFormats = param
}

// GetFemaleNameFormats is a method that provides access to the femaleNameFormats.
func (p *BasePerson) GetFemaleNameFormats() []string {
	return p.femaleNameFormats
}
func (p *BasePerson) SetFemaleNameFormats(param []string) {
	p.femaleNameFormats = param
}

func (p *BasePerson) GetMaleFirstNames() []string {
	return p.maleFirstNames
}

func (p *BasePerson) SetMaleFirstNames(param []string) {
	p.maleFirstNames = param
}

func (p *BasePerson) GetFemaleFirstNames() []string {
	return p.femaleFirstNames
}
func (p *BasePerson) SetFemaleFirstNames(param []string) {
	p.femaleFirstNames = param
}

func (p *BasePerson) GetLastNames() []string {
	return p.lastNames
}

func (p *BasePerson) SetLastNames(param []string) {
	p.lastNames = param
}

func (p *BasePerson) GetMaleTitles() []string {
	return p.maleFirstNames
}
func (p *BasePerson) GetFemaleTitles() []string {
	return p.femaleTitles
}

// GetMaleGender is a method that provides access to the MaleGender value.
func (p *BasePerson) GetMaleGender() string {
	return p.genderMale
}

// GetFemaleGender is a method that provides access to the FemaleGender value.
func (p *BasePerson) GetFemaleGender() string {
	return p.genderFemale
}

// Name generates a random full name.
//
// Parameters:
// - gender (optional): Specify the gender of the generated name (e.g., "male" or "female").
//
// Returns:
// - A string representing a randomly generated full name.
func (p *BasePerson) Name(gender ...string) (string, error) {
	var suppliedGender string
	format := helperInstance.RandomElementString(append(p.maleNameFormats, p.femaleNameFormats...))
	if len(gender) > 0 {
		suppliedGender = gender[0]
		if suppliedGender == GenderMale {
			format = helperInstance.RandomElementString(p.maleNameFormats)
		}
		if suppliedGender == GenderFemale {
			format = helperInstance.RandomElementString(p.femaleNameFormats)
		}
	}
	fullName, err := p.Parse(format, p)
	if err != nil {
		return "", err
	}
	return fullName, nil
}

// FirstName generates a random first name.
//
// Parameters:
// - gender (optional): Specify the gender of the generated first name (e.g., "male" or "female").
//
// Returns:
// - A string representing a randomly generated first name.
func (p *BasePerson) FirstName(gender ...string) (string, error) {
	var suppliedGender string
	if len(gender) > 0 {
		suppliedGender = gender[0]
		if suppliedGender == GenderMale {
			return p.FirstNameMale()
		}
		if suppliedGender == GenderFemale {
			return p.FirstNameFemale()
		}
	}
	firstName, err := p.BaseProvider.Parse(helperInstance.RandomElementString(p.firstNameFormats), p)

	if err != nil {
		return "", err
	}
	return firstName, nil
}

// FirstNameMale generates a random male first name.
//
// Returns:
// - A string representing a randomly generated male first name.
func (p *BasePerson) FirstNameMale() (string, error) {
	firstName, err := p.RandomElementFromStringSlice(p.maleFirstNames)
	if err != nil {
		return "", err
	}
	return firstName, nil
}

// FirstNameFemale generates a random female first name.
//
// Returns:
// - A string representing a randomly generated female first name.
func (p *BasePerson) FirstNameFemale() (string, error) {
	firstName := helperInstance.RandomElementString(p.femaleFirstNames)
	return firstName, nil
}

// LastName generates a random last name.
//
// Returns:
// - A string representing a randomly generated last name.
func (p *BasePerson) LastName() (string, error) {
	lastName := helperInstance.RandomElementString(p.lastNames)
	return lastName, nil
}

// TitleMale generates a random male job title.
//
// Returns:
// - A string representing a randomly generated male job title.
func (p *BasePerson) TitleMale() (string, error) {
	title := helperInstance.RandomElementString(p.maleTitles)
	return title, nil
}

// TitleFemale generates a random female job title.
//
// Returns:
// - A string representing a randomly generated female job title.
func (p *BasePerson) TitleFemale() (string, error) {
	title := helperInstance.RandomElementString(p.femaleTitles)
	return title, nil
}

// Title generates a random job title.
//
// Parameters:
// - gender (optional): Specify the gender of the generated title (e.g., "male" or "female").
//
// Returns:
// - A string representing a randomly generated job title.
func (p *BasePerson) Title(gender ...string) (string, error) {
	var suppliedGender string
	if len(gender) > 0 {
		suppliedGender = gender[0]
		if suppliedGender == GenderMale {
			return p.TitleMale()
		}
		if suppliedGender == GenderFemale {
			return p.TitleFemale()
		}
	}
	return p.BaseProvider.Parse(helperInstance.RandomElementString(p.titleFormats), p)
}

// invokeSelectedMethod
func (p *BasePerson) invokeSelectedMethod(fn func(*BasePerson) string) string {
	return fn(p)
}
