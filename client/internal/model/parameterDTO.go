package model

import (
	"time"
)


type Parameter struct {
	TpsLimitApiKey string 		`json:"tpsLimitApiKey"`
	TpsLimitIP string 			`json:"tpsLimitIP"`
	RequestTimeBlock float64 	`json:"requestTimeBlock"`
	TpsCount int64				`json:"tpsCount"`
	RequestTime time.Time 		`json:"requestTime"`
	ApiKey string 				`json:"apiKey"`
	Ip string 	  				`json:"ip"`
}




