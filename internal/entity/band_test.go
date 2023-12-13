package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBand(t *testing.T) {
	band, err := NewBand("Iron Maiden", 1975)
	assert.NotNil(t, band)
	assert.Nil(t, err)
	assert.Equal(t, band.Name, "Iron Maiden")
	assert.Equal(t, band.Year, uint(1975))
}

func TestValidate(t *testing.T) {
	validBand := Band{
		Name: "Test Band",
		Year: 2000,
	}

	err := validBand.Validate()
	assert.Nil(t, err)

	emptyNameBand := Band{
		Name: "",
		Year: 2000,
	}
	err = emptyNameBand.Validate()
	assert.Equal(t, ErrNameIsEmpty, err)

	invalidYearBand := Band{
		Name: "Test Band",
		Year: 0,
	}
	err = invalidYearBand.Validate()
	assert.Equal(t, ErrYearIsInvalid, err)
}
