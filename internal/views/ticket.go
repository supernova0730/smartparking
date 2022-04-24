package views

type TicketDetailView struct {
	ID             int64  `json:"id"`
	ExpiresAt      string `json:"expires_at"`
	ClientID       int64  `json:"client_id"`
	CarID          int64  `json:"car_id"`
	ParkingPlaceID int64  `json:"parking_place_id"`

	Client       ClientView             `json:"client"`
	Car          CarDetailView          `json:"car"`
	ParkingPlace ParkingPlaceDetailView `json:"parking_place"`
}

type TicketListView struct {
	ID               int64  `json:"id"`
	ExpiresAt        string `json:"expires_at"`
	ClientID         int64  `json:"client_id"`
	ParkingPlaceID   int64  `json:"parking_place_id"`
	CarID            int64  `json:"car_id"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	ParkingNumber    string `json:"parking_number"`
	ParkingZoneTitle string `json:"parking_zone_title"`
	CarTitle         string `json:"car_title"`
	CarPlate         string `json:"car_number"`
}

type PersonalTicketListView struct {
	ID               int64  `json:"id"`
	ParkingNumber    string `json:"parking_number"`
	ParkingZoneTitle string `json:"parking_zone_title"`
	ParkingZoneImage string `json:"parking_zone_image"`
	Expires          string `json:"expires"`
}
