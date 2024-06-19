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

type Config struct {
	RedisDSN string
	CH       ClickHouseConfig
}

type ClickHouseConfig struct {
	ClickHouseDSN string
	DBName        string
	Username      string
	Password      string
}

func New(rcfg Config) (*Repository, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     rcfg.RedisDSN,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx := context.Background()
	if errPingCache := rdb.Ping(ctx).Err(); errPingCache != nil { //nolint:staticcheck
		log.Printf("cannot ping cache, got err %v", errPingCache)
	}

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

	if errPingDB := ch.Ping(ctx); errPingDB != nil { //nolint:staticcheck
		log.Printf("cannot ping db, got err %v", errPingDB)
	}

	if rdb.Ping(ctx).Err() != nil && ch.Ping(ctx) != nil { //nolint:staticcheck
		log.Fatalf("cannot connect to both cache and storage")
	}

	return &Repository{
		URLRepository: &URLRepository{
			rdb,
			ch,
		},
	}, nil
}
