package dto

type CreateBandParams struct {
	Name string `json:"name"`
	Year uint   `json:"year"`
}

type UpdateBandParams struct {
	Name string `json:"name"`
	Year uint   `json:"year"`
}
