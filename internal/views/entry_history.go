package views

type EntryHistoryDetailView struct {
	ID            int64  `json:"id"`
	Image         string `json:"image"`
	Time          string `json:"time"`
	CarID         int64  `json:"car_id"`
	ParkingZoneID int64  `json:"parking_zone_id"`

	Car         CarDetailView   `json:"car"`
	ParkingZone ParkingZoneView `json:"parking_zone"`
}

type EntryHistoryListView struct {
	ID               int64  `json:"id"`
	Time             string `json:"time"`
	CarID            int64  `json:"car_id"`
	ParkingZoneID    int64  `json:"parking_zone_id"`
	CarPlate         string `json:"car_plate"`
	ParkingZoneTitle string `json:"parking_zone_title"`
}

type PersonalEntryHistoryListView struct {
	ID               int64  `json:"id"`
	CarTitle         string `json:"car_title"`
	CarPlate         string `json:"car_plate"`
	EntryTime        string `json:"entry_time"`
	ParkingZoneTitle string `json:"parking_zone_title"`
}
