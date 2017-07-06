package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

// MustRegister registers the provided prometheus.Collectors and the pre-defined
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
