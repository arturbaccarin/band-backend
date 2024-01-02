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
