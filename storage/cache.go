package storage

import "errors"

var ErrNotFound = errors.New("value not found")

type Cache interface {
	Get(key string) (string, error)
	Set(key string, value string) error
	Delete(key string) error
}

type CacheWithMetrics interface {
	Cache
	Metrics
}
