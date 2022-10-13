package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkAtomicCounter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = AtomicCounter()
	}
}

func BenchmarkMutexCounter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = MutexCounter()
	}
}

func Test_SomeTest(t *testing.T) {
	t.Parallel()

	t.Run("sum is calculated right", func(t *testing.T) {
		t.Parallel()

		res := sum(1, 2)

		assert.Equal(t, 3, res)
	})

	t.Run("error for broken string", func(t *testing.T) {
		t.Parallel()

		res, err := CalcSumFromString("text", "3")

		assert.NotNil(t, err)
		t.Log(err)

		assert.Equal(t, 0, res)
	})
}
