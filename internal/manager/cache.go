package manager

import (
	"smartparking/config"
	"smartparking/internal/interfaces/cache"
	cache2 "smartparking/internal/repository/cache"
	"sync"
)

type cacheImpl struct {
	baseRepositoryInit         sync.Once
	baseRepository             cache.Base
	sessionRepositoryInit      sync.Once
	sessionRepository          cache.Session
	entryHistoryRepositoryInit sync.Once
	entryHistoryRepository     cache.EntryHistory
}

func (cm *cacheImpl) BaseRepository() cache.Base {
	cm.baseRepositoryInit.Do(func() {
		if cm.baseRepository == nil {
			cm.baseRepository = cache2.NewBaseRepository(config.CacheConn)
		}
	})
	return cm.baseRepository
}

func (cm *cacheImpl) Session() cache.Session {
	cm.sessionRepositoryInit.Do(func() {
		if cm.sessionRepository == nil {
			cm.sessionRepository = cache2.NewSessionRepository(config.CacheConn)
		}
	})
	return cm.sessionRepository
}

func (cm *cacheImpl) EntryHistory() cache.EntryHistory {
	cm.entryHistoryRepositoryInit.Do(func() {
		if cm.entryHistoryRepository == nil {
			cm.entryHistoryRepository = cache2.NewEntryHistoryRepository(config.CacheConn)
		}
	})
	return cm.entryHistoryRepository
}
