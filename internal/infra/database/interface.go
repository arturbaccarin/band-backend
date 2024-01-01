package database

import "github.com/arturbaccarin/band-backend/internal/entity"

type BandInterface interface {
	Create(entity.Band) error
	SelectByID(string) (entity.Band, error)
	DeleteByID(string) error
	UpdateByID(string, entity.Band) error
	GetList(int) ([]entity.Band, error)
}
