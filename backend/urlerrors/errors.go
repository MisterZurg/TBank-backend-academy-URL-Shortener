package urlerrors

import (
	"errors"
	"fmt"
)

// Error — an error returned by the Secrets Manager SDK to user.
type Error struct {
	Err  error
	Desc string
}

func (e Error) Error() string {
	return fmt.Sprintf("tbank-urlshortener: error — %s: %s", e.Err.Error(), e.Desc)
}

func (e Error) Is(err error) bool {
	return errors.Is(e.Err, err)
}

func GetError(errorString string) error {
	err, ok := stringToError[errorString]
	if !ok {
		return nil
	}
	return err
}

var (
	// Set Up Errors.
	ErrCannotInitConfig     = errors.New("CANNOT_INIT_CONFIG")
	ErrCannotConnectToCache = errors.New("CANNOT_CONNECT_TO_CACHE")
	ErrCannotConnectToDB    = errors.New("CANNOT_CONNECT_TO_DB")

	// Storage Errors.
	ErrCannotFindURL       = errors.New("CANNOT_FIND_URL")
	ErrCannotScanURLFromDB = errors.New("CANNOT_SCAN_URL_FROM_DB")

	// Backend Errors.
	ErrEmptyURL = errors.New("EMPTY_URL")

	ErrInternalAppError = errors.New("INTERNAL_APP_ERROR")

	//nolint:gochecknoglobals
	stringToError = map[string]error{
		ErrCannotInitConfig.Error():     ErrCannotInitConfig,
		ErrCannotConnectToCache.Error(): ErrCannotConnectToCache,
		ErrCannotConnectToDB.Error():    ErrCannotConnectToDB,

		ErrCannotFindURL.Error():       ErrCannotFindURL,
		ErrCannotScanURLFromDB.Error(): ErrCannotScanURLFromDB,

		ErrEmptyURL.Error(): ErrEmptyURL,

		ErrInternalAppError.Error(): ErrInternalAppError,
	}
)
