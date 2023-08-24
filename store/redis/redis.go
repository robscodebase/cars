package redis

import (
	"cars/config"
	"fmt"
	"github.com/go-redis/redis/v7"
	"strconv"
	"time"
)

const defaultSessionDuration = 3600

type Redis struct {
	client *redis.Client
}

func NewRedis() (*Redis, error) {
	conf := config.Get().Redis
	client := redis.NewClient(&redis.Options{
		Addr: conf.Host + ":" + conf.Port,
		DB:   0,
	})
	_, err := client.Ping().Result()
	if err != nil {
		return nil, fmt.Errorf("failed to ping redis: %v", err)
	}
	return &Redis{
		client: client,
	}, nil
}

func (r *Redis) SetToken(token string) error {
	d, _ := strconv.Atoi(config.Get().Redis.SessionDuration)
	if d == 0 {
		return fmt.Errorf("session duration is not set, using default value: %v", defaultSessionDuration)
	}
	return r.client.Set(token, "", defaultSessionDuration*time.Second).Err()
}

func (r *Redis) TokenExists(token string) (int64, error) {
	return r.client.Exists(token).Result()
}

func (r *Redis) DeleteToken(token string) error {
	return r.client.Del(token).Err()
}
