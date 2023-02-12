package main

import (
	"fmt"
	"os"

	"mac-laren/order/internal/handler"
	"mac-laren/order/internal/router"
	"mac-laren/pkg/config"
	"mac-laren/pkg/logger"
	"mac-laren/pkg/redis"
)

const APP_NAME = "order"

func main() {

	currentPath, _ := os.Getwd()
	// init config
	config, err := config.InitConfig(APP_NAME, currentPath, "config")
	if err != nil {
		logger.Logger.Error().Err(err).Msg("error while initializing config")
	}

	// init redis client
	redisClient := redis.GetRedisClient(config)
	redisRepository := redis.NewRedisRepository(redisClient)

	// initialize router
	router := router.New()
	v1 := router.Group("/api/v1")
	handler := handler.NewHandler(redisRepository)
	handler.Register(v1)

	router.Start(fmt.Sprintf("%s:%s", config.GetString("hostname"), config.GetString("port")))

}
