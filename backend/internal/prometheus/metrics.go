package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	RedirectNum = promauto.NewCounter(prometheus.CounterOpts{
		Name: "tbank_url_shortener_usage_redirect",
		Help: "The total number of processed events",
	})

	CreateURLS = promauto.NewCounter(prometheus.CounterOpts{
		Name: "tbank_url_shortener_usage_create",
		Help: "The total number of processed events",
	})

	GotURLFromCache = promauto.NewCounter(prometheus.CounterOpts{
		Name: "tbank_url_shortener_cache_usage",
		Help: "The total number of processed events",
	})

	GotURLFromDB = promauto.NewCounter(prometheus.CounterOpts{
		Name: "tbank_url_shortener_db_usage",
		Help: "The total number of processed events",
	})
)
