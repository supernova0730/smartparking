package views

type ParkingPlaceDetailView struct {
	ID     string `json:"id,omitempty"`
	Number string `json:"number,omitempty"`
	IsBusy bool   `json:"is_busy"`

	ParkingZone ParkingZoneView `json:"parking_zone,omitempty"`
}

type ParkingPlaceListView struct {
	ID               string `json:"id,omitempty"`
	Number           string `json:"number,omitempty"`
	IsBusy           bool   `json:"is_busy"`
	ParkingZoneTitle string `json:"parking_zone_title,omitempty"`
	ParkingZoneID    string `json:"parking_zone_id,omitempty"`
}
