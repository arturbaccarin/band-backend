package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/arturbaccarin/band-backend/internal/dto"
	"github.com/arturbaccarin/band-backend/internal/entity"
	"github.com/arturbaccarin/band-backend/internal/infra/database"
	"github.com/arturbaccarin/band-backend/pkg/auth"
)

var ErrWrongPassword = errors.New("user or password is invalid")

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
func (u UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var createUserParams dto.CreateUserParams

	err := json.NewDecoder(r.Body).Decode(&createUserParams)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	user, err := entity.NewUser(createUserParams.Name, createUserParams.Email, createUserParams.Password)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	err = u.UserDB.Create(*user)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Create godoc
//
//	@Summary	SignIn
//	@Tags		user
//	@Accept		json
//	@Produce	json
//	@Param		request	body	dto.SignInParams	true	"sign in params"
//	@Success      200  {object}  dto.GetJWTOutput
//	@Failure	500	{object}	ErrorResponse
//	@Router		/users/signin [post]
func (u UserHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var signInParams dto.SignInParams

	err := json.NewDecoder(r.Body).Decode(&signInParams)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	user, err := u.UserDB.FindByEmail(signInParams.Email)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if !user.ValidatePassword(signInParams.Password) {
		ErrorResponse(w, http.StatusBadRequest, ErrWrongPassword)
		return
	}

	sub := strconv.FormatUint(uint64(user.ID), 10)

	tokenString, err := auth.GenerateJWT(sub)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	accessToken := dto.GetJWTOutput{AccessToken: tokenString}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(accessToken)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}
}
