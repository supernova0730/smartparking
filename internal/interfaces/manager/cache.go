package manager

import "smartparking/internal/interfaces/cache"

type Cache interface {
	BaseRepository() cache.Base
	Session() cache.Session
	EntryHistory() cache.EntryHistory
}
