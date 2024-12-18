package database

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/tiago-g-sales/rate-limiter-goexpert/client/internal/model"
)

type Database interface {
	ConsultarParametros(context.Context, model.Parameter ) *model.Parameter
	InserirParametros(ctx context.Context, parameter model.Parameter )
}

type DatabaseImpl struct {	

	Client *redis.Client
}

type ParameterRedis struct {
	TpsLimitApiKey 		string 	`redis:"tpsLimitApiKey"`
	TpsLimitIP 			string 	`redis:"tpsLimitIP"`
	RequestTimeBlock  	int 	`redis:"requestTimeBlock"`
	TpsCount 			int    	`redis:"tpsCount"`
	RequestTime 		int 	`redis:"requestTime"`
	ApiKey 				string 	`redis:"apiKey"`
	Ip					string 	`redis:"ip"`
 }
 


func (db DatabaseImpl) ConsultarParametros(ctx context.Context, parameter model.Parameter) *model.Parameter {

	// Ping the Redis server to check the connection
	pong, err := db.Client.Ping(ctx).Result()
	if err != nil {
	log.Fatal("Error connecting to Redis:", err)
	}
	fmt.Println("Connected to Redis:", pong)

	var p ParameterRedis
	err = db.Client.Get(ctx, parameter.Ip ).Scan(&p)
	if err != nil {
		return nil		
	}

	parameterResult := model.Parameter{	
		TpsLimitApiKey: p.TpsLimitApiKey,
		TpsLimitIP: p.TpsLimitIP,
		RequestTimeBlock: float64(p.RequestTimeBlock),
		TpsCount: int64(p.TpsCount),
		RequestTime: time.Unix(int64(p.RequestTime), 0),  
		ApiKey: p.ApiKey,
		Ip: p.Ip,
	}

	return &parameterResult
}

func (db DatabaseImpl) InserirParametros(ctx context.Context, parameter model.Parameter) {


	p := ParameterRedis{
		TpsLimitApiKey: parameter.TpsLimitApiKey,
		TpsLimitIP: parameter.TpsLimitIP,
		RequestTimeBlock: int(parameter.RequestTimeBlock),
		TpsCount: int(parameter.TpsCount),
		RequestTime: int(parameter.RequestTime.Unix()),  
		ApiKey: parameter.ApiKey,
		Ip: parameter.Ip,
	}

	value, err := p.MarshalBinary()
	if err != nil {
		log.Fatal("Error MarshalBinary to Redis:", err)
	}

	// Ping the Redis server to check the connection
	pong, err := db.Client.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Error connecting to Redis:", err)
	}
	fmt.Println("Connected to Redis:", pong)

	err = db.Client.Set(ctx, parameter.Ip, value, 5*time.Minute).Err()
	if err != nil {
		log.Fatal("Error Insert to Redis:", err)
	}

	v, err := db.Client.Get(ctx, parameter.Ip).Result()
	if err != nil {
		log.Fatal("Error Select to Redis:", err)
	}
	fmt.Printf("The name of the parameter is %s \n", v)



}


func (m *ParameterRedis) MarshalBinary() ([]byte, error) {
	
	return json.Marshal(m)
}

