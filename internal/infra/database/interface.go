package database

import "github.com/arturbaccarin/band-backend/internal/entity"

type BandInterface interface {
	Create(band entity.Band) error
	SelectByID(ID string) (entity.Band, error)
	DeleteByID(ID string) error
	UpdateByID(ID string, band entity.Band) error
	GetList(page int) ([]entity.Band, error)
}

type UserInterface interface {
	Create(user entity.User) error
	FindByEmail(email, password string) (entity.User, error)
}
