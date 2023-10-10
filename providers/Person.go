package providers

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const GenderFemale string = "female"
const GenderMale string = "male"

type Person struct {
	BaseProvider
	titleFormats      []func(*Person) string
	firstNameFormats  []func(*Person) string
	maleNameFormats   [][]func(*Person) string
	femaleNameFormats [][]func(*Person) string
	maleFirstNames    []string
	femaleFirstNames  []string
	lastNames         []string
	maleTitles        []string
	femaleTitles      []string
	genderMale        string
	genderFemale      string
}

// GetTitleFormats is a method that provides access to the titleFormats.
func (p *Person) GetTitleFormats() []func(*Person) string {
	p.titleFormats = []func(*Person) string{
		(*Person).TitleMale,
		(*Person).TitleFemale,
	}
	return p.titleFormats
}

// GetFirstNameFormats is a method that provides access to the firstNameFormats.
func (p *Person) GetFirstNameFormats() []func(*Person) string {
	p.firstNameFormats = []func(*Person) string{
		(*Person).FirstNameMale, (*Person).FirstNameFemale,
	}
	return p.firstNameFormats
}

// GetMaleNameFormats is a method that provides access to the maleNameFormats.
func (p *Person) GetMaleNameFormats() [][]func(*Person) string {
	//p.maleNameFormats = []string{"{{ firstNameMale }} {{ lastName }}"}
	p.maleNameFormats = [][]func(*Person) string{
		{(*Person).FirstNameMale, (*Person).LastName},
	}
	return p.maleNameFormats
}

// GetFemaleNameFormats is a method that provides access to the femaleNameFormats.
func (p *Person) GetFemaleNameFormats() [][]func(*Person) string {
	p.femaleNameFormats = [][]func(*Person) string{
		{(*Person).FirstNameFemale, (*Person).LastName},
	}
	return p.femaleNameFormats
}

// GetMaleGender is a method that provides access to the MaleGender value.
func (p *Person) GetMaleGender() string {
	p.genderMale = "male"
	return p.genderMale
}

// GetFemaleGender is a method that provides access to the FemaleGender value.
func (p *Person) GetFemaleGender() string {
	p.genderFemale = "female"
	return p.genderFemale
}

func (p *Person) GetMaleFirstNames() []string {
	p.maleFirstNames = []string{"John", "Harry"}
	return p.maleFirstNames
}

func (p *Person) GetFemaleFirstNames() []string {
	p.femaleFirstNames = []string{"Jane", "Clara"}
	return p.femaleFirstNames
}

func (p *Person) GetLastNames() []string {
	p.lastNames = []string{"Doe", "Kane"}
	return p.lastNames
}

func (p *Person) GetMaleTitles() []string {
	p.maleFirstNames = []string{"Mr.", "Dr.", "Prof."}
	return p.maleFirstNames
}
func (p *Person) GetFemaleTitles() []string {
	p.femaleTitles = []string{"Mrs.", "Ms.", "Miss", "Dr.", "Prof."}
	return p.femaleTitles
}

func (p *Person) Name(gender ...string) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	var suppliedGender string
	var fullName string
	var formats []func(*Person) string
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
	fullName = parseToFullName(formats, fullName)
	return fullName
}

func parseToFullName(formats []func(*Person) string, fullName string) string {
	for i := range formats {
		selectedMethod := formats[i]
		currentName := invokeSelectedMethod(selectedMethod, &Person{})
		if len(fullName) < 0 {
			fullName = currentName
		} else {
			fullName = fullName + " " + currentName
		}
	}
	return strings.Trim(fullName, " ")
}

func (p *Person) FirstName(gender ...string) string {
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
		firstName = invokeSelectedMethod(selectedMethod, &Person{})
	}

	return firstName
}

// FirstNameMale returns a male first name for Person
func (p *Person) FirstNameMale() string {
	firstName, err := p.RandomElementFromStringSlice(p.GetMaleFirstNames())
	if err != nil {
		panic(err.Error())
	}
	return firstName
}

// FirstNameFemale returns a female first name for Person
func (p *Person) FirstNameFemale() string {
	firstName, err := p.BaseProvider.RandomElementFromStringSlice(p.GetFemaleFirstNames())
	if err != nil {
		panic(err.Error())
	}
	return firstName
}

// LastName returns a last name for Person
func (p *Person) LastName() string {
	lastName, err := p.BaseProvider.RandomElementFromStringSlice(p.GetLastNames())
	if err != nil {
		panic(err.Error())
	}
	return lastName
}

// TitleMale returns a male title for Person
func (p *Person) TitleMale() string {
	title, err := p.BaseProvider.RandomElementFromStringSlice(p.GetMaleTitles())
	if err != nil {
		panic(err.Error())
	}
	return title
}

// TitleFemale returns a female title for Person
func (p *Person) TitleFemale() string {
	title, err := p.BaseProvider.RandomElementFromStringSlice(p.GetFemaleTitles())
	if err != nil {
		panic(err.Error())
	}
	return title
}

func (p *Person) Title(gender ...string) string {
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
		randomIndex := rand.Intn(len(p.GetFirstNameFormats()))
		selectedMethod := p.GetTitleFormats()[randomIndex]
		fmt.Println(randomIndex)
		title = invokeSelectedMethod(selectedMethod, &Person{})
	}
	return title
}

// invokeSelectedMethod
func invokeSelectedMethod(fn func(*Person) string, personStruct *Person) string {
	return fn(personStruct)
}
