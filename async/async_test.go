package async

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAsyncSet(t *testing.T) {
	ac := InitAsyncCache()
	to := time.Millisecond

	ctxBase := context.Background()
	ctx, _ := context.WithTimeout(ctxBase, to)

	err := ac.Add(ctx, "key", "value")
	if err != ErrTimeout {
		t.Error("Expected timeout")
	}

	to = time.Millisecond * 3
	ctx, _ = context.WithTimeout(ctxBase, to*5)

	err = ac.Add(ctx, "key", "value")
	if err != nil {
		t.Errorf("Expected Set %v", err)
	}

}

func TestAsyncGet(t *testing.T) {
	ac := InitAsyncCache()
	to := time.Millisecond
	key := "key"
	value := "value"

	ctxBase := context.Background()
	ctx, _ := context.WithTimeout(ctxBase, to)

	_ = ac.Add(ctx, key, value)

	val, err := ac.Get(ctx, key)
	if err != ErrTimeout {
		t.Error("Expected timeout")
	} else {
		assert.Equal(t, val, "")
	}
	to = time.Millisecond * 3
	ctx, _ = context.WithTimeout(ctxBase, to*50)

	val, err = ac.Get(ctx, key)
	if err != nil {
		t.Errorf("Expected Set %v", err)
	}
	assert.Equal(t, val, value)

}
