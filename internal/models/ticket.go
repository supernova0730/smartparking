package models

import (
	"smartparking/internal/consts"
	"smartparking/internal/dtos"
	"smartparking/internal/views"
	"smartparking/pkg/tools"
	"time"
)

type ListTicket []Ticket
type Ticket struct {
	ID             int64     `gorm:"column:id"`
	ExpiresAt      time.Time `gorm:"column:expires_at"`
	ClientID       int64     `gorm:"column:client_id"`
	ParkingPlaceID int64     `gorm:"column:parking_place_id"`

	Client       *Client
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
		ID:             tools.Int64ToString(t.ID),
		ExpiresAt:      t.ExpiresAt.Format(consts.CustomFormat),
		ClientID:       tools.Int64ToString(t.ClientID),
		ParkingPlaceID: tools.Int64ToString(t.ParkingPlaceID),
	}
	if t.Client != nil {
		view.Client = t.Client.ToView()
	}
	if t.ParkingPlace != nil {
		view.ParkingPlace = t.ParkingPlace.ToView()
	}
	return view
}

func (list ListTicket) ToView() (result []views.TicketListView) {
	for _, t := range list {
		view := views.TicketListView{
			ID:             tools.Int64ToString(t.ID),
			ExpiresAt:      t.ExpiresAt.Format(consts.CustomFormat),
			ClientID:       tools.Int64ToString(t.ClientID),
			ParkingPlaceID: tools.Int64ToString(t.ParkingPlaceID),
		}
		if t.Client != nil {
			view.ClientFullName = t.Client.GetFullName()
		}
		if t.ParkingPlace != nil {
			view.ParkingNumber = t.ParkingPlace.ToView().Number
			if t.ParkingPlace.ParkingZone != nil {
				view.ParkingZoneTitle = t.ParkingPlace.ParkingZone.ToView().Title
				view.ParkingZoneID = tools.Int64ToString(t.ParkingPlace.ParkingZone.ID)
			}
		}
		result = append(result, view)
	}
	return
}

func (list ListTicket) ToPersonalTicketListView() (result []views.PersonalTicketListView) {
	for _, t := range list {
		view := views.PersonalTicketListView{
			ID:      tools.Int64ToString(t.ID),
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
