package providers

import "github.com/harmlessprince/goFaker/constants"

type BaseMedical struct {
	BaseProvider
}

func (m *BaseMedical) getBloodGroupFormat() string {
	return "{{BloodType}}{{BloodRh}}"
}

func (m *BaseMedical) BloodType() string {
	bloodType, err := m.BaseProvider.RandomElementFromStringSlice(constants.BloodTypes)
	if err != nil {
		panic(err.Error())
		return ""
	}
	return bloodType
}

func (m *BaseMedical) BloodRh() string {
	bloodRh, err := m.BaseProvider.RandomElementFromStringSlice(constants.BloodRhFactors)
	if err != nil {
		panic(err.Error())
		return ""
	}
	return bloodRh
}

func (m *BaseMedical) BloodGroup() string {
	bloodGroup, err := m.BaseProvider.Parse(m.getBloodGroupFormat(), m)
	if err != nil {
		panic(err.Error())
		return ""
	}
	return bloodGroup
}
