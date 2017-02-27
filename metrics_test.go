package metrics

import (
	"testing"

	"context"
	"time"
)

func TestMetrics_NewContext(t *testing.T) {
	NewContext(context.Background(), nil)
}

func TestMetrics_ElapsedFromContext(t *testing.T) {
	ctx := NewContext(context.Background(), nil)
	time.Sleep(10 * time.Millisecond) // Fake some work.

	elapsed := ElapsedFromContext(ctx)
	if elapsed < time.Duration(10)*time.Millisecond {
		t.Errorf("Expected elapsed time to be greater than 10ms, but was %s", elapsed)
	}
}

func TestMetrics_ElapsedAfterContextCancel(t *testing.T) {
	elapsedCh := make(chan time.Duration, 1)
	ctx := NewContext(context.Background(), func(elapsedFromCb time.Duration) {
		elapsedCh <- elapsedFromCb
	})
	time.Sleep(10 * time.Millisecond) // Fake some work.

	CancelContext(ctx)
	elapsed1 := ElapsedFromContext(ctx)
	if elapsed1 < time.Duration(10)*time.Millisecond {
		t.Errorf("Expected elapsed time to be greater than 10ms, but was %s", elapsed1)
	}
	elapsed2 := <-elapsedCh
	if elapsed1 != elapsed2 {
		t.Errorf("Expected elapsed time from callback to the the same than from value")
	}
}
