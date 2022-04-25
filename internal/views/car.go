package views

type CarDetailView struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Plate    string `json:"plate"`
	IsActive bool   `json:"is_active"`

	Client ClientView `json:"client"`
}

type CarListView struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Plate    string `json:"plate"`
	IsActive bool   `json:"is_active"`
	ClientID int64  `json:"client_id"`
}
