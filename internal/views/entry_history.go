package views

type EntryHistoryDetailView struct {
	ID    string `json:"id,omitempty"`
	Image string `json:"image,omitempty"`
	Time  string `json:"time,omitempty"`

	Car         CarDetailView   `json:"car,omitempty"`
	ParkingZone ParkingZoneView `json:"parking_zone,omitempty"`
}

type EntryHistoryListView struct {
	ID               string `json:"id,omitempty"`
	Time             string `json:"time,omitempty"`
	CarPlate         string `json:"car_plate,omitempty"`
	ParkingZoneTitle string `json:"parking_zone_title,omitempty"`
	CarID            string `json:"car_id,omitempty"`
	ParkingZoneID    string `json:"parking_zone_id,omitempty"`
}

type PersonalEntryHistoryListView struct {
	ID               string `json:"id,omitempty"`
	CarTitle         string `json:"car_title,omitempty"`
	CarPlate         string `json:"car_plate,omitempty"`
	EntryTime        string `json:"entry_time,omitempty"`
	ParkingZoneTitle string `json:"parking_zone_title,omitempty"`
}
