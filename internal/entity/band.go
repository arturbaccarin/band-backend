package entity

import "errors"

var (
	ErrNameIsEmpty   = errors.New("name is empty")
	ErrYearIsInvalid = errors.New("year is invalid")
)

type Band struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Year uint   `json:"year"`
}

func NewBand(name string, year uint) (*Band, error) {
	band := &Band{
		Name: name,
		Year: year,
	}

	err := band.Validate()
	if err != nil {
		return nil, err
	}

	return band, nil
}

func (b *Band) Validate() error {
	if b.Name == "" {
		return ErrNameIsEmpty
	}

	if b.Year <= 0 {
		return ErrYearIsInvalid
	}

	return nil
}
