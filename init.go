package metrics

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type key int

const (
	elapsedKey key = 0
)

func MustRegister(collectors ...prometheus.Collector) {
	appName := os.Args[0]
	log.Printf("Registering metrics for %s\n", appName)

	prometheus.MustRegister(PrometheusHTTPRequestCount)
	prometheus.MustRegister(PrometheusHTTPRequestLatency)
	prometheus.MustRegister(PrometheusHTTPResponseCount)
	prometheus.MustRegister(PrometheusWebsocketMessageCount)
	prometheus.MustRegister(PrometheusWebsocketMessageSize)

	for _, collector := range collectors {
		prometheus.MustRegister(collector)
	}
}

type StoppedFunc func(elapsed time.Duration)

type elapsedRecord struct {
	start   time.Time
	elapsed time.Duration
	cancel  context.CancelFunc
}

func NewContext(parent context.Context, stopped StoppedFunc) context.Context {
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

func StartFromContext(ctx context.Context) time.Time {
	return ctx.Value(elapsedKey).(*elapsedRecord).start
}

func ElapsedFromContext(ctx context.Context) time.Duration {
	return ctx.Value(elapsedKey).(*elapsedRecord).elapsed
}

func CancelContext(ctx context.Context) {
	recordPtr := ctx.Value(elapsedKey).(*elapsedRecord)
	if recordPtr != nil {
		recordPtr.cancel()
	}
}
