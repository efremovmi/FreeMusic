package metrics

import "github.com/prometheus/client_golang/prometheus"

var UploadCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "upload_request_count",
		Help: "No of request handled by upload handler",
	},
)

func RegisterMetrics() {
	prometheus.MustRegister(UploadCounter)
}
