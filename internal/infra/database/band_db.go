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

func (b Band) SelectByID(ID string) (entity.Band, error) {
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

func (b Band) DeleteByID(ID string) error {
	query := `
		DELETE FROM band
		WHERE id = ?;
	`

	_, err := b.DB.Exec(query, ID)
	if err != nil {
		return err
	}

	return nil
}

func (b Band) UpdateByID(ID string, band entity.Band) error {
	query := `
		UPDATE band
		SET name = ?, year = ?
		WHERE id = ?;
	`

	_, err := b.DB.Exec(query, band.Name, band.Year, ID)
	if err != nil {
		return err
	}

	return nil
}

func (b Band) GetList(page int) ([]entity.Band, error) {
	var bands []entity.Band

	query := `
		SELECT 
			*
		FROM
			band
		LIMIT 5
		OFFSET ?;
	`

	offset := 5 * (page - 1)

	rows, err := b.DB.Query(query, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var band entity.Band

		err := rows.Scan(&band.ID, &band.Name, &band.Year)
		if err != nil {
			return nil, err
		}

		bands = append(bands, band)
	}

	return bands, nil
}
