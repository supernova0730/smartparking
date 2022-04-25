package views

type TicketDetailView struct {
	ID             string `json:"id,omitempty"`
	ExpiresAt      string `json:"expires_at,omitempty"`
	ClientID       string `json:"client_id,omitempty"`
	ParkingPlaceID string `json:"parking_place_id,omitempty"`

	Client       ClientView             `json:"client,omitempty"`
	ParkingPlace ParkingPlaceDetailView `json:"parking_place,omitempty"`
}

type TicketListView struct {
	ID               string `json:"id,omitempty"`
	ExpiresAt        string `json:"expires_at,omitempty"`
	ClientID         string `json:"client_id,omitempty"`
	ParkingPlaceID   string `json:"parking_place_id,omitempty"`
	FirstName        string `json:"first_name,omitempty"`
	LastName         string `json:"last_name,omitempty"`
	ParkingNumber    string `json:"parking_number,omitempty"`
	ParkingZoneTitle string `json:"parking_zone_title,omitempty"`
}

type PersonalTicketListView struct {
	ID               string `json:"id,omitempty"`
	ParkingNumber    string `json:"parking_number,omitempty"`
	ParkingZoneTitle string `json:"parking_zone_title,omitempty"`
	ParkingZoneImage string `json:"parking_zone_image,omitempty"`
	Expires          string `json:"expires,omitempty"`
}
