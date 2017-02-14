package metrics

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// key is an unexported type for keys defined in this package.
type key int

// elapsedKey is the key for elapsedRecord which holds start end elapsed time
// in Contexts.
const elapsedKey key = 0

// MustRegister regusters the provides prometheus.Collectors and the pre defined
// prometheus.Collectors with the prometheus.DefaultRegisterer and panics of any
// error occurs.
func MustRegister(collectors ...prometheus.Collector) {
	prometheus.MustRegister(PrometheusHTTPRequestCount)
	prometheus.MustRegister(PrometheusHTTPRequestLatency)
	prometheus.MustRegister(PrometheusHTTPResponseCount)
	prometheus.MustRegister(PrometheusWebsocketMessageCount)
	prometheus.MustRegister(PrometheusWebsocketMessageSize)

	for _, collector := range collectors {
		prometheus.MustRegister(collector)
	}
}

type elapsedRecord struct {
	start   time.Time
	elapsed time.Duration
	cancel  context.CancelFunc
}

// NewContext returns a new Context that carries the start time and calls the
// provided stopped function then the created Context is cancelled. The stopped
// function can be nil in which case nothing is called.
func NewContext(parent context.Context, stopped func(elapsed time.Duration)) context.Context {
	ctx, cancel := context.WithCancel(parent)
	recordPtr := &elapsedRecord{
		start:  time.Now(),
		cancel: cancel,
	}
	ctx = context.WithValue(ctx, elapsedKey, recordPtr)
	go func() {
		<-ctx.Done()
		recordPtr.elapsed = time.Now().Sub(recordPtr.start)
		if stopped != nil {
			stopped(recordPtr.elapsed)
		}
	}()

	return ctx
}

// StartFromContext returns the start time from the provided Context.
func StartFromContext(ctx context.Context) time.Time {
	return ctx.Value(elapsedKey).(*elapsedRecord).start
}

// ElapsedFromContext returns the elapsed time from the provided Context.
func ElapsedFromContext(ctx context.Context) time.Duration {
	return ctx.Value(elapsedKey).(*elapsedRecord).elapsed
}

// CancelContext cancels the provided Context if it carries start time.
func CancelContext(ctx context.Context) {
	recordPtr := ctx.Value(elapsedKey).(*elapsedRecord)
	if recordPtr != nil {
		recordPtr.cancel()
	}
}
