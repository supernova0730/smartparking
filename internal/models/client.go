package models

import (
	"smartparking/internal/dtos"
	"smartparking/internal/views"
	"smartparking/pkg/tools"
	"time"
)

type ListClient []Client
type Client struct {
	ID           int64     `gorm:"column:id"`
	FirstName    string    `gorm:"column:first_name"`
	LastName     string    `gorm:"column:last_name"`
	Email        string    `gorm:"column:email;unique"`
	Phone        string    `gorm:"column:phone;unique"`
	Password     string    `gorm:"column:password"`
	NumberOfCars int       `gorm:"column:number_of_cars"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}

func (Client) TableName() string {
	return "client"
}

func (c Client) GetFullName() string {
	return c.FirstName + " " + c.LastName
}

func (c Client) ToView() views.ClientView {
	return views.ClientView{
		ID:           tools.Int64ToString(c.ID),
		FirstName:    c.FirstName,
		LastName:     c.LastName,
		Email:        c.Email,
		Phone:        c.Phone,
		NumberOfCars: tools.IntToString(c.NumberOfCars),
		CreatedAt:    c.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    c.UpdatedAt.Format(time.RFC3339),
	}
}

func (list ListClient) ToView() (result []views.ClientView) {
	for _, client := range list {
		result = append(result, client.ToView())
	}
	return
}

func (c *Client) SetFromRegisterDTO(dto dtos.RegisterDTO) {
	c.FirstName = dto.FirstName
	c.LastName = dto.LastName
	c.Email = dto.Email
	c.Phone = dto.Phone
	c.Password = dto.Password
}

func (c *Client) SetFromLoginDTO(dto dtos.LoginDTO) {
	c.Email = dto.Email
	c.Password = dto.Password
}

func (c *Client) SetFromUpdateDTO(dto dtos.ClientUpdateDTO) {
	c.FirstName = dto.FirstName
	c.LastName = dto.LastName
	c.Phone = dto.Phone
}
