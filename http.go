package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var PrometheusHTTPRequestCount = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_request_count",
		Help: "The number of HTTP requests.",
	},
	[]string{"method", "type", "endpoint"},
)

var PrometheusHTTPRequestLatency = prometheus.NewSummaryVec(
	prometheus.SummaryOpts{
		Name: "http_request_latency",
		Help: "The latency of HTTP requests.",
	},
	[]string{"method", "type", "endpoint"},
)

var PrometheusHTTPResponseCount = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_response_count",
		Help: "The number of HTTP responses.",
	},
	[]string{"method", "type", "endpoint", "code"},
)

var PrometheusWebsocketMessageCount = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "websocket_message_count",
		Help: "The number of Websocket messages.",
	},
	[]string{"type", "direction"},
)

var PrometheusWebsocketMessageSize = prometheus.NewSummaryVec(
	prometheus.SummaryOpts{
		Name: "websocket_message_size",
		Help: "The size of websocket messages.",
	},
	[]string{"type", "direction"},
)
