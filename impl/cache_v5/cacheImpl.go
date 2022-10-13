package cache_v5

import (
	"antchel/storage"
	"sync"
)

type simpleCache struct {
	storage sync.Map
}

func NewStorage() storage.Cache {
	return &simpleCache{}
}

func NewStorageSized(size int) storage.Cache {
	return &simpleCache{}
}

func (s *simpleCache) Set(key, value string) error {
	s.storage.Store(key, value)
	return nil
}

func (s *simpleCache) Get(key string) (string, error) {
	val, ok := s.storage.Load(key)
	if !ok {
		return "", storage.ErrNotFound
	}
	return val.(string), nil
}

func (s *simpleCache) Delete(key string) error {
	s.storage.Delete(key)
	return nil
}
