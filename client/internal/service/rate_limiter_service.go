package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
	"github.com/tiago-g-sales/rate-limiter-goexpert/client/internal/config"
	"github.com/tiago-g-sales/rate-limiter-goexpert/client/internal/database"
	"github.com/tiago-g-sales/rate-limiter-goexpert/client/internal/model"
)


func FormatParameter(ctx context.Context) *model.Parameter {
	
	requestTimeBlock := viper.GetString("HTTP_REQUEST_TIME_BLOCK")
	duration, errConv := time.ParseDuration(requestTimeBlock)
	if errConv != nil {
		fmt.Println("Error Parse HTTP_REQUEST_TIME_BLOCK")
	}

	apiKeyHttpTPS := viper.GetString("HTTP_REQUEST_APIKEY_TPS")
	apiKey := strings.Split(apiKeyHttpTPS, ":")


	parameter := model.Parameter{	
		ApiKeyParameter:apiKey[0],
		TpsLimitApiKey: apiKey[1],
		TpsLimitIP: viper.GetString("HTTP_REQUEST_IP_TPS"),
		RequestTimeBlock: float64(duration),
		RequestBlocked: false,
	}

	return &parameter

}

func GetParameter(ctx context.Context, parameter model.Parameter) *model.Parameter {
	
	database :=  database.DatabaseImpl{
		Client: config.ConfigRedis(),
	}
	
	p := database.ConsultarParametros(ctx, parameter)
	
	return p

}

func UpdateRateLimiter(ctx context.Context, parameter model.Parameter){

	database :=  database.DatabaseImpl{
	Client: config.ConfigRedis(),
	}
	parameter.TpsCount++
	database.AtualizarParametros(ctx, parameter)
	
	fmt.Println("Update rate limiter")
}

func InserirParametros(ctx context.Context, parameter model.Parameter) {
	
	time_now := time.Now()

	database :=  database.DatabaseImpl{
		Client: config.ConfigRedis(),
	}
	parameter.TpsCount = 1
	parameter.RequestTime = time_now
	database.InserirParametros(ctx, parameter)
	
}