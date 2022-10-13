package cache_v3

import (
	"antchel/storage"
	"sync"
)

type simpleCache struct {
	storage map[string]string
	m       sync.Mutex
}

func NewStorage() storage.Cache {
	return &simpleCache{storage: make(map[string]string),
		m: sync.Mutex{}}
}

func NewStorageSized(size int) storage.Cache {
	return &simpleCache{storage: make(map[string]string, size),
		m: sync.Mutex{}}

}

func (s *simpleCache) Set(key, value string) error {
	s.m.Lock()
	s.storage[key] = value
	s.m.Unlock()
	return nil
}

func (s *simpleCache) Get(key string) (string, error) {
	s.m.Lock()
	v, ok := s.storage[key]
	s.m.Unlock()
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
