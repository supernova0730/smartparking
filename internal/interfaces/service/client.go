package service

import "smartparking/internal/models"

type Client interface {
	GetByID(id int64) (result models.Client, err error)
	GetByPhone(phone string) (result models.Client, err error)
	GetByEmail(email string) (result models.Client, err error)
	GetAll() (result models.ListClient, err error)
	Create(model models.Client) (result models.Client, err error)
	Update(id int64, model models.Client) (result models.Client, err error)
	Delete(id int64) (err error)
}
