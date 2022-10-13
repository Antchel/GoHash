package cachev4

import (
	"antchel/storage"
	"sync"
)

type impl struct {
	storage map[string]string
	m       sync.Mutex
	total   int64
}

func New() storage.CacheWithMetrics {
	return &impl{
		storage: make(map[string]string),
		m:       sync.Mutex{},
		total:   0,
	}
}

func (s *impl) Set(key, value string) error {
	s.m.Lock()
	s.storage[key] = value
	s.m.Unlock()
	return nil
}

func (s *impl) Get(key string) (string, error) {
	s.m.Lock()
	v, ok := s.storage[key]
	s.m.Unlock()
	if !ok {
		return "", storage.ErrNotFound
	}
	return v, nil
}

func (s *impl) Delete(key string) error {
	s.m.Lock()
	delete(s.storage, key)
	s.m.Unlock()
	return nil
}

func (s *impl) TotalAmount() int64 {
	return s.total
}
