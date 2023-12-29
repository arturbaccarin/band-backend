package handler

import (
	"encoding/json"
	"net/http"

	"github.com/arturbaccarin/band-backend/internal/dto"
	"github.com/arturbaccarin/band-backend/internal/entity"
	"github.com/arturbaccarin/band-backend/internal/infra/database"
	"github.com/go-chi/chi"
)

type BandHandler struct {
	BandDB database.BandInterface
}

type ErrorResponse struct {
	Error string `json:"error"`
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
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}

	band, err := entity.NewBand(createBandParams.Name, createBandParams.Year)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}

	err = h.BandDB.Create(*band)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *BandHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "ID")
	band, err := h.BandDB.SelectByID(ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
	}

	if band.ID == 0 {
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(band)
}

func (h *BandHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "ID")

	err := h.BandDB.DeleteByID(ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
	}

	w.WriteHeader(http.StatusNoContent)
}
