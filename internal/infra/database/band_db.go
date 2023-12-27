package database

import (
	"database/sql"

	"github.com/arturbaccarin/band-backend/internal/entity"
)

type Band struct {
	DB *sql.DB
}

func NewBand(db *sql.DB) *Band {
	return &Band{
		DB: db,
	}
}

func (b Band) SelectByID(ID uint) (entity.Band, error) {
	var band entity.Band

	query := `
		SELECT 
			* 
		FROM
			band
		WHERE
			id = ?; 
	`

	row := b.DB.QueryRow(query, ID)

	err := row.Scan(&band.ID, &band.Name, &band.Year)
	if err != nil {
		return band, err
	}

	return band, nil
}

func (b Band) Create(band entity.Band) error {
	query := `
		INSERT INTO band (name, year) 
		VALUES (?, ?);
	`

	_, err := b.DB.Exec(query, band.Name, band.Year)
	if err != nil {
		return err
	}

	return nil
}
