package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"smartparking/internal/models"
	"smartparking/pkg/logger"
)

type jobRepository struct {
	db *gorm.DB
}

func NewJobRepository(db *gorm.DB) *jobRepository {
	return &jobRepository{db: db}
}

func (repo *jobRepository) GetByCode(code string) (result models.Job, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("jobRepository.GetByCode failed", zap.Error(err), zap.String("code", code))
		}
	}()

	err = repo.db.Model(&models.Job{}).Where("code = ?", code).First(&result).Error
	return
}

func (repo *jobRepository) GetAll() (result models.ListJob, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("jobRepository.GetAll failed", zap.Error(err))
		}
	}()

	err = repo.db.Model(&models.Job{}).Find(&result).Error
	return
}

func (repo *jobRepository) GetAllActive() (result models.ListJob, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("jobRepository.GetAllActive failed", zap.Error(err))
		}
	}()

	err = repo.db.Model(&models.Job{}).Where("is_active = ?", true).Find(&result).Error
	return
}

func (repo *jobRepository) Enable(id int64) (err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("jobRepository.Enable failed", zap.Error(err), zap.Int64("id", id))
		}
	}()

	err = repo.db.Model(&models.Job{}).Where("id = ?", id).Update("is_active", true).Error
	return
}

func (repo *jobRepository) Disable(id int64) (err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("jobRepository.Disable failed", zap.Error(err), zap.Int64("id", id))
		}
	}()

	err = repo.db.Model(&models.Job{}).Where("id = ?", id).Update("is_active", false).Error
	return
}

func (repo *jobRepository) SetIsRunning(id int64, isRunning bool) (err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("jobRepository.SetIsRunning failed", zap.Error(err), zap.Int64("id", id), zap.Bool("isRunning", isRunning))
		}
	}()

	err = repo.db.Model(&models.Job{}).Where("id = ?", id).Update("is_running", isRunning).Error
	return
}
