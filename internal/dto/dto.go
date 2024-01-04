package dto

type CreateBandParams struct {
	Name string `json:"name"`
	Year uint   `json:"year"`
}

type UpdateBandParams struct {
	Name string `json:"name"`
	Year uint   `json:"year"`
}

type CreateUserParams struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
