// Package metrics provides init for Prometheus metrics
package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Count of all HTTP requests",
	}, []string{"code", "method"})
)

// Create provides initialization of Prometheus
func Create() {
	r := prometheus.NewRegistry()
	r.MustRegister(httpRequestsTotal)
}
