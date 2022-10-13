package cache

import (
	"sync"
	"time"
)

const timeout = time.Millisecond * 2

type Cache struct {
	c map[string]string
	m *sync.RWMutex
}

func InitCache() *Cache {
	return &Cache{
		c: make(map[string]string),
		m: new(sync.RWMutex),
	}
}

func (c *Cache) Add(key string, value string) {
	c.m.Lock()
	time.Sleep(timeout)
	c.c[key] = value
	c.m.Unlock()
}

func (c *Cache) Get(key string) (value string, ok bool) {
	c.m.RLock()

	time.Sleep(timeout)

	value, ok = c.c[key]
	c.m.RUnlock()
	return
}

func (c *Cache) Delete(key string) {
	c.m.Lock()
	time.Sleep(timeout)
	delete(c.c, key)
	c.m.Unlock()
}
