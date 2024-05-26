package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	RequestsTotal prometheus.Counter
	StatusCodes   *prometheus.CounterVec
)

func InitMetrics(prefix string) {
	RequestsTotal = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: prefix + "requests_processed_total",
			Help: "The total number of processed requests",
		},
	)

	StatusCodes = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: prefix + "http_status_codes",
			Help: "HTTP response status codes",
		},
		[]string{"status_code"},
	)
}
