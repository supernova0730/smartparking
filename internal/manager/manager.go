package manager

import (
	"smartparking/internal/interfaces/manager"
	"sync"
)

var (
	MManager manager.Manager = &managerImpl{}
)

type managerImpl struct {
	repositoryInit sync.Once
	repository     manager.Repository
	cacheInit      sync.Once
	cache          manager.Cache
	serviceInit    sync.Once
	service        manager.Service
	processorInit  sync.Once
	processor      manager.Processor
	controllerInit sync.Once
	controller     manager.Controller
}

func (m *managerImpl) Repository() manager.Repository {
	m.repositoryInit.Do(func() {
		if m.repository == nil {
			m.repository = &repositoryImpl{}
		}
	})
	return m.repository
}

func (m *managerImpl) Cache() manager.Cache {
	m.cacheInit.Do(func() {
		if m.cache == nil {
			m.cache = &cacheImpl{}
		}
	})
	return m.cache
}

func (m *managerImpl) Service() manager.Service {
	m.serviceInit.Do(func() {
		if m.service == nil {
			m.service = &serviceImpl{}
		}
	})
	return m.service
}

func (m *managerImpl) Processor() manager.Processor {
	m.processorInit.Do(func() {
		if m.processor == nil {
			m.processor = &processorImpl{}
		}
	})
	return m.processor
}

func (m *managerImpl) Controller() manager.Controller {
	m.controllerInit.Do(func() {
		if m.controller == nil {
			m.controller = &controllerImpl{}
		}
	})
	return m.controller
}
