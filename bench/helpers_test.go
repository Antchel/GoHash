package test

import (
	"antchel/storage"
	"errors"
	"fmt"
	"sync"
)

func emulateLoad(c storage.Cache, paralelFactor int) {
	wg := &sync.WaitGroup{}

	for i := 0; i < paralelFactor; i++ {
		key := fmt.Sprintf("%d-key", i)
		value := fmt.Sprintf("%d-value", i)
		wg.Add(1)
		go func(k string) {
			err := c.Set(k, value)
			defer wg.Done()
			if err != nil {
				panic(err)
			}
		}(key)

		wg.Add(1)
		go func(k, v string) {
			_, err := c.Get(k)
			if err != nil && !errors.Is(err, storage.ErrNotFound) {
				panic(err)
			}
			defer wg.Done()
		}(key, value)

		wg.Add(1)
		go func(k string) {
			err := c.Delete(k)
			if err != nil {
				panic(err)
			}
			defer wg.Done()
		}(key)
	}
}

func emulateBalancedLoad(c storage.Cache, paralelFactor int) {
	wg := &sync.WaitGroup{}

	for i := 0; i < paralelFactor/10; i++ {
		key := fmt.Sprintf("%d-key", i)
		value := fmt.Sprintf("%d-value", i)
		wg.Add(1)
		go func(k string) {
			err := c.Set(k, value)
			defer wg.Done()
			if err != nil {
				panic(err)
			}
		}(key)

		wg.Add(1)
		go func(k string) {
			err := c.Delete(k)
			if err != nil {
				panic(err)
			}
			defer wg.Done()
		}(key)
	}

	for i := 0; i < paralelFactor; i++ {
		key := fmt.Sprintf("%d-key", i)
		value := fmt.Sprintf("%d-value", i)

		wg.Add(1)
		go func(k, v string) {
			_, err := c.Get(k)
			if err != nil && !errors.Is(err, storage.ErrNotFound) {
				panic(err)
			}
			defer wg.Done()
		}(key, value)

	}
}
