package prometheus

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	// Technical metrics
	HttpRequestsTotal   *prometheus.CounterVec
	HttpRequestDuration *prometheus.HistogramVec

	// Business metrics
}

func NewMetrics(reg prometheus.Registerer) *Metrics {
	m := &Metrics{}

	m.HttpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"path", "method"},
	)
	reg.MustRegister(m.HttpRequestsTotal)

	m.HttpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Histogram of the duration of HTTP requests.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path", "method"},
	)
	reg.MustRegister(m.HttpRequestDuration)

	return m
}
