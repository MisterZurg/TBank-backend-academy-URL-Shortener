package repository

import (
	"context"
	"github.com/MisterZurg/TBank-backend-academy-URL-Shortener/backend/urlerrors"
	"log"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/lithammer/shortuuid"
	"github.com/redis/go-redis/v9"

	"github.com/MisterZurg/TBank-backend-academy-URL-Shortener/backend/internal/prometheus"
)

// URLRepository — entity that represents Redis and ClickHouse
type URLRepository struct {
	// *sqlx.DB
	cache *redis.Client
	db    driver.Conn
}

// GetURL — checks if shortURL exists in Redis and ClickHouse, if so returns full url
func (ur *URLRepository) GetURL(shortURL string) (string, error) {
	ctx := context.Background()

	longURL, err := ur.getFromCache(ctx, shortURL)
	if err == redis.Nil {
		log.Printf("NO IN CACHE")

		longURL, err = ur.getFromDB(ctx, shortURL)
		if err != nil {
			return "", err
		}
	}

	prometheus.GotURLFromDB.Inc()
	return longURL, nil
}

// PostURL — ...
func (ur *URLRepository) PostURL(longURL string) (string, error) {
	var shortURL string
	ctx := context.Background()
	log.Printf("got %s", longURL)
	shortURL = shortuuid.NewWithNamespace(longURL)
	log.Printf("convSh %s", shortURL)
	_, err := ur.insertCache(ctx, shortURL, longURL)
	if err != nil {
		return "", err
	}

	gotLongURL, _ := ur.getFromDB(ctx, shortURL)
	log.Printf("got _%s_", gotLongURL)
	if gotLongURL != "" {
		log.Printf("УЖЕ СУЩЕСТВУЕТ %s", shortURL)
		return shortURL, nil
	}

	_, err = ur.insertDB(ctx, shortURL, longURL)
	if err != nil {
		return "", err
	}

	return shortURL, nil
}

// getFromCache — helper for checking shortURL in Redis
func (ur *URLRepository) getFromCache(ctx context.Context, shortURL string) (string, error) {
	longURL, err := ur.cache.Get(ctx, shortURL).Result()
	if longURL != "" {
		prometheus.GotURLFromCache.Inc()
		return longURL, nil
	}

	return "", err
}

// getFromDB — helper for checking shortURL in ClickHouse
func (ur *URLRepository) getFromDB(ctx context.Context, shortURL string) (string, error) {
	log.Printf("getFromDB %s ", shortURL)
	row := ur.db.QueryRow(
		ctx,
		`SELECT long_url FROM tbank_academy.short_to_long WHERE short_url == $1`,
		shortURL,
	)

	var longURL string
	if err := row.Scan(&longURL); err != nil {
		return "", urlerrors.Error{
			Err:  urlerrors.ErrCannotScanURLFromDB,
			Desc: err.Error(),
		}
	}

	return longURL, nil
}

// insertCache — helper for inserting shortURL -> longURL into Redis
func (ur *URLRepository) insertCache(ctx context.Context, shortURL, longURL string) (string, error) {
	// Zero expiration means the key has no expiration time.
	err := ur.cache.Set(ctx, shortURL, longURL, 0).Err()
	if err != nil {
		return "", err
	}
	return shortURL, nil
}

// insertDB — helper for inserting shortURL -> longURL into ClickHouse
func (ur *URLRepository) insertDB(ctx context.Context, shortURL, longURL string) (string, error) {
	log.Printf("insertDB %s %s", shortURL, longURL)
	_ = ur.db.QueryRow(
		ctx,
		`INSERT INTO tbank_academy.short_to_long (short_url, long_url) VALUES ($1, $2)`,
		shortURL,
		longURL,
	)
	return shortURL, nil
}
