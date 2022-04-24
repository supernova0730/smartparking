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

type sessionRepository struct {
	*BaseRepository
	ttl time.Duration
}

func NewSessionRepository(db *memcache.Client) *sessionRepository {
	return &sessionRepository{
		BaseRepository: NewBaseRepository(db),
		ttl:            config.GlobalConfig.Cache.TTL,
	}
}

func (repo *sessionRepository) Get(id int64) (result models.Session, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("sessionCache.Get failed", zap.Error(err), zap.Int64("id", id))
		}
	}()

	data, err := repo.BaseRepository.Get(repo.key(id))
	if err != nil {
		return
	}

	err = gob.NewDecoder(bytes.NewBuffer(data)).Decode(&result)
	return
}

func (repo *sessionRepository) Set(id int64, value models.Session) (err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("sessionCache.Set failed", zap.Error(err), zap.Int64("id", id), zap.Any("value", value))
		}
	}()

	buff := bytes.Buffer{}
	err = gob.NewEncoder(&buff).Encode(value)
	if err != nil {
		return
	}

	return repo.BaseRepository.SetWithTTL(repo.key(id), buff.Bytes(), repo.ttl)
}

func (repo *sessionRepository) key(id int64) string {
	return fmt.Sprintf("session_%d", id)
}
