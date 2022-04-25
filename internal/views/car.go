package views

type CarDetailView struct {
	ID       string `json:"id,omitempty"`
	Title    string `json:"title,omitempty"`
	Plate    string `json:"plate,omitempty"`
	IsActive bool   `json:"is_active,omitempty"`

	Client ClientView `json:"client,omitempty"`
}

type CarListView struct {
	ID             string `json:"id,omitempty"`
	Title          string `json:"title,omitempty"`
	Plate          string `json:"plate,omitempty"`
	IsActive       bool   `json:"is_active,omitempty"`
	ClientFullName string `json:"client_full_name,omitempty"`
	ClientID       string `json:"client_id,omitempty"`
}
