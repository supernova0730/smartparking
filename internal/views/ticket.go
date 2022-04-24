package views

type TicketDetailView struct {
	ID             int64  `json:"id"`
	ExpiresAt      string `json:"expires_at"`
	ClientID       int64  `json:"client_id"`
	ParkingPlaceID int64  `json:"parking_place_id"`

	Client       ClientView             `json:"client"`
	ParkingPlace ParkingPlaceDetailView `json:"parking_place"`
}

type TicketListView struct {
	ID               int64  `json:"id"`
	ExpiresAt        string `json:"expires_at"`
	ClientID         int64  `json:"client_id"`
	ParkingPlaceID   int64  `json:"parking_place_id"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	ParkingNumber    string `json:"parking_number"`
	ParkingZoneTitle string `json:"parking_zone_title"`
}

type PersonalTicketListView struct {
	ID               int64  `json:"id"`
	ParkingNumber    string `json:"parking_number"`
	ParkingZoneTitle string `json:"parking_zone_title"`
	ParkingZoneImage string `json:"parking_zone_image"`
	Expires          string `json:"expires"`
}
