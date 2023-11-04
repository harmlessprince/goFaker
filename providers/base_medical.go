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

// BloodType generates a random blood type (A, B, AB, O).
//
// Returns:
// - A string representing a randomly generated blood type (A, B, AB, O).
func (m *BaseMedical) BloodType() (string, error) {
	bloodType, err := m.BaseProvider.RandomElementFromStringSlice(constants.BloodTypes)
	if err != nil {
		return "", err
	}
	return bloodType, nil
}

// BloodRh generates a random blood Rh factor (+ or -).
//
// Returns:
// - A string representing a randomly generated blood Rh factor, either "+" or "-".
func (m *BaseMedical) BloodRh() (string, error) {
	bloodRh, err := m.BaseProvider.RandomElementFromStringSlice(constants.BloodRhFactors)
	if err != nil {
		return "", err
	}
	return bloodRh, nil
}

// BloodGroup generates a random blood group, including the ABO blood type (A, B, AB, O) and the Rh factor (+ or -).
//
// Returns:
// - A string representing a randomly generated blood group, such as "A+", "B-", "AB+", "O-", etc.
func (m *BaseMedical) BloodGroup() (string, error) {
	bloodGroup, err := m.BaseProvider.Parse(m.getBloodGroupFormat(), m)
	if err != nil {
		return "", err
	}
	return bloodGroup, nil
}
