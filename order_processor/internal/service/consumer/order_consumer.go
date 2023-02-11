package consumer

import (
	"encoding/json"
	"fmt"
	"mac-laren/order_processor/internal/model"
	"mac-laren/pkg/config"
	"mac-laren/pkg/logger"
	"mac-laren/pkg/redis"
	"time"
)

func ConsumeOrder(config *config.Config, redisRepository *redis.RedisRepository) {

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

		fmt.Println("result", order)

		// save to mysql
		// notify use
	}

}
