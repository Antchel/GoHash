package test

import (
	"antchel/storage"
	"errors"
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func emulateLoad(t *testing.T, c storage.Cache, paralelFactor int) {
	wg := &sync.WaitGroup{}

	for i := 0; i < paralelFactor; i++ {
		key := fmt.Sprintf("%d-key", i)
		value := fmt.Sprintf("%d-value", i)
		wg.Add(1)
		go func(k string) {
			err := c.Set(k, value)
			assert.NoError(t, err)
			defer wg.Done()
		}(key)

		wg.Add(1)
		go func(k, v string) {
			val, err := c.Get(k)
			if !errors.Is(err, storage.ErrNotFound) {
				assert.Equal(t, v, val)
			}
			defer wg.Done()
		}(key, value)

		wg.Add(1)
		go func(k string) {
			err := c.Delete(k)
			assert.NoError(t, err)
			defer wg.Done()
		}(key)
	}
}
