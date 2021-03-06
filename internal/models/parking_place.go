package models

import (
	"smartparking/internal/dtos"
	"smartparking/internal/views"
	"smartparking/pkg/tools"
)

type ListParkingPlace []ParkingPlace
type ParkingPlace struct {
	ID            int64  `gorm:"column:id"`
	Number        string `gorm:"column:number"`
	IsBusy        bool   `gorm:"column:is_busy"`
	ParkingZoneID int64  `gorm:"column:parking_zone_id"`

	ParkingZone *ParkingZone
}

func (ParkingPlace) TableName() string {
	return "parking_place"
}

func (pp ParkingPlace) ToView() views.ParkingPlaceDetailView {
	view := views.ParkingPlaceDetailView{
		ID:     tools.Int64ToString(pp.ID),
		Number: pp.Number,
		IsBusy: pp.IsBusy,
	}
	if pp.ParkingZone != nil {
		view.ParkingZone = pp.ParkingZone.ToView()
	}
	return view
}

func (list ListParkingPlace) ToView() (result []views.ParkingPlaceListView) {
	for _, parkingPlace := range list {
		view := views.ParkingPlaceListView{
			ID:            tools.Int64ToString(parkingPlace.ID),
			Number:        parkingPlace.Number,
			IsBusy:        parkingPlace.IsBusy,
			ParkingZoneID: tools.Int64ToString(parkingPlace.ParkingZoneID),
		}
		if parkingPlace.ParkingZone != nil {
			view.ParkingZoneTitle = parkingPlace.ParkingZone.Title
		}
		result = append(result, view)
	}
	return
}

func (pp *ParkingPlace) SetFromCreateDTO(dto dtos.ParkingPlaceCreateDTO) {
	pp.Number = dto.Number
	pp.ParkingZoneID = dto.ParkingZoneID
}

func (pp *ParkingPlace) SetFromUpdateDTO(dto dtos.ParkingPlaceUpdateDTO) {
	pp.Number = dto.Number
	pp.IsBusy = dto.IsBusy
	pp.ParkingZoneID = dto.ParkingZoneID
}
