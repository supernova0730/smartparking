package models

import (
	"smartparking/internal/consts"
	"smartparking/internal/dtos"
	"smartparking/internal/views"
	"time"
)

type ListTicket []Ticket
type Ticket struct {
	ID             int64     `gorm:"column:id"`
	ExpiresAt      time.Time `gorm:"column:expires_at"`
	CarID          int64     `gorm:"column:car_id"`
	ClientID       int64     `gorm:"column:client_id"`
	ParkingPlaceID int64     `gorm:"column:parking_place_id"`

	Client       *Client
	Car          *Car
	ParkingPlace *ParkingPlace
}

func (Ticket) TableName() string {
	return "ticket"
}

func (t Ticket) IsExpired() bool {
	return t.ExpiresAt.Before(time.Now())
}

func (t Ticket) ToView() views.TicketDetailView {
	view := views.TicketDetailView{
		ID:             t.ID,
		ExpiresAt:      t.ExpiresAt.Format(consts.CustomFormat),
		ClientID:       t.ClientID,
		CarID:          t.CarID,
		ParkingPlaceID: t.ParkingPlaceID,
	}
	if t.Client != nil {
		view.Client = t.Client.ToView()
	}
	if t.Car != nil {
		view.Car = t.Car.ToView()
	}
	if t.ParkingPlace != nil {
		view.ParkingPlace = t.ParkingPlace.ToView()
	}
	return view
}

func (list ListTicket) ToView() (result []views.TicketListView) {
	for _, t := range list {
		view := views.TicketListView{
			ID:             t.ID,
			ExpiresAt:      t.ExpiresAt.Format(consts.CustomFormat),
			ClientID:       t.ClientID,
			CarID:          t.CarID,
			ParkingPlaceID: t.ParkingPlaceID,
		}
		if t.Client != nil {
			view.FirstName = t.Client.ToView().FirstName
			view.LastName = t.Client.ToView().LastName
		}
		if t.Car != nil {
			view.CarTitle = t.Car.Title
			view.CarPlate = t.Car.Plate
		}
		if t.ParkingPlace != nil {
			view.ParkingNumber = t.ParkingPlace.ToView().Number
			if t.ParkingPlace.ParkingZone != nil {
				view.ParkingZoneTitle = t.ParkingPlace.ParkingZone.ToView().Title
			}
		}
		result = append(result, view)
	}
	return
}

func (list ListTicket) ToPersonalTicketListView() (result []views.PersonalTicketListView) {
	for _, t := range list {
		view := views.PersonalTicketListView{
			ID:      t.ID,
			Expires: t.ExpiresAt.Format(consts.CustomFormat),
		}
		if t.ParkingPlace != nil {
			view.ParkingNumber = t.ParkingPlace.Number
		}
		if t.ParkingPlace != nil && t.ParkingPlace.ParkingZone != nil {
			view.ParkingZoneTitle = t.ParkingPlace.ParkingZone.Title
			view.ParkingZoneImage = t.ParkingPlace.ParkingZone.Image
		}
		result = append(result, view)
	}
	return
}

func (t *Ticket) SetFromUpdateDTO(dto dtos.TicketUpdateDTO) error {
	expiresAt, err := time.Parse(consts.CustomFormat, dto.ExpiresAt)
	if err != nil {
		return err
	}
	t.ExpiresAt = expiresAt
	t.ClientID = dto.ClientID
	t.ParkingPlaceID = dto.ParkingPlaceID
	return nil
}
