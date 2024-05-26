package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// TODO: rewrite
const prefix = "router_"

var RequestsTotal = promauto.NewCounter(
	prometheus.CounterOpts{
		Name: prefix + "requests_processed_total",
		Help: "The total number of processed requests",
	},
)

var StatusCodes = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: prefix + "http_status_codes",
		Help: "HTTP response status codes",
	},
	[]string{"status_code"},
)
