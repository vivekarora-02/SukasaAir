package redis

import (
	"context"
	"time"
)

// RedisClientInterface defines the required Redis methods.
type RedisClientInterface interface {
	SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error)
	Del(ctx context.Context, keys ...string) error
}
