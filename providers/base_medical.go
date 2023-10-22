package providers

import (
	"github.com/harmlessprince/goFaker/constants"
	"log"
)

type BaseMedical struct {
	BaseProvider
}

func (m *BaseMedical) getBloodGroupFormat() string {
	return "{{BloodType}}{{BloodRh}}"
}

func (m *BaseMedical) BloodType() string {
	bloodType, err := m.BaseProvider.RandomElementFromStringSlice(constants.BloodTypes)
	if err != nil {
		log.Fatal(err.Error())
	}
	return bloodType
}

func (m *BaseMedical) BloodRh() string {
	bloodRh, err := m.BaseProvider.RandomElementFromStringSlice(constants.BloodRhFactors)
	if err != nil {
		log.Fatal(err.Error())
	}
	return bloodRh
}

func (m *BaseMedical) BloodGroup() string {
	bloodGroup, err := m.BaseProvider.Parse(m.getBloodGroupFormat(), m)
	if err != nil {
		log.Fatal(err.Error())
	}
	return bloodGroup
}
