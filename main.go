package main

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

func MutexCounter() int {
	// Executable goroutinec number
	goRoutineCounter := 0
	wg := sync.WaitGroup{}
	m := sync.Mutex{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			m.Lock()
			goRoutineCounter++
			m.Unlock()
			time.Sleep(time.Microsecond)
			m.Lock()
			goRoutineCounter--
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	return goRoutineCounter
}

func AtomicCounter() int32 {
	goRoutineCounter := int32(0)
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&goRoutineCounter, 1)

			time.Sleep(time.Microsecond)

			atomic.AddInt32(&goRoutineCounter, -1)
			wg.Done()
		}()
	}
	wg.Wait()
	return goRoutineCounter
}

func main() {
	fmt.Println(50)
}

func sum(a, b int) int {
	return a + b
}

func CalcSumFromString(x, y string) (int, error) {
	a, err := strconv.Atoi(x)

	if err != nil {
		return 0, err
	}
	b, err := strconv.Atoi(y)
	if err != nil {
		return 0, err
	}
	return a + b, nil
}
