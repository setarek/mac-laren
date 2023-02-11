package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"

	"mac-laren/order_processor/internal/model"
	"mac-laren/order_processor/internal/repository"
	"mac-laren/order_processor/internal/service/consumer"
	"mac-laren/pkg/config"
	"mac-laren/pkg/logger"
	"mac-laren/pkg/mysql"
	"mac-laren/pkg/redis"
)

const APP_NAME = "order_processor"

func main() {

	currentPath := "/go/src/order_processor"
	config, err := config.InitConfig(APP_NAME, currentPath, "config")
	if err != nil {
		logger.Logger.Error().Err(err).Msg("error while initializing config")
	}

	redisClient := redis.GetRedisClient(config)
	redisRepository := redis.NewRedirectRepository(redisClient)

	db, err := mysql.InitConnection(config)
	if err != nil {
		logger.Logger.Error().Err(err).Msg("error while initializing mysql connection")
		os.Exit(3)
	}

	err = model.MigrateDB(db)
	if err != nil {
		logger.Logger.Error().Err(err).Msg("error while migrating tables")
		os.Exit(3)
	}

	dbRepository := repository.NewOrderRepository(db)

	content, err := ioutil.ReadFile(fmt.Sprintf("%s/order_processor.lua", currentPath))
	if err != nil {
		logger.Logger.Error().Err(err).Msg("error while reading lua script")
	}
	result, err := redisRepository.FunctionLoad(string(content))
	if err != nil {
		logger.Logger.Error().Err(err).Msg("error while loading lua script at db redis")
	}
	logger.Logger.Info().Msgf("lua script loaded at db redis. result: %v", result)

	go consumer.ConsumeOrder(config, redisRepository, dbRepository)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	for {

		select {

		case <-interrupt:
			return
		}
	}

}
