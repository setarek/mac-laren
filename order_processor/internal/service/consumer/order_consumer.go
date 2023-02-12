package consumer

import (
	"encoding/json"
	"mac-laren/order_processor/internal/model"
	"mac-laren/order_processor/internal/repository"
	"mac-laren/pkg/config"
	"mac-laren/pkg/logger"
	"mac-laren/pkg/redis"
	"time"
)

func ConsumeOrder(config *config.Config, redisRepository *redis.RedisRepository, repository *repository.OrderRepository) {

	for {

		result, err := redisRepository.FCallLPop(redis.NEW_ORDER)
		if err != nil {
			logger.Logger.Error().Err(err).Msg("error while retrieving order from redis")
		}

		if result == nil {
			time.Sleep(1 * time.Second)
			continue
		}

		var order model.Order
		err = json.Unmarshal([]byte(result.(string)), &order)
		if err != nil {
			logger.Logger.Error().Err(err).Msg("error while unmarshalling order")
		}

		logger.Logger.Info().Str("title", order.Title).Msg("new order comes for process")

		err = repository.CreateOrder(order)
		if err != nil {
			logger.Logger.Error().Err(err).Msg("error while creating order")
		}

	}

}
