package main

import (
	"github.com/MisterZurg/TBank-backend-academy-URL-Shortener/backend/config"
	"github.com/MisterZurg/TBank-backend-academy-URL-Shortener/backend/internal/handler"
	"github.com/MisterZurg/TBank-backend-academy-URL-Shortener/backend/internal/repository"
	"github.com/MisterZurg/TBank-backend-academy-URL-Shortener/backend/internal/service"
)

func main() {
	cfg, _ := config.New()

	repo, _ := repository.New(repository.Config{
		RedisDSN: cfg.GetRedisDSN(),
		CH: repository.ClickHouseConfig{
			ClickHouseDSN: cfg.GetClickHouseDSN(),
			DBName:        cfg.CHDBName,
			Username:      cfg.CHUser,
			Password:      cfg.CHPassword,
		},
	})

	svc := service.New(repo)

	e := handler.RegisterHandlers(svc)

	e.Logger.Fatal(e.Start(cfg.GetAppAddress()))
}
