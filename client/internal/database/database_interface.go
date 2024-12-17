package database

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

type Database interface {
	ConsultarParametros(context.Context )
	InserirParametros()
}

type DatabaseImpl struct {	

	Client *redis.Client
}


func (db DatabaseImpl) ConsultarParametros(ctx context.Context) {

	// Ping the Redis server to check the connection
	pong, err := db.Client.Ping(ctx).Result()
	if err != nil {
	log.Fatal("Error connecting to Redis:", err)
	}
	fmt.Println("Connected to Redis:", pong)

}