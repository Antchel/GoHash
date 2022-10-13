package test

import (
	"antchel/impl/cache_v1"
	"antchel/impl/cache_v2"
	"antchel/impl/cache_v3"
	"antchel/impl/cache_v5"
	"testing"
)

const paralelFactor = 1_000_000

func BenchmarkNoMutex(b *testing.B) {
	b.Skip("panic in mutex")

	c := cache_v1.NewStorage()

	for i := 0; i < b.N; i++ {
		emulateLoad(c, paralelFactor)
	}
}

func BenchmarkRWMutexLoad(b *testing.B) {
	// b.Skip()

	c := cache_v2.NewStorage()

	for i := 0; i < b.N; i++ {
		emulateLoad(c, paralelFactor)
	}
}

func BenchmarkMutexLoad(b *testing.B) {
	// b.Skip()

	c := cache_v3.NewStorage()

	for i := 0; i < b.N; i++ {
		emulateLoad(c, paralelFactor)
	}
}

func BenchmarkSyncMapLoad(b *testing.B) {
	// b.Skip()

	c := cache_v5.NewStorage()

	for i := 0; i < b.N; i++ {
		emulateLoad(c, paralelFactor)
	}
}

func BenchmarkRWMutex_Balanced_Load(b *testing.B) {
	b.Skip()

	c := cache_v2.NewStorage()

	for i := 0; i < b.N; i++ {
		emulateBalancedLoad(c, paralelFactor)
	}
}

func BenchmarkMutex_Balanced_Load(b *testing.B) {
	b.Skip()

	c := cache_v3.NewStorage()

	for i := 0; i < b.N; i++ {
		emulateBalancedLoad(c, paralelFactor)
	}
}
