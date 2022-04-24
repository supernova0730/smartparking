package dtos

type TaxCreateUpdateDTO struct {
	Title    string `json:"title"`
	Duration string `json:"duration"`
	Price    int    `json:"price"`
}
