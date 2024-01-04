package handler

import (
	"encoding/json"
	"net/http"

	"github.com/arturbaccarin/band-backend/internal/dto"
	"github.com/arturbaccarin/band-backend/internal/entity"
	"github.com/arturbaccarin/band-backend/internal/infra/database"
)

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHander(userDB database.UserInterface) *UserHandler {
	return &UserHandler{
		UserDB: userDB,
	}
}

// Create godoc
//
//	@Summary	Create an user
//	@Tags		user
//	@Accept		json
//	@Produce	json
//	@Param		request	body	dto.CreateUserParams	true	"user request"
//	@Success	201
//	@Failure	500	{object}	ErrorResponse
//	@Router		/users [post]
//	@Security	ApiKeyAuth
func (u UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var createUserParams dto.CreateUserParams

	err := json.NewDecoder(r.Body).Decode(&createUserParams)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}

	user, err := entity.NewUser(createUserParams.Name, createUserParams.Email, createUserParams.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}

	err = u.UserDB.Create(*user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
}
