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

func (r *RedisRepository) FunctionLoad(function string) (interface{}, error) {
	return r.rc.Do("FUNCTION", "LOAD", function).Result()
}

func (r *RedisRepository) FCallLPop(key string) (interface{}, error) {
	result, err := r.rc.Do("FCALL", "lpop_order", "1", key).Result()
	if result == nil && err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return result, nil
}
