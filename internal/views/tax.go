package views

type TaxView struct {
	ID       int64  `json:"id,omitempty"`
	Title    string `json:"title,omitempty"`
	Duration string `json:"duration,omitempty"`
	Price    int    `json:"price,omitempty"`
}
