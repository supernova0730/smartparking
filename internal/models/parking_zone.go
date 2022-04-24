package models

import (
	"smartparking/internal/dtos"
	"smartparking/internal/views"
)

type ListParkingZone []ParkingZone
type ParkingZone struct {
	ID    int64  `gorm:"column:id"`
	Title string `gorm:"column:title"`
	Image string `gorm:"column:image"`
}

func (ParkingZone) TableName() string {
	return "parking_zone"
}

func (pz ParkingZone) ToView() views.ParkingZoneView {
	return views.ParkingZoneView{
		ID:    pz.ID,
		Title: pz.Title,
		Image: pz.Image,
	}
}

func (list ListParkingZone) ToView() (result []views.ParkingZoneView) {
	for _, parkingZone := range list {
		result = append(result, parkingZone.ToView())
	}
	return
}

func (pz *ParkingZone) SetFromCreateUpdateDTO(dto dtos.ParkingZoneCreateUpdateDTO) {
	pz.Title = dto.Title
	pz.Image = dto.Image
}
