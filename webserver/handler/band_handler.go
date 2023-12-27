package handler

import (
	"encoding/json"
	"net/http"

	"github.com/arturbaccarin/band-backend/internal/dto"
	"github.com/arturbaccarin/band-backend/internal/entity"
	"github.com/arturbaccarin/band-backend/internal/infra/database"
)

type BandHandler struct {
	BandDB database.BandInterface
}

func NewBandHandler(bandDB database.BandInterface) *BandHandler {
	return &BandHandler{
		BandDB: bandDB,
	}
}

func (h *BandHandler) Create(w http.ResponseWriter, r *http.Request) {
	var createBandParams dto.CreateBandParams

	err := json.NewDecoder(r.Body).Decode(&createBandParams)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	band, err := entity.NewBand(createBandParams.Name, createBandParams.Year)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.BandDB.Create(*band)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
