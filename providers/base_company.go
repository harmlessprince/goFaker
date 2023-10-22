package providers

import "log"

type BaseCompany struct {
	BaseProvider
	formats        []string
	companySuffix  []string
	companyName    []string
	jobTitleFormat []string
}

func NewBaseCompany() *BaseCompany {
	company := &BaseCompany{}
	company.SetCompanySuffix()
	company.SetFormats()
	company.SetCompanyName()
	company.SetJobTitleFormats()
	return company
}
func (bc *BaseCompany) SetFormats(formats ...[]string) {
	bc.formats = []string{
		"{{CompanyName}} {{CompanySuffix}}",
		"{{CompanyName}}",
	}
	if len(formats) > 0 {
		bc.formats = formats[0]
	}
}

func (bc *BaseCompany) GetCompanySuffix() []string {
	return bc.companySuffix
}

func (bc *BaseCompany) SetCompanySuffix(suffixes ...[]string) {
	bc.companySuffix = []string{
		"Ltd",
	}
	if len(suffixes) > 0 {
		bc.companySuffix = suffixes[0]
	}
}

func (bc *BaseCompany) GetCompanyName() []string {
	return bc.companyName
}

func (bc *BaseCompany) SetCompanyName(names ...[]string) {
	bc.companyName = []string{
		"Percepta", "Defendify", "Intrepid Travel", "Aceable", "Exela", "Ibotta", "Kaboom Fireworks", "Compass Mortgage", "Marathon", "Technologent",
	}
	if len(names) > 0 {
		bc.companyName = names[0]
	}
}

func (bc *BaseCompany) GetFormats() []string {
	return bc.formats
}

func (bc *BaseCompany) SetJobTitleFormats(formats ...[]string) {
	bc.jobTitleFormat = []string{
		"{{Word}}",
	}
	if len(formats) > 0 {
		bc.jobTitleFormat = formats[0]
	}
}

func (bc *BaseCompany) GetJobTitleFormats() []string {
	return bc.jobTitleFormat
}

func (bc *BaseCompany) Company() string {
	format, err := bc.RandomElementFromStringSlice(bc.GetFormats())
	if err != nil {
		log.Fatal(err.Error())
		return ""
	}
	parse, err := bc.BaseProvider.Parse(format, bc)
	if err != nil {
		log.Fatal(err.Error())
		return ""
	}
	return parse
}

func (bc *BaseCompany) CompanySuffix() string {
	suffix, err := bc.RandomElementFromStringSlice(bc.GetCompanySuffix())
	if err != nil {
		log.Fatal(err.Error())
		return ""
	}
	return suffix
}
func (bc *BaseCompany) CompanyName() string {
	name, err := bc.RandomElementFromStringSlice(bc.GetCompanyName())
	if err != nil {
		log.Fatal(err.Error())
	}
	return name
}

func (bc *BaseCompany) CompanyPrefix() string {
	prefix, err := bc.RandomElementFromStringSlice(bc.GetCompanySuffix())
	if err != nil {
		log.Fatal(err.Error())
	}
	return prefix
}

func (bc *BaseCompany) JobTitle() string {
	format, err := bc.RandomElementFromStringSlice(bc.GetJobTitleFormats())
	if err != nil {
		log.Fatal(err.Error())
	}
	parse, err := bc.BaseProvider.Parse(format, &BaseLorem{})
	if err != nil {
		log.Fatal(err.Error())
	}
	return parse
}
