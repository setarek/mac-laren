package handler

import (
	"mac-laren/pkg/redis"
)

type OrderHandler struct {
	RedisRepository *redis.RedisRepository
}

func NewHandler(redisRepository *redis.RedisRepository) *OrderHandler {
	return &OrderHandler{
		RedisRepository: redisRepository,
	}
}
