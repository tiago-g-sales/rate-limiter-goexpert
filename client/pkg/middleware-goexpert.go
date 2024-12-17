package pkg

import (
	"context"
	"fmt"
	"net/http"

	"github.com/tiago-g-sales/rate-limiter-goexpert/client/internal/config"
	"github.com/tiago-g-sales/rate-limiter-goexpert/client/internal/database"
)


func RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add rate limiter logic here
		
		ctx := context.Background()

		database :=  database.DatabaseImpl{
			Client: config.ConfigRedis(),
		}
		
		database.ConsultarParametros(ctx)

		
		fmt.Println("Teste rate limiter")
		
		next.ServeHTTP(w, r)
	})
}