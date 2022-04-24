package dtos

type TicketUpdateDTO struct {
	ExpiresAt      string `json:"expires_at"`
	ClientID       int64  `json:"client_id"`
	ParkingPlaceID int64  `json:"parking_place_id"`
}

type BuyTicketDTO struct {
	TaxID          int64 `json:"tax_id" validate:"required"`
	ParkingPlaceID int64 `json:"parking_place_id" validate:"required"`
}
