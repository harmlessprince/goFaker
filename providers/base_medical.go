package providers

import (
	"github.com/harmlessprince/goFaker/constants"
)

type BloodInterface interface {
	BloodType() (string, error)
	BloodRh() (string, error)
	BloodGroup() (string, error)
}
type BaseMedical struct {
	BaseProvider
}

func (m *BaseMedical) getBloodGroupFormat() string {
	return "{{BloodType}}{{BloodRh}}"
}

func (m *BaseMedical) BloodType() (string, error) {
	bloodType, err := m.BaseProvider.RandomElementFromStringSlice(constants.BloodTypes)
	if err != nil {
		return "", err
	}
	return bloodType, nil
}

func (m *BaseMedical) BloodRh() (string, error) {
	bloodRh, err := m.BaseProvider.RandomElementFromStringSlice(constants.BloodRhFactors)
	if err != nil {
		return "", err
	}
	return bloodRh, nil
}

func (m *BaseMedical) BloodGroup() (string, error) {
	bloodGroup, err := m.BaseProvider.Parse(m.getBloodGroupFormat(), m)
	if err != nil {
		return "", err
	}
	return bloodGroup, nil
}
