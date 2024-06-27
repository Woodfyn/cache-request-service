package core

import "errors"

var (
	ErrCacheIsExpired           = errors.New("cache is expired")
	ErrCacheIsExpiredOrNotFound = errors.New("cache is expired or not found")
)
