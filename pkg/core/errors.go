package core

import "errors"

var (
	ErrCacheIsExpiredOrNotFound = errors.New("cache is expired or not found")

	ErrUnknownType = errors.New("unknown type")
)
