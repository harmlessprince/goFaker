package tests

import (
	"github.com/harmlessprince/goFaker/providers"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

var addressInstanceTest = providers.NewBaseAddress()

//func TestName(t *testing.T) {
//	t.Run("description", func(t *testing.T) {})
//	t.Run("description", func(t *testing.T) {})
//}

func TestCityName(t *testing.T) {
	t.Run("Validate CityName is generated", func(t *testing.T) {

		cityName, _ := addressInstanceTest.CityName()
		assert.True(t, len(cityName) > 0)
	})
}

func TestCity(t *testing.T) {
	t.Run("Validate City is generated", func(t *testing.T) {
		city, _ := addressInstanceTest.City()
		assert.True(t, len(city) > 0)
	})
}
func TestCityPrefix(t *testing.T) {
	t.Run("Validate CityPrefix is generated", func(t *testing.T) {
		expected, _ := addressInstanceTest.CityPrefix()
		assert.True(t, len(expected) > 0)
	})
}

func TestStreetPrefix(t *testing.T) {
	t.Run("Validate street prefix is generated", func(t *testing.T) {
		expected, _ := addressInstanceTest.StreetPrefix()
		assert.True(t, len(expected) > 0)
	})
}

func TestCountry(t *testing.T) {
	t.Run("Validate Country is generated", func(t *testing.T) {
		expected, _ := addressInstanceTest.Country()
		assert.True(t, len(expected) > 0)
	})
}

func TestStreetAddress(t *testing.T) {
	t.Run("Validate StreetAddress is generated", func(t *testing.T) {
		expected, _ := addressInstanceTest.StreetAddress()
		assert.True(t, len(expected) > 0)
	})
}
func TestStreetName(t *testing.T) {
	t.Run("Validate StreetName is generated", func(t *testing.T) {
		expected, _ := addressInstanceTest.StreetName()
		assert.True(t, len(expected) > 0)
	})
}

func TestLatitude(t *testing.T) {
	t.Run("Validate Latitude is generated", func(t *testing.T) {
		expected, _ := addressInstanceTest.Latitude()
		assert.True(t, expected <= 90 && expected >= -90)
	})
}
func TestLongitude(t *testing.T) {
	t.Run("Validate Valid Longitude is generated", func(t *testing.T) {
		expected, _ := addressInstanceTest.Longitude()
		assert.True(t, expected <= 180 && expected >= -180)
	})
}

func TestAddress(t *testing.T) {
	t.Run("Validate Valid Address is generated", func(t *testing.T) {
		expected, _ := addressInstanceTest.Address()
		assert.True(t, len(expected) > 0)
	})
}
