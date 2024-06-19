package handler

import (
	"net/http"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"

	"github.com/MisterZurg/TBank-backend-academy-URL-Shortener/backend/internal/prometheus"
	"github.com/MisterZurg/TBank-backend-academy-URL-Shortener/backend/internal/service"
)

// RegisterHandlers — ...
func RegisterHandlers(svc *service.Service) *echo.Echo {
	e := echo.New()

	e.Use(echoprometheus.NewMiddleware("myapp"))            // adds middleware to gather metrics
	e.GET("/short-it/metrics", echoprometheus.NewHandler()) // adds route to serve gathered metrics

	e.GET("/hello-world", func(c echo.Context) error {
		prometheus.TotalOpsProcessed.Inc()
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/short-it", svc.ShortenURL)
	e.GET("/short-it/:short_url", svc.GetURL)
	return e
}
