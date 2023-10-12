package providers

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

	if len(formats) > 0 {
		bc.formats = formats[0]
	} else {
		bc.formats = []string{
			"{{CompanyName}} {{CompanySuffix}}",
			"{{CompanyName}}",
		}
	}
}

func (bc *BaseCompany) GetCompanySuffix() []string {
	return bc.companySuffix
}

func (bc *BaseCompany) SetCompanySuffix(suffixes ...[]string) {
	if len(suffixes) > 0 {
		bc.companySuffix = suffixes[0]
	} else {
		bc.companySuffix = []string{
			"Ltd",
		}
	}
}

func (bc *BaseCompany) GetCompanyName() []string {
	return bc.companyName
}

func (bc *BaseCompany) SetCompanyName(names ...[]string) {
	if len(names) > 0 {
		bc.companyName = names[0]
	} else {
		bc.companyName = []string{
			"Percepta", "Defendify", "Intrepid Travel", "Aceable", "Exela", "Ibotta", "Kaboom Fireworks", "Compass Mortgage", "Marathon", "Technologent",
		}
	}
}

func (bc *BaseCompany) GetFormats() []string {
	return bc.formats
}

func (bc *BaseCompany) SetJobTitleFormats(formats ...[]string) {
	if len(formats) > 0 {
		bc.jobTitleFormat = formats[0]
	} else {
		bc.jobTitleFormat = []string{
			"{{Word}}",
		}
	}
}

func (bc *BaseCompany) GetJobTitleFormats() []string {
	return bc.jobTitleFormat
}

func (bc *BaseCompany) Company() string {
	format, err := bc.RandomElementFromStringSlice(bc.GetFormats())
	if err != nil {
		panic(err.Error())
		return ""
	}
	parse, err := bc.BaseProvider.Parse(format, bc)
	if err != nil {
		panic(err.Error())
		return ""
	}
	return parse
}

func (bc *BaseCompany) CompanySuffix() string {
	suffix, err := bc.RandomElementFromStringSlice(bc.GetCompanySuffix())
	if err != nil {
		panic(err.Error())
		return ""
	}
	return suffix
}
func (bc *BaseCompany) CompanyName() string {
	name, err := bc.RandomElementFromStringSlice(bc.GetCompanyName())
	if err != nil {
		panic(err.Error())
		return ""
	}
	return name
}

func (bc *BaseCompany) CompanyPrefix() string {
	prefix, err := bc.RandomElementFromStringSlice(bc.GetCompanySuffix())
	if err != nil {
		panic(err.Error())
		return ""
	}
	return prefix
}

func (bc *BaseCompany) JobTitle() string {
	format, err := bc.RandomElementFromStringSlice(bc.GetJobTitleFormats())
	if err != nil {
		panic(err.Error())
		return ""
	}
	parse, err := bc.BaseProvider.Parse(format, &BaseLorem{})
	if err != nil {
		panic(err.Error())
		return ""
	}
	return parse
}
