package database

import (
	"database/sql"

	"github.com/arturbaccarin/bands-backend/internal/entity"
)

type Band struct {
	DB *sql.DB
}

func NewBand(db *sql.DB) *Band {
	return &Band{
		DB: db,
	}
}

func (b Band) SelectByID(ID int) (entity.Band, error) {
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

func (b Band) Create(name string, year int) error {
	query := `
		INSERT INTO band (name, year) 
		VALUES (?, ?);
	`

	_, err := b.DB.Exec(query, name, year)
	if err != nil {
		return err
	}

	return nil
}
