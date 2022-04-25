package models

import (
	"smartparking/internal/dtos"
	"smartparking/internal/views"
	"smartparking/pkg/tools"
)

type ListCar []Car
type Car struct {
	ID       int64  `gorm:"column:id"`
	Title    string `gorm:"column:title"`
	Plate    string `gorm:"column:plate"`
	IsActive bool   `gorm:"column:is_active"`
	ClientID int64  `gorm:"column:client_id"`

	Client *Client
}

func (Car) TableName() string {
	return "car"
}

func (c Car) ToView() views.CarDetailView {
	view := views.CarDetailView{
		ID:       tools.Int64ToString(c.ID),
		Title:    c.Title,
		Plate:    c.Plate,
		IsActive: c.IsActive,
	}
	if c.Client != nil {
		view.Client = c.Client.ToView()
	}
	return view
}

func (list ListCar) ToView() (result []views.CarListView) {
	for _, car := range list {
		view := views.CarListView{
			ID:       tools.Int64ToString(car.ID),
			Title:    car.Title,
			Plate:    car.Plate,
			IsActive: car.IsActive,
			ClientID: tools.Int64ToString(car.ClientID),
		}
		if car.Client != nil {
			view.ClientFullName = car.Client.GetFullName()
		}
		result = append(result, view)
	}
	return
}

func (c *Car) SetFromCreateUpdateDTO(dto dtos.CarCreateUpdateDTO) {
	c.Title = dto.Title
	c.Plate = dto.Plate
	c.IsActive = dto.IsActive
	c.ClientID = dto.ClientID
}

func (c *Car) SetFromPersonalCarCreateDTO(dto dtos.PersonalCarCreateDTO) {
	c.Title = dto.Title
	c.Plate = dto.Plate
}
