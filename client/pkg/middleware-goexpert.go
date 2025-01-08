package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/tiago-g-sales/rate-limiter-goexpert/client/internal/config"
	"github.com/tiago-g-sales/rate-limiter-goexpert/client/internal/database"
	"github.com/tiago-g-sales/rate-limiter-goexpert/client/internal/model"
)

const(
	INVALID_IP_ADRESS = "invalid Ip Adress"
	MSG_BLOCKED = "you have reached the maximum number of requests or actions allowed within a certain time frame"
	ERROR = "Error Limit exceeded"
)


func RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add rate limiter logic here
		isValid := true
		ctx := context.Background()
		database :=  database.DatabaseImpl{
			Client: config.ConfigRedis(),
		}
	

		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			http.Error(w, INVALID_IP_ADRESS, http.StatusUnprocessableEntity)
			return
		}	
		parameter := model.Parameter{
			Ip: ip,
		}

		p := database.ConsultarParametros(ctx, parameter)
		if p != nil {
			fmt.Println("Validate rate limiter")
			isValid = validateRateLimiter(*p)
		}
		
		if !isValid {
			bloquerRequest(*p)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(model.MsgBlocked{Message: MSG_BLOCKED})	
		}

		
		next.ServeHTTP(w, r)
	})
}

func bloquerRequest( parameter model.Parameter ) {
	
	ctx := context.Background()
	database :=  database.DatabaseImpl{
		Client: config.ConfigRedis(),
	}

	parameter.RequestBlocked = true
	database.BloquerParametros(ctx, parameter)	

}



func validateRateLimiter( parameter model.Parameter ) bool{


	if parameter.RequestBlocked == bool(true) {
		fmt.Println("Request bloqued!")
		return false
	}

	now := time.Now()
	tpsIpTimeRequest := float64(now.Sub(parameter.RequestTime).Seconds())
	parameter.TpsCount++
	tpsRequest := parameter.TpsCount / tpsIpTimeRequest
	var tpsLimit float64
	fmt.Println("TpsRequest: ", tpsRequest)
	if parameter.ApiKeyRequest  != "" {
		tpsLimit =  convertStrintToFloat64(parameter.TpsLimitApiKey)
	} else {
		tpsLimit = convertStrintToFloat64(parameter.TpsLimitIP)
	}

	if tpsRequest > tpsLimit {
		
		return false
	}

	return true
}

func convertStrintToFloat64(s string) float64 {
	i, err := strconv.ParseFloat(s, 10)
	if err != nil {
		fmt.Println("Error converting string to int64")
	}
	return i
}