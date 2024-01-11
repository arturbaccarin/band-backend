package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/arturbaccarin/band-backend/internal/dto"
	"github.com/arturbaccarin/band-backend/internal/entity"
	"github.com/arturbaccarin/band-backend/internal/infra/database"
	"github.com/arturbaccarin/band-backend/pkg/auth"
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
func (u UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var createUserParams dto.CreateUserParams

	err := json.NewDecoder(r.Body).Decode(&createUserParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := entity.NewUser(createUserParams.Name, createUserParams.Email, createUserParams.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = u.UserDB.Create(*user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := u.UserDB.FindByEmail(signInParams.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !user.ValidatePassword(signInParams.Password) {
		http.Error(w, "user or password is invalid", http.StatusForbidden)
		return
	}

	sub := strconv.FormatUint(uint64(user.ID), 10)

	tokenString, err := auth.GenerateJWT(sub)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	accessToken := dto.GetJWTOutput{AccessToken: tokenString}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(accessToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
