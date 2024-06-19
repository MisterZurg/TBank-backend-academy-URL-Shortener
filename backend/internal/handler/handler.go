package handler

import (
	"net/http"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/MisterZurg/TBank-backend-academy-URL-Shortener/backend/internal/service"
)

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
)

func RegisterHandlers(svc *service.Service) *echo.Echo {
	e := echo.New()

	e.Use(echoprometheus.NewMiddleware("myapp"))            // adds middleware to gather metrics
	e.GET("/short-it/metrics", echoprometheus.NewHandler()) // adds route to serve gathered metrics

	e.GET("/hello-world", func(c echo.Context) error {
		opsProcessed.Inc()
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/short-it", svc.ShortenURL)
	e.GET("/short-it/:short_url", svc.GetURL)

	return e
}
