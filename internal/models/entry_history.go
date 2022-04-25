package models

import (
	"smartparking/internal/consts"
	"smartparking/internal/views"
	"smartparking/pkg/tools"
	"time"
)

type ListEntryHistory []EntryHistory
type EntryHistory struct {
	ID            int64     `gorm:"column:id"`
	Image         string    `gorm:"column:image"`
	Time          time.Time `gorm:"column:time"`
	CarID         int64     `gorm:"column:car_id"`
	ParkingZoneID int64     `gorm:"column:parking_zone_id"`

	Car         *Car
	ParkingZone *ParkingZone
}

func (EntryHistory) TableName() string {
	return "entry_history"
}

func (eh EntryHistory) ToView() views.EntryHistoryDetailView {
	view := views.EntryHistoryDetailView{
		ID:    tools.Int64ToString(eh.ID),
		Image: eh.Image,
		Time:  eh.Time.Format(consts.CustomFormat),
	}
	if eh.Car != nil {
		view.Car = eh.Car.ToView()
	}
	if eh.ParkingZone != nil {
		view.ParkingZone = eh.ParkingZone.ToView()
	}
	return view
}

func (list ListEntryHistory) ToView() (result []views.EntryHistoryListView) {
	for _, eh := range list {
		view := views.EntryHistoryListView{
			ID:            tools.Int64ToString(eh.ID),
			Time:          eh.Time.Format(consts.CustomFormat),
			CarID:         tools.Int64ToString(eh.CarID),
			ParkingZoneID: tools.Int64ToString(eh.ParkingZoneID),
		}
		if eh.Car != nil {
			view.CarPlate = eh.Car.ToView().Plate
		}
		if eh.ParkingZone != nil {
			view.ParkingZoneTitle = eh.ParkingZone.ToView().Title
		}
		result = append(result, view)
	}
	return
}

func (list ListEntryHistory) ToPersonalEntryHistoryView() (result []views.PersonalEntryHistoryListView) {
	for _, eh := range list {
		view := views.PersonalEntryHistoryListView{
			ID:        tools.Int64ToString(eh.ID),
			EntryTime: eh.Time.Format(consts.CustomFormat),
		}
		if eh.Car != nil {
			view.CarTitle = eh.Car.Title
			view.CarPlate = eh.Car.Plate
		}
		if eh.ParkingZone != nil {
			view.ParkingZoneTitle = eh.ParkingZone.Title
		}
		result = append(result, view)
	}
	return
}
