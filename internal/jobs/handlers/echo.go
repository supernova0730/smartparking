package handlers

import (
	"smartparking/internal/interfaces/manager"
	"smartparking/internal/models"
	"smartparking/pkg/logger"
)

type Echo struct {
	m   manager.Manager
	job models.Job
}

func NewEcho(m manager.Manager, job models.Job) *Echo {
	return &Echo{
		m:   m,
		job: job,
	}
}

func (h *Echo) GetSchedule() string {
	return h.job.Schedule
}

func (h *Echo) Do() {
	_ = h.m.Repository().Job().SetIsRunning(h.job.ID, true)
	defer func() {
		_ = h.m.Repository().Job().SetIsRunning(h.job.ID, false)
	}()

	logger.Log.Info("ECHO")
}
