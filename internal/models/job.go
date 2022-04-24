package models

type ListJob []Job
type Job struct {
	ID        int64  `gorm:"column:id"`
	Code      string `gorm:"column:code;unique"`
	Schedule  string `gorm:"column:schedule"`
	IsActive  bool   `gorm:"column:is_active;default:false"`
	IsRunning bool   `gorm:"column:is_running;default:false"`
}

func (Job) TableName() string {
	return "jobs"
}
