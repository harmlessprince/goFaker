package providers

type CompanyInterface interface {
	Company() (string, error)
	CompanySuffix() (string, error)
	JobTitle() (string, error)
}

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

// Company generates a random company name.
//
// Returns:
// - A string representing a randomly generated company name.
func (bc *BaseCompany) Company() (string, error) {
	format := helperInstance.RandomElementString(bc.GetFormats())
	parse, err := bc.BaseProvider.Parse(format, bc)
	if err != nil {
		return "", err
	}
	return parse, nil
}

// CompanySuffix generates a random company suffix.
//
// Returns:
// - A string representing a randomly generated company suffix.
func (bc *BaseCompany) CompanySuffix() (string, error) {
	suffix := helperInstance.RandomElementString(bc.GetCompanySuffix())
	return suffix, nil
}
func (bc *BaseCompany) CompanyName() (string, error) {
	name := helperInstance.RandomElementString(bc.GetCompanyName())
	return name, nil
}

func (bc *BaseCompany) CompanyPrefix() (string, error) {
	prefix := helperInstance.RandomElementString(bc.GetCompanySuffix())
	return prefix, nil
}

// JobTitle generates a random job title.
//
// Returns:
// - A string representing a randomly generated job title.
func (bc *BaseCompany) JobTitle() (string, error) {
	format := helperInstance.RandomElementString(bc.GetJobTitleFormats())
	parse, err := bc.BaseProvider.Parse(format, &BaseLorem{})
	if err != nil {
		return "", err
	}
	return parse, nil
}
