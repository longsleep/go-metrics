package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

// PrometheusHTTPRequestCount is a prometheus.CounterVec with name
// http_request_count and method, type and endpoint labels.
var PrometheusHTTPRequestCount = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_request_count",
		Help: "The number of HTTP requests.",
	},
	[]string{"method", "type", "endpoint"},
)

// PrometheusHTTPRequestLatency is a new prometheus.SummaryVec with name
// http_request_latency and method, type and endpoint labels.
var PrometheusHTTPRequestLatency = prometheus.NewSummaryVec(
	prometheus.SummaryOpts{
		Name: "http_request_latency",
		Help: "The latency of HTTP requests.",
	},
	[]string{"method", "type", "endpoint"},
)

// PrometheusHTTPResponseCount is a prometheus.CounterVec wuth name
// http_response_count and method, type, endpoint and code labels.
var PrometheusHTTPResponseCount = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_response_count",
		Help: "The number of HTTP responses.",
	},
	[]string{"method", "type", "endpoint", "code"},
)

// PrometheusWebsocketMessageCount is a prometheus.CounterVec with name
// websocket_message_count and type and direction labels.
var PrometheusWebsocketMessageCount = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "websocket_message_count",
		Help: "The number of Websocket messages.",
	},
	[]string{"type", "direction"},
)

// PrometheusWebsocketMessageSize is a prometheus.SummaryVec with name
// websocket_message_size and type and direction labels.
var PrometheusWebsocketMessageSize = prometheus.NewSummaryVec(
	prometheus.SummaryOpts{
		Name: "websocket_message_size",
		Help: "The size of websocket messages.",
	},
	[]string{"type", "direction"},
)
