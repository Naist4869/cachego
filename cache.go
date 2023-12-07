package cachego

import (
	"context"
	"time"
)

type (
	// Cache is the top-level cache interface
	Cache interface {
		// Contains check if a cached key exists
		Contains(key string) bool

		// Delete remove the cached key
		Delete(key string) error

		// Fetch retrieve the cached key value
		Fetch(key string) (string, error)

		// FetchMulti retrieve multiple cached keys value
		FetchMulti(keys []string) map[string]string

		// Flush remove all cached keys
		Flush() error

		// Save cache a value by key
		Save(key string, value string, lifeTime time.Duration) error
	}
)

type (
	CtxCache interface {
		Cache

		FetchWithCtx(ctx context.Context, key string) (string, error)

		SaveWithCtx(ctx context.Context, key string, value string, lifeTime time.Duration) error
	}
)
