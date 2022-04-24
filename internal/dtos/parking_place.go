package dtos

type ParkingPlaceCreateDTO struct {
	Number        string `json:"number"`
	ParkingZoneID int64  `json:"parking_zone_id"`
}

type ParkingPlaceUpdateDTO struct {
	Number        string `json:"number"`
	IsBusy        bool   `json:"is_busy"`
	ParkingZoneID int64  `json:"parking_zone_id"`
}
