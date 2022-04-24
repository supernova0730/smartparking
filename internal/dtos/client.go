package dtos

type ClientUpdateDTO struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone" validate:"len=10"`
}
