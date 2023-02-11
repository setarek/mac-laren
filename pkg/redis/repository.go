package redis

import "github.com/go-redis/redis"

const (
	NEW_ORDER = "ORDER:NEW"
)

type RedisRepository struct {
	rc *redis.Client
}

func NewRedirectRepository(rc *redis.Client) *RedisRepository {
	return &RedisRepository{rc: rc}
}

func (r *RedisRepository) LPush(key string, value interface{}) (int64, error) {
	return r.rc.LPush(key, value).Result()
}

func (r *RedisRepository) RPop(key string) (string, error) {
	result, err := r.rc.RPop(key).Result()
	if err != nil && err != redis.Nil {
		return "", err
	} else if err == redis.Nil {
		return "", nil
	}
	return result, nil
}
