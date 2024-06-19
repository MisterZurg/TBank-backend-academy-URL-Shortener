package main

import (
	"github.com/labstack/gommon/log"

	"github.com/MisterZurg/TBank-backend-academy-URL-Shortener/backend/config"
	"github.com/MisterZurg/TBank-backend-academy-URL-Shortener/backend/internal/handler"
	"github.com/MisterZurg/TBank-backend-academy-URL-Shortener/backend/internal/repository"
	"github.com/MisterZurg/TBank-backend-academy-URL-Shortener/backend/internal/service"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal("Fuck: Cannot get Config")
	}

	repo, err := repository.New(repository.Config{
		RedisDSN: cfg.GetRedisDSN(),
		CH: repository.ClickHouseConfig{
			ClickHouseDSN: cfg.GetClickHouseDSN(),
			DBName:        cfg.CHDBName,
			Username:      cfg.CHUser,
			Password:      cfg.CHPassword,
		},
	})
	if err != nil {
		log.Fatalf("Fuck: Cannot get repo %e", err)
	}

	svc := service.New(repo)

	e := handler.RegisterHandlers(svc)

	e.Logger.Fatal(e.Start(cfg.GetAppAddress()))
}
