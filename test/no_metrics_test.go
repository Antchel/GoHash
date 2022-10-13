package test

import (
	"antchel/cache_v2"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Cache(t *testing.T) {
	t.Parallel()

	testCache := cache_v2.NewStorage()

	t.Run("correct cache", func(t *testing.T) {
		t.Parallel()

		key := "someKey"
		value := "someValue"

		err := testCache.Set(key, value)
		assert.NoError(t, err)
		storedValue, err := testCache.Get(key)
		assert.NoError(t, err)
		assert.Equal(t, value, storedValue)
	})

	t.Run("correct update", func(t *testing.T) {
		t.Parallel()

		key := "someKey"
		value := "someValue"

		err := testCache.Set(key, value)
		assert.NoError(t, err)
		storedValue, err := testCache.Get(key)
		assert.NoError(t, err)
		assert.Equal(t, value, storedValue)

		newValue := "someValue2"
		err = testCache.Set(key, newValue)
		assert.NoError(t, err)
		newStoredValue, err := testCache.Get(key)
		assert.NoError(t, err)
		assert.Equal(t, newValue, newStoredValue)

	})

	t.Run("No data race", func(t *testing.T) {
		t.Parallel()

		parallel_factor := 100_000
		emulateLoad(t, testCache, parallel_factor)
	})

}
