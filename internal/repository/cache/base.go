package cache

import (
	"github.com/bradfitz/gomemcache/memcache"
	"time"
)

type BaseRepository struct {
	db *memcache.Client
}

func NewBaseRepository(db *memcache.Client) *BaseRepository {
	return &BaseRepository{db: db}
}

func (repo *BaseRepository) Get(key string) ([]byte, error) {
	item, err := repo.db.Get(key)
	if err != nil {
		return nil, err
	}
	return item.Value, nil
}

func (repo *BaseRepository) Set(key string, value []byte) error {
	return repo.db.Set(&memcache.Item{Key: key, Value: value})
}

func (repo *BaseRepository) SetWithTTL(key string, value []byte, ttl time.Duration) error {
	return repo.db.Set(&memcache.Item{Key: key, Value: value, Expiration: int32(ttl.Seconds())})
}

func (repo *BaseRepository) Delete(key string) error {
	return repo.db.Delete(key)
}
