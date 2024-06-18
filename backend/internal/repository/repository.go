package repository

import (
	"context"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/labstack/gommon/log"
	"github.com/redis/go-redis/v9"
)

type Repository struct {
	*URLRepository
}

type RepositoryConfig struct {
	RedisDSN string
	CH       ClickHouseConfig
}

type ClickHouseConfig struct {
	ClickHouseDSN string
	DBName        string
	Username      string
	Password      string
}

func New(rcfg RepositoryConfig) (*Repository, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     rcfg.RedisDSN,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx := context.Background()
	if errPingCache := rdb.Ping(ctx); errPingCache != nil {
		log.Printf("cannot ping cache, got err %v", errPingCache)
	}

	// TODO: connect to DB here
	ch, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{rcfg.CH.ClickHouseDSN},
		Auth: clickhouse.Auth{
			Database: rcfg.CH.DBName,
			Username: rcfg.CH.Username,
			Password: rcfg.CH.Password,
		},
	})
	if err != nil {
		log.Printf("cannot ping cache, got err %v", err)
	}

	if errPingCache := rdb.Ping(ctx); errPingCache != nil {
		log.Printf("cannot ping cache, got err %v", errPingCache)
	}

	if rdb == nil && ch == nil {
		log.Fatalf("cannot connect to both cache and storage")
	}

	return &Repository{
		URLRepository: &URLRepository{
			rdb,
			ch,
		},
	}, nil
}
