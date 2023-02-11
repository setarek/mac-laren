package main

import (
	"fmt"

	"mac-laren/order/internal/handler"
	"mac-laren/order/internal/router"
	"mac-laren/pkg/config"
	"mac-laren/pkg/redis"
)

const APP_NAME = "order"

func main() {
	fmt.Println("hello world")

	config, err := config.InitConfig(APP_NAME, "config")
	if err != nil {
		panic(fmt.Errorf("error while initializing config: %v+", err))
	}

	redisClient := redis.GetRedisClient(config)
	redisRepository := redis.NewRedirectRepository(redisClient)

	router := router.New()
	v1 := router.Group("/api/v1")
	handler := handler.NewHandler(redisRepository)
	handler.Register(v1)

	router.Start(fmt.Sprintf("%s:%s", config.GetString("hostname"), config.GetString("port")))

}
