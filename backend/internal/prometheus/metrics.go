package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	TotalOpsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "tbank_processed_ops_total",
		Help: "The total number of processed events",
	})

	RedirectNum = promauto.NewCounter(prometheus.CounterOpts{
		Name: "tbank_url_shortener_usage_redirect",
		Help: "The number of redirects",
	})

	CreateURLS = promauto.NewCounter(prometheus.CounterOpts{
		Name: "tbank_url_shortener_usage_create",
		Help: "The number of short urls creation",
	})

	GotURLFromCache = promauto.NewCounter(prometheus.CounterOpts{
		Name: "tbank_url_shortener_cache_usage",
		Help: "The number of cache usage",
	})

	GotURLFromDB = promauto.NewCounter(prometheus.CounterOpts{
		Name: "tbank_url_shortener_db_usage",
		Help: "The number of db usage",
	})

	TotalErrors = promauto.NewCounter(prometheus.CounterOpts{
		Name: "tbank_url_shortener_errors_total",
		Help: "The number of errors",
	})
)
