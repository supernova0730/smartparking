package views

type EntryHistoryDetailView struct {
	ID    int64  `json:"id,omitempty"`
	Image string `json:"image,omitempty"`
	Time  string `json:"time,omitempty"`

	Car         CarDetailView   `json:"car,omitempty"`
	ParkingZone ParkingZoneView `json:"parking_zone,omitempty"`
}

type EntryHistoryListView struct {
	ID               int64  `json:"id,omitempty"`
	Time             string `json:"time,omitempty"`
	CarID            int64  `json:"car_id,omitempty"`
	ParkingZoneID    int64  `json:"parking_zone_id,omitempty"`
	CarPlate         string `json:"car_plate,omitempty"`
	ParkingZoneTitle string `json:"parking_zone_title,omitempty"`
}

type PersonalEntryHistoryListView struct {
	ID               int64  `json:"id,omitempty"`
	CarTitle         string `json:"car_title,omitempty"`
	CarPlate         string `json:"car_plate,omitempty"`
	EntryTime        string `json:"entry_time,omitempty"`
	ParkingZoneTitle string `json:"parking_zone_title,omitempty"`
}
