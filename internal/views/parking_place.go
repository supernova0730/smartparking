package views

type ParkingPlaceDetailView struct {
	ID            int64  `json:"id"`
	Number        string `json:"number"`
	IsBusy        bool   `json:"is_busy"`
	ParkingZoneID int64  `json:"parking_zone_id"`

	ParkingZone ParkingZoneView `json:"parking_zone"`
}

type ParkingPlaceListView struct {
	ID            int64  `json:"id"`
	Number        string `json:"number"`
	IsBusy        bool   `json:"is_busy"`
	ParkingZoneID int64  `json:"parking_zone_id"`
}
