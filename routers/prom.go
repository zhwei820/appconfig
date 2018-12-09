package routers

import "github.com/prometheus/client_golang/prometheus"

var (
	RequestCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "http_request_total_code",
			Help: "total request code controller",
		},
	)
)

func init() {
	prometheus.MustRegister(RequestCounter)
}
