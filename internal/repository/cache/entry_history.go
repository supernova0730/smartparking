package cache

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"go.uber.org/zap"
	"smartparking/config"
	"smartparking/internal/models"
	"smartparking/pkg/logger"
	"time"
)

type entryHistoryRepository struct {
	*BaseRepository
	ttl time.Duration
}

func NewEntryHistoryRepository(db *memcache.Client) *entryHistoryRepository {
	return &entryHistoryRepository{
		BaseRepository: NewBaseRepository(db),
		ttl:            config.GlobalConfig.Cache.TTL,
	}
}

func (repo *entryHistoryRepository) Get(id int64) (result models.EntryHistory, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("entryHistoryRepository.Get failed", zap.Error(err), zap.Int64("id", id))
		}
	}()

	data, err := repo.BaseRepository.Get(repo.key(id))
	if err != nil {
		return
	}

	err = gob.NewDecoder(bytes.NewBuffer(data)).Decode(&result)
	return
}

func (repo *entryHistoryRepository) Set(id int64, value models.EntryHistory) (err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("entryHistoryRepository.Set failed", zap.Error(err), zap.Int64("id", id), zap.Any("value", value))
		}
	}()

	buff := bytes.Buffer{}
	err = gob.NewEncoder(&buff).Encode(value)
	if err != nil {
		return
	}

	return repo.BaseRepository.SetWithTTL(repo.key(id), buff.Bytes(), repo.ttl)
}

func (repo *entryHistoryRepository) Delete(id int64) (err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("entryHistoryRepository.Delete failed", zap.Error(err), zap.Int64("id", id))
		}
	}()

	return repo.BaseRepository.Delete(repo.key(id))
}

func (repo *entryHistoryRepository) key(id int64) string {
	return fmt.Sprintf("entry_history_%d", id)
}
