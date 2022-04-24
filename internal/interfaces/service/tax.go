package service

import "smartparking/internal/models"

type Tax interface {
	GetByID(id int64) (result models.Tax, err error)
	GetAll() (result []models.Tax, err error)
	GetAllOrderByPrice() (result []models.Tax, err error)
	Create(model models.Tax) (result models.Tax, err error)
	Update(id int64, model models.Tax) (result models.Tax, err error)
	Delete(id int64) (err error)
}
