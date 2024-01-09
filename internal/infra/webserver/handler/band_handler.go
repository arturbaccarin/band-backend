package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/arturbaccarin/band-backend/internal/dto"
	"github.com/arturbaccarin/band-backend/internal/entity"
	"github.com/arturbaccarin/band-backend/internal/infra/database"
	"github.com/go-chi/chi"
)

type BandHandler struct {
	BandDB database.BandInterface
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
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	band, err := entity.NewBand(createBandParams.Name, createBandParams.Year)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	err = h.BandDB.Create(*band)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
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
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if band.ID == 0 {
		ErrorResponse(w, http.StatusNoContent, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(band)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}
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
		ErrorResponse(w, http.StatusBadRequest, err)
		return
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
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	band, err := entity.NewBand(updateBandParams.Name, updateBandParams.Year)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	err = h.BandDB.UpdateByID(ID, *band)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetList godoc
//
//	@Summary	Get a list of bands
//	@Tags		band
//	@Accept		json
//	@Produce	json
//	@Param		page	query		string	false	"page number"
//	@Success	200		{object}	[]entity.Band
//	@Failure	404
//	@Failure	500	{object}	ErrorResponse
//	@Router		/bands [get]
//	@Security	ApiKeyAuth
func (h *BandHandler) GetList(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	bands, err := h.BandDB.GetList(pageInt)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(bands)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}
}
