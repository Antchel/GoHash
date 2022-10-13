package cache_v2

import (
	"antchel/storage"
	"sync"
)

type simpleCache struct {
	storage map[string]string
	m       sync.RWMutex
}

func NewStorage() storage.Cache {
	return &simpleCache{storage: make(map[string]string),
		m: sync.RWMutex{}}
}

func NewStorageSized(size int) storage.Cache {
	return &simpleCache{storage: make(map[string]string, size),
		m: sync.RWMutex{}}

}

func (s *simpleCache) Set(key, value string) error {
	s.m.Lock()
	s.storage[key] = value
	s.m.Unlock()
	return nil
}

func (s *simpleCache) Get(key string) (string, error) {
	s.m.RLock()
	v, ok := s.storage[key]
	s.m.RUnlock()
	if !ok {
		return "", storage.ErrNotFound
	}
	return v, nil
}

func (s *simpleCache) Delete(key string) error {
	s.m.Lock()
	delete(s.storage, key)
	s.m.Unlock()
	return nil
}
