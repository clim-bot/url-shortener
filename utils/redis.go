package utils

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/clim-bot/url-shortener/config"
)

func SetCache(key string, value string, expiration time.Duration) error {
	err := config.RedisClient.Set(context.Background(), key, value, expiration).Err()
	return err
}

func GetCache(key string) (string, error) {
	val, err := config.RedisClient.Get(context.Background(), key).Result()
	return val, err
}

func BlacklistToken(ctx context.Context, token string, expiration time.Duration) error {
	return config.RedisClient.Set(ctx, token, "blacklisted", expiration).Err()
}

func IsTokenBlacklisted(ctx context.Context, token string) bool {
	val, err := config.RedisClient.Get(ctx, token).Result()
	if err == redis.Nil {
		return false
	}
	return val == "blacklisted"
}
