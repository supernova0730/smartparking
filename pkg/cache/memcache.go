package cache

import (
	"github.com/bradfitz/gomemcache/memcache"
)

func Conn(host, port string) (*memcache.Client, error) {
	address := host + ":" + port
	mc := memcache.New(address)
	if err := mc.Ping(); err != nil {
		return nil, err
	}
	return mc, nil
}
