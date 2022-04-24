package views

type TaxView struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Duration string `json:"duration"`
	Price    int    `json:"price"`
}
