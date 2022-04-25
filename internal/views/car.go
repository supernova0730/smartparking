package views

type CarDetailView struct {
	ID       int64  `json:"id,omitempty"`
	Title    string `json:"title,omitempty"`
	Plate    string `json:"plate,omitempty"`
	IsActive bool   `json:"is_active,omitempty"`

	Client ClientView `json:"client,omitempty"`
}

type CarListView struct {
	ID       int64  `json:"id,omitempty"`
	Title    string `json:"title,omitempty"`
	Plate    string `json:"plate,omitempty"`
	IsActive bool   `json:"is_active,omitempty"`
	ClientID int64  `json:"client_id,omitempty"`
}
