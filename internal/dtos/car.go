package dtos

type CarCreateUpdateDTO struct {
	Title    string `json:"title"`
	Plate    string `json:"plate"`
	IsActive bool   `json:"is_active"`
	ClientID int64  `json:"client_id"`
}

type PersonalCarCreateDTO struct {
	Title string `json:"title" validate:"required"`
	Plate string `json:"plate" validate:"required"`
}
