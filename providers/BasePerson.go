package providers

import (
	"math/rand"
	"strings"
	"time"
)

const GenderFemale string = "female"
const GenderMale string = "male"

type BasePerson struct {
	BaseProvider
	titleFormats      []func(*BasePerson) string
	firstNameFormats  []func(*BasePerson) string
	maleNameFormats   [][]func(*BasePerson) string
	femaleNameFormats [][]func(*BasePerson) string
	maleFirstNames    []string
	femaleFirstNames  []string
	lastNames         []string
	maleTitles        []string
	femaleTitles      []string
	genderMale        string
	genderFemale      string
}

func NewBasePerson() *BasePerson {
	p := &BasePerson{}
	p.SetMaleNameFormats()
	p.SetFemaleNameFormats()
	p.SetFirstNameFormats()
	p.SetMaleFirstNames()
	p.SetFemaleFirstNames()
	p.SetLastNames()
	return p
}

// GetTitleFormats is a method that provides access to the titleFormats.
func (p *BasePerson) GetTitleFormats() []func(*BasePerson) string {
	p.titleFormats = []func(*BasePerson) string{
		(*BasePerson).TitleMale,
		(*BasePerson).TitleFemale,
	}
	return p.titleFormats
}

// SetFirstNameFormats is a method that provides access to set the firstNameFormats.
func (p *BasePerson) SetFirstNameFormats(param ...[]func(*BasePerson) string) {
	if len(param) > 0 {
		p.firstNameFormats = param[0]
	} else {
		p.firstNameFormats = []func(*BasePerson) string{
			(*BasePerson).FirstNameMale, (*BasePerson).FirstNameFemale,
		}
	}
}

// GetFirstNameFormats is a method that provides access to the firstNameFormats.
func (p *BasePerson) GetFirstNameFormats() []func(*BasePerson) string {
	return p.firstNameFormats
}

// GetMaleNameFormats is a method that provides access to the maleNameFormats.
func (p *BasePerson) GetMaleNameFormats() [][]func(*BasePerson) string {
	p.maleNameFormats = [][]func(*BasePerson) string{
		{(*BasePerson).FirstNameMale, (*BasePerson).LastName},
	}
	return p.maleNameFormats
}
func (p *BasePerson) SetMaleNameFormats(param ...[][]func(*BasePerson) string) {
	if len(param) > 0 {
		p.maleNameFormats = param[0]
	} else {
		p.maleNameFormats = [][]func(*BasePerson) string{
			{(*BasePerson).FirstNameMale, (*BasePerson).LastName},
		}
	}
}

// GetFemaleNameFormats is a method that provides access to the femaleNameFormats.
func (p *BasePerson) GetFemaleNameFormats() [][]func(*BasePerson) string {
	return p.femaleNameFormats
}
func (p *BasePerson) SetFemaleNameFormats(param ...[][]func(*BasePerson) string) {
	if len(param) > 0 {
		p.femaleNameFormats = param[0]
	} else {
		p.femaleNameFormats = [][]func(*BasePerson) string{
			{(*BasePerson).FirstNameFemale, (*BasePerson).LastName},
		}
	}
}

func (p *BasePerson) GetMaleFirstNames() []string {
	return p.maleFirstNames
}

func (p *BasePerson) SetMaleFirstNames(param ...[]string) {
	if len(param) > 0 {
		p.maleFirstNames = param[0]
	} else {
		p.maleFirstNames = []string{"John", "Harry"}
	}
}

func (p *BasePerson) GetFemaleFirstNames() []string {
	return p.femaleFirstNames
}
func (p *BasePerson) SetFemaleFirstNames(param ...[]string) {
	if len(param) > 0 {
		p.femaleFirstNames = param[0]
	} else {
		p.femaleFirstNames = []string{"Jane", "Clara"}
	}
}

func (p *BasePerson) GetLastNames() []string {
	return p.lastNames
}

func (p *BasePerson) SetLastNames(param ...[]string) {
	if len(param) > 0 {
		p.lastNames = param[0]
	} else {
		p.lastNames = []string{"Doe", "Kane"}
	}
}

func (p *BasePerson) GetMaleTitles() []string {
	p.maleFirstNames = []string{"Mr.", "Dr.", "Prof."}
	return p.maleFirstNames
}
func (p *BasePerson) GetFemaleTitles() []string {
	p.femaleTitles = []string{"Mrs.", "Ms.", "Miss", "Dr.", "Prof."}
	return p.femaleTitles
}

// GetMaleGender is a method that provides access to the MaleGender value.
func (p *BasePerson) GetMaleGender() string {
	p.genderMale = "male"
	return p.genderMale
}

// GetFemaleGender is a method that provides access to the FemaleGender value.
func (p *BasePerson) GetFemaleGender() string {
	p.genderFemale = "female"
	return p.genderFemale
}
func (p *BasePerson) Name(gender ...string) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	var suppliedGender string
	var fullName string
	var formats []func(*BasePerson) string
	if len(gender) > 0 {
		suppliedGender = gender[0]
		if suppliedGender == GenderMale {
			randomIndex := rand.Intn(len(p.GetMaleNameFormats()))
			formats = p.GetMaleNameFormats()[randomIndex]
		}
		if suppliedGender == GenderFemale {
			randomIndex := rand.Intn(len(p.GetFemaleNameFormats()))
			formats = p.GetFemaleNameFormats()[randomIndex]
		}
	} else {
		allNameFormat := append(p.GetFemaleNameFormats(), p.GetMaleNameFormats()...)
		randomIndex := rand.Intn(len(allNameFormat))
		formats = allNameFormat[randomIndex]
	}
	fullName = p.parseToFullName(formats, fullName)
	return fullName
}

func (p *BasePerson) parseToFullName(formats []func(*BasePerson) string, fullName string) string {
	for i := range formats {
		selectedMethod := formats[i]
		currentName := p.invokeSelectedMethod(selectedMethod)
		if len(fullName) < 0 {
			fullName = currentName
		} else {
			fullName = fullName + " " + currentName
		}
	}
	return strings.Trim(fullName, " ")
}

func (p *BasePerson) FirstName(gender ...string) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	var suppliedGender string
	var firstName string
	if len(gender) > 0 {
		suppliedGender = gender[0]
		if suppliedGender == GenderMale {
			firstName = p.FirstNameMale()
		}
		if suppliedGender == GenderFemale {
			firstName = p.FirstNameFemale()
		}
	} else {
		randomIndex := rand.Intn(len(p.GetFirstNameFormats()))
		selectedMethod := p.GetFirstNameFormats()[randomIndex]
		firstName = p.invokeSelectedMethod(selectedMethod)
	}

	return firstName
}

// FirstNameMale returns a male first name for BasePerson
func (p *BasePerson) FirstNameMale() string {
	firstName, err := p.RandomElementFromStringSlice(p.GetMaleFirstNames())
	if err != nil {
		panic(err.Error())
	}
	return firstName
}

// FirstNameFemale returns a female first name for BasePerson
func (p *BasePerson) FirstNameFemale() string {
	firstName, err := p.BaseProvider.RandomElementFromStringSlice(p.GetFemaleFirstNames())
	if err != nil {
		panic(err.Error())
	}
	return firstName
}

// LastName returns a last name for BasePerson
func (p *BasePerson) LastName() string {
	lastName, err := p.BaseProvider.RandomElementFromStringSlice(p.GetLastNames())
	if err != nil {
		panic(err.Error())
	}
	return lastName
}

// TitleMale returns a male title for BasePerson
func (p *BasePerson) TitleMale() string {
	title, err := p.BaseProvider.RandomElementFromStringSlice(p.GetMaleTitles())
	if err != nil {
		panic(err.Error())
	}
	return title
}

// TitleFemale returns a female title for BasePerson
func (p *BasePerson) TitleFemale() string {
	title, err := p.BaseProvider.RandomElementFromStringSlice(p.GetFemaleTitles())
	if err != nil {
		panic(err.Error())
	}
	return title
}

func (p *BasePerson) Title(gender ...string) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	var suppliedGender string
	var title string
	if len(gender) > 0 {
		suppliedGender = gender[0]
		if suppliedGender == GenderMale {
			title = p.TitleMale()
		}
		if suppliedGender == GenderFemale {
			title = p.TitleFemale()
		}
	} else {
		formats := p.GetTitleFormats()
		randomIndex := rand.Intn(len(formats))
		selectedMethod := formats[randomIndex]
		title = p.invokeSelectedMethod(selectedMethod)
	}
	return title
}

// invokeSelectedMethod
func (p *BasePerson) invokeSelectedMethod(fn func(*BasePerson) string) string {
	return fn(p)
}
