package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/redis/go-redis/v9/maintnotifications"
)

type Config struct {
	Host    string `env:"REDIS_HOST"`
	Port    int    `env:"REDIS_PORT"`
	RedisDB int    `env:"REDIS_DB"`
}

type Client struct {
	*redis.Client
}

var redisInstance *Client

func Init(ctx context.Context, c *Config) error {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Host, c.Port),
		Password: "",
		DB:       c.RedisDB,
		MaintNotificationsConfig: &maintnotifications.Config{
			Mode: maintnotifications.ModeDisabled,
		},
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("failed to connect to redis client: %v", err)
	}

	redisInstance = &Client{client}

	return nil
}

func Set(ctx context.Context, k string, v any, exp time.Duration) error {
	return redisInstance.Client.Set(ctx, k, v, exp).Err()
}

func Get(ctx context.Context, k string) string {
	return redisInstance.Client.Get(ctx, k).String()
}
