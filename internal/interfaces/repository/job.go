package repository

import "smartparking/internal/models"

type Job interface {
	GetByCode(code string) (result models.Job, err error)
	GetAll() (result []models.Job, err error)
	GetAllActive() (result []models.Job, err error)
	Enable(id int64) (err error)
	Disable(id int64) (err error)
	SetIsRunning(id int64, isRunning bool) (err error)
}
