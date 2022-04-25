package views

type TaxView struct {
	ID       string `json:"id,omitempty"`
	Title    string `json:"title,omitempty"`
	Duration string `json:"duration,omitempty"`
	Price    string `json:"price,omitempty"`
}
