package service

import (
	"context"
	"fmt"
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

	parameter := model.Parameter{	
		TpsLimitApiKey: viper.GetString("HTTP_REQUEST_APIKEY_TPS"),
		TpsLimitIP: viper.GetString("HTTP_REQUEST_IP_TPS"),
		RequestTimeBlock: float64(duration),
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

func ValidateRateLimiter(ctx context.Context, parameter model.Parameter){

	fmt.Println("Validating rate limiter")
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