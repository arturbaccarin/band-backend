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

// Create godoc
//
//	@Summary	Create a band
//	@Tags		band
//	@Accept		json
//	@Produce	json
//	@Param		request	body	dto.CreateBandParams	true	"band request"
//	@Success	201
//	@Failure	500	{object}	ErrorResponse
//	@Router		/bands [post]
//	@Security	ApiKeyAuth
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

// GetByID godoc
//
//	@Summary	Get a band
//	@Tags		band
//	@Accept		json
//	@Produce	json
//	@Param		ID	path		int	true	"band ID"
//	@Success	200	{object}	entity.Band
//	@Failure	404
//	@Failure	500	{object}	ErrorResponse
//	@Router		/bands/{ID} [get]
//	@Security	ApiKeyAuth
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

// DeleteByID godoc
//
//	@Summary	Delete a band
//	@Tags		band
//	@Accept		json
//	@Produce	json
//	@Param		ID	path	string	true	"band ID"
//	@Success	204
//	@Failure	404
//	@Failure	500	{object}	ErrorResponse
//	@Router		/bands/{ID} [delete]
//	@Security	ApiKeyAuth
func (h *BandHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "ID")

	err := h.BandDB.DeleteByID(ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
	}

	w.WriteHeader(http.StatusNoContent)
}

// UpdateByID godoc
//
//	@Summary	Update a band
//	@Tags		band
//	@Accept		json
//	@Produce	json
//	@Param		ID		path	string					true	"band ID"
//	@Param		request	body	dto.UpdateBandParams	true	"band request"
//	@Success	204
//	@Failure	404
//	@Failure	500	{object}	ErrorResponse
//	@Router		/bands/{ID} [put]
//	@Security	ApiKeyAuth
func (h *BandHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	var updateBandParams dto.UpdateBandParams

	ID := chi.URLParam(r, "ID")

	err := json.NewDecoder(r.Body).Decode(&updateBandParams)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}

	band, err := entity.NewBand(updateBandParams.Name, updateBandParams.Year)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}

	err = h.BandDB.UpdateByID(ID, *band)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
}
