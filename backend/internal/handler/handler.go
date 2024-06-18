package handler

import (
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"net/http"

	"github.com/MisterZurg/TBank_backend_academy_URL_Shortener/backend/internal/service"
)

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
//type EchoRouter interface {
//	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
//	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
//	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
//	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
//	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
//	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
//	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
//	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
//	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
//}

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

	e.GET("/sss", func(c echo.Context) error {
		opsProcessed.Inc()
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/short-it", svc.ShortenURL)
	e.GET("/short-it/:short_url", svc.GetURL)

	return e
}
