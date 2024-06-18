package repository

import (
	"context"
	"log"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/lithammer/shortuuid"
	"github.com/redis/go-redis/v9"

	"github.com/MisterZurg/TBank_URL_shortener/backend/urlerrors"
)

// URLRepository — ...
type URLRepository struct {
	// *sqlx.DB
	cache *redis.Client
	db    driver.Conn
}

// GetURL — ...
func (ur *URLRepository) GetURL(shortURL string) (string, error) {
	var longURl string

	ctx := context.Background()
	// log.Printf("repo got from svc :> %s", url)
	longURl, err := ur.cache.Get(ctx, shortURL).Result()

	if err == redis.Nil {
		log.Printf("NO IN CACHE")
		row := ur.db.QueryRow(
			ctx,
			`SELECT long_url FROM tbank_academy.short_to_long WHERE short_url == $1`,
			shortURL,
		)
		if err := row.Scan(&longURl); err != nil {
			log.Printf("%v", urlerrors.Error{
				Err:  urlerrors.ErrCannotScanURLFromDB,
				Desc: err.Error(),
			})
		}

	}

	return longURl, nil
}

// PostURL — ...
func (ur *URLRepository) PostURL(url string) (string, error) {
	var shortURL string
	ctx := context.Background()

	// log.Printf("%s", url)
	shortURL = shortuuid.NewWithNamespace(url)

	// Zero expiration means the key has no expiration time.
	err := ur.cache.Set(ctx, shortURL, url, 0).Err()
	if err != nil {
		return "", err
	}

	row := ur.db.QueryRow(
		ctx,
		`SELECT * FROM tbank_academy.short_to_long long_url WHERE short_url == $1`,
		shortURL,
	)
	if row != nil {
		log.Printf("УЖЕ СУЩЕСТВУЕТ %s", shortURL)
	}
	return shortURL, nil
}
