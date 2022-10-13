package cache_v1

import "antchel/storage"

type simpleCache struct {
	storage map[string]string
}

func NewStorage() storage.Cache {
	return &simpleCache{storage: make(map[string]string)}
}

func NewStorageSized(size int) storage.Cache {
	return &simpleCache{storage: make(map[string]string, size)}
}

func (s *simpleCache) Set(key, value string) error {
	s.storage[key] = value
	return nil
}

func (s *simpleCache) Get(key string) (string, error) {
	val, ok := s.storage[key]
	if !ok {
		return "", storage.ErrNotFound
	}
	return val, nil
}

func (s *simpleCache) Delete(key string) error {
	delete(s.storage, key)
	return nil
}
