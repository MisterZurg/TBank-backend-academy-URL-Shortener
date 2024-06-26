package service

import (
	"github.com/MisterZurg/TBank-backend-academy-URL-Shortener/backend/internal/prometheus"
	"github.com/MisterZurg/TBank-backend-academy-URL-Shortener/backend/urlerrors"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Repository — ...
type Repository interface {
	PostURL(url string) (string, error)
	GetURL(shortURL string) (string, error)
}

// Service — ...
type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

type GetURLResponse struct {
	LongURL string `json:"long_url" param:"long_url" query:"long_url"`
}

type PostURLParams struct {
	LongURL string `json:"long_url"`
}

type PostURLResponse struct {
	ShortURL string `json:"short_url,omitempty"`
}

// ShortenURL — creates short url from given
//
// @Summary      Creates short url from given
// @Description  Creates short url from given
// @Accept       json
// @Produce      json
// @Success      200  {string}  string
// @Router       /short-it [post]
func (s *Service) ShortenURL(c echo.Context) error {
	prometheus.TotalOpsProcessed.Inc()

	params := new(PostURLParams)
	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, "----")
	}

	if params.LongURL == "" {
		prometheus.TotalErrors.Inc()
		return c.JSON(http.StatusBadRequest, urlerrors.ErrEmptyURL)
	}

	short, err := s.repo.PostURL(params.LongURL)
	if err != nil {
		prometheus.TotalErrors.Inc()
		return c.JSON(http.StatusBadRequest, urlerrors.ErrInternalAppError)
	}

	return c.String(http.StatusOK, short)
}

// GetURL — redirects user to short url
//
// @Summary      Redirects user
// @Description  Redirects user
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        short_url   path    string  true  "short_url"
// @Success      302  {string}  string
// @Failure      400  {object}  urlerrors.Error
// @Failure      404  {object}  urlerrors.Error
// @Router       /short-it/{short_url} [get]
func (s *Service) GetURL(c echo.Context) error {
	prometheus.TotalOpsProcessed.Inc()

	shorten := c.Param("short_url")
	if shorten == "" {
		prometheus.TotalErrors.Inc()
		return c.JSON(http.StatusBadRequest, urlerrors.ErrEmptyURL)
	}

	redirectURL, err := s.repo.GetURL(shorten)
	if err != nil {
		prometheus.TotalErrors.Inc()
		return c.JSON(http.StatusNotFound, urlerrors.ErrCannotFindURL)
	}
	// c.Response().Header().Set("HX-Redirect", redirectURL)
	prometheus.RedirectNum.Inc()
	return c.Redirect(http.StatusFound, redirectURL)
}
