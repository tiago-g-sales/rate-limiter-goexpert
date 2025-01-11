package model

import (
	"time"
)


type Parameter struct {
	TpsLimitApiKey 				string 		`json:"tpsLimitApiKey"`
	TpsLimitIP 					string 		`json:"tpsLimitIP"`
	RequestTimeBlock 			float64 	`json:"requestTimeBlock"`
	TpsCount 					float64		`json:"tpsCount"`
	RequestTime 				time.Time 	`json:"requestTime"`
	ApiKeyParameter 			string 		`json:"apiKeyParameter"`
	ApiKeyRequest 				string 		`json:"apiKeyRequest"`
	Ip 							string 	  	`json:"ip"`
	InitialTimeRequestBlocked 	time.Time  	`json:"initialTimeRequestBlocked"`
	RequestBlocked				bool 		
	
}

type MsgBlocked struct {
	Message 		string 		`json:"message"`
}




