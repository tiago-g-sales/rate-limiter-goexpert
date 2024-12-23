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
	ApiKeyParameter  	string 	`redis:"apiKeyParameter"`
	ApiKeyRequest  		string 	`redis:"apiKeyRequest"`
	Ip					string 	`redis:"ip"`
	RequestBlocked		bool 	`redis:"requestBlocked"`
 }
 


func (db DatabaseImpl) ConsultarParametros(ctx context.Context, parameter model.Parameter) *model.Parameter {

	defer db.Client.Close()
	// Ping the Redis server to check the connection
	pong, err := db.Client.Ping(ctx).Result()
	if err != nil {
	log.Fatal("Error connecting to Redis:", err)
	}
	fmt.Println("Connected to Redis:", pong)

	p := ParameterRedis{}
	value, err := db.Client.Get(ctx, parameter.Ip ).Result()
	if err != nil {
		return nil		
	}
	err = p.UnmarshalBinary([]byte(value))
	if err != nil {
		return nil		
	}

	parameterResult := model.Parameter{	
		TpsLimitApiKey: p.TpsLimitApiKey,
		TpsLimitIP: p.TpsLimitIP,
		RequestTimeBlock: float64(p.RequestTimeBlock),
		TpsCount: float64(p.TpsCount),
		RequestTime: time.Unix(int64(p.RequestTime), 0),  
		ApiKeyParameter: p.ApiKeyParameter,
		ApiKeyRequest: p.ApiKeyRequest,
		Ip: p.Ip,
		RequestBlocked: p.RequestBlocked,
	}

	return &parameterResult
}

func (db DatabaseImpl) InserirParametros(ctx context.Context, parameter model.Parameter) {

	defer db.Client.Close()
	p := ParameterRedis{
		TpsLimitApiKey: parameter.TpsLimitApiKey,
		TpsLimitIP: parameter.TpsLimitIP,
		RequestTimeBlock: int(parameter.RequestTimeBlock),
		TpsCount: int(parameter.TpsCount),
		RequestTime: int(parameter.RequestTime.Unix()),  
		ApiKeyParameter: parameter.ApiKeyParameter,
		ApiKeyRequest: parameter.ApiKeyRequest,
		Ip: parameter.Ip,
		RequestBlocked: parameter.RequestBlocked,
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

	err = db.Client.Set(ctx, parameter.Ip, value, 10*time.Minute).Err()
	if err != nil {
		log.Fatal("Error Insert to Redis:", err)
	}

	v, err := db.Client.Get(ctx, parameter.Ip).Result()
	if err != nil {
		log.Fatal("Error Select to Redis:", err)
	}
	fmt.Printf("The name of the parameter is %s \n", v)



}

func (db DatabaseImpl) AtualizarParametros(ctx context.Context, parameter model.Parameter) {

	defer db.Client.Close()
	p := ParameterRedis{
		TpsLimitApiKey: parameter.TpsLimitApiKey,
		TpsLimitIP: parameter.TpsLimitIP,
		RequestTimeBlock: int(parameter.RequestTimeBlock),
		TpsCount: int(parameter.TpsCount),
		RequestTime: int(parameter.RequestTime.Unix()),  
		ApiKeyParameter: parameter.ApiKeyParameter,
		ApiKeyRequest: parameter.ApiKeyRequest,
		Ip: parameter.Ip,
		RequestBlocked: parameter.RequestBlocked,
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

	err = db.Client.Set(ctx, parameter.Ip, value, 10*time.Minute).Err()
	if err != nil {
		log.Fatal("Error Insert to Redis:", err)
	}

	v, err := db.Client.Get(ctx, parameter.Ip).Result()
	if err != nil {
		log.Fatal("Error Select to Redis:", err)
	}
	fmt.Printf("The name of the parameter is %s \n", v)



}



func (db DatabaseImpl) BloquerParametros(ctx context.Context, parameter model.Parameter) {
	
	defer db.Client.Close()

	p := ParameterRedis{
 
		ApiKeyParameter: parameter.ApiKeyParameter,
		ApiKeyRequest: parameter.ApiKeyRequest,
		Ip: parameter.Ip,
		RequestBlocked: parameter.RequestBlocked,
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
	fmt.Println("RequestTimeBlock :", p.RequestTimeBlock)

	err = db.Client.Del(ctx, p.Ip).Err()
	if err != nil {
		log.Fatal("Error Delete to Redis:", err)
	}

	//err = db.Client.Set(ctx, p.Ip, value, 0).Err()
	//if err != nil {
//		log.Fatal("Error Insert to Redis:", err)
//	}

	fmt.Printf("The name of the parameter is %s \n",value)



}


func (m *ParameterRedis) MarshalBinary() ([]byte, error) {
	
	return json.Marshal(m)
}

func (m *ParameterRedis) UnmarshalBinary(data []byte) error {
    return json.Unmarshal(data, m)
}