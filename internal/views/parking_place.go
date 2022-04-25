package views

type ParkingPlaceDetailView struct {
	ID     int64  `json:"id,omitempty"`
	Number string `json:"number,omitempty"`
	IsBusy bool   `json:"is_busy,omitempty"`

	ParkingZone ParkingZoneView `json:"parking_zone,omitempty"`
}

type ParkingPlaceListView struct {
	ID               int64  `json:"id,omitempty"`
	Number           string `json:"number,omitempty"`
	IsBusy           bool   `json:"is_busy,omitempty"`
	ParkingZoneTitle string `json:"parking_zone_title,omitempty"`
}
