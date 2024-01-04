package database

import (
	"database/sql"

	"github.com/arturbaccarin/band-backend/internal/entity"
)

type User struct {
	DB *sql.DB
}

func NewUser(db *sql.DB) *User {
	return &User{
		DB: db,
	}
}

func (u User) Create(user entity.User) error {
	query := `
		INSERT INTO user (name, email, password)
		VALUE (?, ?, ?);
	`

	_, err := u.DB.Exec(query, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (u User) FindByEmail(email string) (entity.User, error) {
	var user entity.User

	query := `
		SELECT 
			*
		FROM
			user
		WHERE
			email = ?;
	`

	row := u.DB.QueryRow(query, email)

	err := row.Scan(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}
