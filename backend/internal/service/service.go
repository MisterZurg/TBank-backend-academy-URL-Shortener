package service

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	PostURL(url string) (string, error)
	GetURL(shortURL string) (string, error)
}

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

func (s *Service) ShortenURL(c echo.Context) error {
	params := new(PostURLParams)
	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	if params.LongURL == "" {
		return c.JSON(http.StatusBadRequest, nil)
	}

	short, err := s.repo.PostURL(params.LongURL)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	return c.String(http.StatusOK, short)
}

func (s *Service) GetURL(c echo.Context) error {
	shorten := c.Param("short_url")
	// log.Printf("svc got param short_url :> %s", shorten)
	if shorten == "" {
		return c.JSON(http.StatusBadRequest, nil)
	}

	redirectURL, err := s.repo.GetURL(shorten)
	if err != nil {
		return c.JSON(http.StatusNotFound, nil)
	}
	// c.Response().Header().Set("HX-Redirect", redirectURL)
	return c.Redirect(http.StatusFound, redirectURL)
}
