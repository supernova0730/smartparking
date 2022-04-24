package cache

import "time"

type Base interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
	SetWithTTL(key string, value []byte, ttl time.Duration) error
	Delete(key string) error
}
