package config

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)


func ConfigRedis() *redis.Client {
	
	client := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("REDIS_HOST")+ viper.GetString("REDIS_PORT"), // Replace with your Redis server address
		Password: viper.GetString("REDIS_PASSWORD"),                // No password for local development
		DB:       0,                 // Default DB
	   })

	return client   
}