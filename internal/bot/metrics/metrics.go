package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	MessagesTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "bot_messages_total",
		Help: "Total number of handled bot updates/messages.",
	}, []string{"handler", "update_type", "chat_type"})

	HandlerErrors = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "bot_handler_errors_total",
		Help: "Total number of errors in handlers.",
	}, []string{"handler"})

	HandlerDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "bot_handler_duration_seconds",
		Help:    "Handler duration in seconds.",
		Buckets: prometheus.DefBuckets,
	}, []string{"handler"})

	InFlight = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "bot_handlers_in_flight",
		Help: "Number of handlers currently executing.",
	})
)

func Serve(listenAddr string) error {
	http.Handle("/metrics", promhttp.Handler())
	return http.ListenAndServe(listenAddr, nil)
}
